From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN whoami
RUN go mod init github.com/thedevopsschool/user-service
RUN go mod tidy
RUN go build .
CMD ["/app/user-service"]

