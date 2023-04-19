From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build .
CMD ["/app/user-service"]

