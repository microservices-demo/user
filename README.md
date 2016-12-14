#User Service
[![Build Status](https://travis-ci.org/microservices-demo/user.svg?branch=master)](https://travis-ci.org/microservices-demo/user)
[![Coverage Status](https://coveralls.io/repos/github/microservices-demo/user/badge.svg?branch=master)](https://coveralls.io/github/microservices-demo/user?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/microservices-demo/user)](https://goreportcard.com/report/github.com/microservices-demo/user)
[![](https://images.microbadger.com/badges/image/weaveworksdemos/user.svg)](http://microbadger.com/images/weaveworksdemos/user "Get your own image badge on microbadger.com")

This service covers user account storage, to include cards and addresses

## Build

### Go

```bash
make build
```

### Docker

```bash
make dockerdev
```

## Test

```bash
make test

```

## Run

```bash
make dockerruntest
```

## Check

```bash
chromium-browser http://localhost:8084/health
```

## Use

Test user account passwords can be found in the comments in `users-db-test/scripts/customer-insert.js`

### Customers

```bash
chromium-browser http://localhost:8084/customers
```

### Cards
```bash
chromium-browser http://localhost:8084/cards
```

### Addresses

```bash
chromium-browser http://localhost:8084/addresses
```

### Login
```bash
chromium-browser http://localhost:8084/login
```

### Register

```bash
chromium-browser http://localhost:8084/register
```

## Push

```bash
make dockertravisbuild
```