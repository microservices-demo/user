From alpinelinux/golang

WORKDIR /app
COPY . /app
RUN ls -l | grep go.mod
RUN sudo go mod tidy
RUN go build .
CMD ["/app/user-service"]

