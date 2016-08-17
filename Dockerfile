FROM golang:1.6-wheezy
ENV sourcesdir /go/src/github.com/microservices-demo/user/
ENV MONGO_HOST localhost
ENV HATEAOS user
ENV USER_DATABASE mongodb

COPY . ${sourcesdir}

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
RUN echo "deb http://repo.mongodb.org/apt/debian wheezy/mongodb-org/3.2 main" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list
RUN apt-get update
RUN apt-get install -y mongodb-org
RUN mkdir -p /data/db
RUN ${sourcesdir}scripts/mongo_create_insert.sh
RUN go get -v github.com/Masterminds/glide && cd ${sourcesdir} && glide install && go install

ENTRYPOINT ${sourcesdir}scripts/start.sh
EXPOSE 8084
