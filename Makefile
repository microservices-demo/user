NAME = weaveworksdemos/user
INSTANCE = user

default: build

pre: 
	go get -v github.com/Masterminds/glide

deps:
	pre
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
	docker build -t $(NAME)-dev .

docker: build
	docker build -t $(NAME) -f Dockerfile-release .

rundev:
	docker run --rm -p 8084:8084 $(NAME)-dev

clean: 
	rm -rf bin
	rm -rf vendor

build: 
	mkdir -p bin 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(INSTANCE) main.go
