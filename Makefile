NAME = weaveworksdemos/user
INSTANCE = user
TESTDB = weaveworkstestuserdb
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
	docker build --no-cache -t $(INSTANCE)-dev .

dockertestdb:
	docker build -t $(TESTDB) -f users-db-test/Dockerfile --no-cache users-db-test/

dockerruntest: dockertestdb dockerdev
	docker run --name my$(TESTDB) -d -h my$(TESTDB) $(TESTDB)
	docker run --name $(INSTANCE)-dev -d -p 8084:8084 --link my$(TESTDB) -e MONGO_HOST="my$(TESTDB):27017" $(INSTANCE)-dev

docker: build
	docker build -t $(NAME) -f Dockerfile-release .

dockertravis: build
	ifeq ($(TRAVIS_BRANCH), "master")
		TAG="snapshot"
	else
		TAG=$(TRAVIS_COMMIT)
	endif
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker build -t $(NAME):$(TAG) -f Dockerfile-release .
	docker push $(NAME):$(TAG)

dockertest: dockerruntest
	scripts/testcontainer.sh
	docker stop my$(TESTDB) $(INSTANCE)-dev
	-docker rm my$(TESTDB)
	-docker rm $(TESTDB)

clean: 
	rm -rf bin
	rm -rf vendor
	-docker stop $(INSTANCE)-dev
	-docker stop my$(TESTDB)
	-docker rm my$(TESTDB)
	-docker rm $(INSTANCE)-dev

build: 
	mkdir -p bin 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(INSTANCE) main.go
