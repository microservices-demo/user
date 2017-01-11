#User Service
[![Build Status](https://travis-ci.org/microservices-demo/user.svg?branch=master)](https://travis-ci.org/microservices-demo/user)
[![Coverage Status](https://coveralls.io/repos/github/microservices-demo/user/badge.svg?branch=master)](https://coveralls.io/github/microservices-demo/user?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/microservices-demo/user)](https://goreportcard.com/report/github.com/microservices-demo/user)
[![](https://images.microbadger.com/badges/image/weaveworksdemos/user.svg)](http://microbadger.com/images/weaveworksdemos/user "Get your own image badge on microbadger.com")

This service covers user account storage, to include cards and addresses

## Bugs, Feature Requests and Contributing
We'd love to see community contributions. We like to keep it simple and use Github issues to track bugs and feature requests and pull requests to manage contributions.

>## API Spec

Checkout the API Spec [here](http://microservices-demo.github.io/api/index?url=https://raw.githubusercontent.com/microservices-demo/user/master/apispec/user.json)

>## Build

### Using Go natively

```bash
make build
```

### Using Docker Compose

```bash
docker-compose build
```

>## Test

```bash
make test
```

>## Run

### Natively
```bash
docker-compose up -d user-db
./bin/user -port=8080 -database=mongodb -mongo-host=localhost:27017
```

### Using Docker Compose
```bash
docker-compose up
```

>## Check

```bash
curl http://localhost:8080/health
```

>## Use

Test user account passwords can be found in the comments in `users-db-test/scripts/customer-insert.js`

### Customers

```bash
curl http://localhost:8080/customers
```

### Cards
```bash
curl http://localhost:8080/cards
```

### Addresses

```bash
curl http://localhost:8080/addresses
```

### Login
```bash
curl http://localhost:8080/login
```

### Register

```bash
curl http://localhost:8080/register
```

## Push

```bash
make dockertravisbuild
```

## Test Zipkin

To test with Zipkin

```
make
docker-compose -f docker-compose-zipkin.yml build
docker-compose -f docker-compose-zipkin.yml up
```
It takes about 10 seconds to seed data

you should see it at:
[http://localhost:9411/](http://localhost:9411)

be sure to hit the "Find Traces" button.  You may need to reload the page.

when done you can run:
```
docker-compose -f docker-compose-zipkin.yml down
```
