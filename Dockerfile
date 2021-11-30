FROM golang:1.17-buster AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /user

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /user /user
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/user", "-port=8080"]
