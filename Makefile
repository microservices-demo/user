NAME = weaveworksdemos/user
INSTANCE = user

default: build

build:
	docker build -t $(NAME)-dev .

copy:
	docker create --name $(INSTANCE) $(NAME)-dev
	docker cp $(INSTANCE):/app/main $(shell pwd)/app
	docker rm $(INSTANCE)

release:
	docker build -t $(NAME) -f Dockerfile-release .

run:
	docker run --rm -p 8084:80 --name $(INSTANCE) $(NAME)
