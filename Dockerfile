From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN whoami
RUN chown -R go:build ../app 
RUN go mod init github.com/thedevopsschool/user-service
RUN go mod tidy
RUN go build .
CMD ["/app/user-service"]

