NAME = weaveworksdemos/user
INSTANCE = user
TESTDB = weaveworkstestuserdb
OPENAPI = $(INSTANCE)-testopenapi

ifeq ($(TRAVIS_BRANCH), master)
	TAG=snapshot
else
	TAG=$(TRAVIS_COMMIT)
endif

default: build
pre: 
	go get -v github.com/Masterminds/glide

deps: pre
	glide install

rm-deps:
	rm -rf vendor

test:
	@glide novendor|xargs go test -v

cover:
	@glide novendor|xargs go test -v -covermode=count

coverprofile:
	go get github.com/modocache/gover
	go test -v -covermode=count -coverprofile=profile.coverprofile
	go test -v -covermode=count -coverprofile=db.coverprofile ./db
	go test -v -covermode=count -coverprofile=mongo.coverprofile ./db/mongodb
	go test -v -covermode=count -coverprofile=api.coverprofile ./api
	gover
	mv gover.coverprofile cover.profile
	rm *.coverprofile


dockerdev:
	docker build -t $(INSTANCE)-dev .

dockertestdb:
	docker build -t $(TESTDB) -f users-db-test/Dockerfile users-db-test/

dockerruntest: dockertestdb dockerdev
	docker run --name my$(TESTDB) -d -h my$(TESTDB) $(TESTDB)
	docker run --name $(INSTANCE)-dev -d -p 8084:8084 --link my$(TESTDB) -e MONGO_HOST="my$(TESTDB):27017" $(INSTANCE)-dev

docker: build
	docker build -t $(NAME) -f Dockerfile-release .

dockertravisbuild: build
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker build -t $(NAME):$(TAG) -f Dockerfile-release .
	docker push $(NAME):$(TAG)

dockertest: dockerruntest
	scripts/testcontainer.sh
	docker run -h openapi --name $(OPENAPI) --link user-dev -v $(PWD)/apispec/:/tmp/specs/\
		weaveworksdemos/openapi /tmp/specs/$(INSTANCE).json\
		http://$(INSTANCE)-dev:8084/\
		-f /tmp/specs/hooks.js

testapi: dockerruntest testapirunning

clean: 
	rm -rf bin
	rm -rf vendor
	-docker stop $(INSTANCE)-dev
	-docker stop $(OPENAPI)
	-docker stop my$(TESTDB)
	-docker rm my$(TESTDB)
	-docker rm $(INSTANCE)-dev
	-docker rm $(OPENAPI)

build: 
	mkdir -p bin 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(INSTANCE) main.go
