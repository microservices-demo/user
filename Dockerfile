From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN sudo go mod tidy
RUN go build .
CMD ["/app/user-service"]

