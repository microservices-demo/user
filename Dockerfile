From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN chmod 777 go.mod
RUN go mod tidy
RUN go build .
CMD ["/app/user-service"]

