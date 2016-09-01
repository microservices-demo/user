#User Service
[![Build Status](https://travis-ci.org/microservices-demo/user.svg?branch=master)](https://travis-ci.org/microservices-demo/user)
[![Coverage Status](https://coveralls.io/repos/github/microservices-demo/user/badge.svg?branch=master)](https://coveralls.io/github/microservices-demo/user?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/microservices-demo/user)](https://goreportcard.com/report/github.com/microservices-demo/user)
[![](https://images.microbadger.com/badges/image/weaveworksdemos/user.svg)](http://microbadger.com/images/weaveworksdemos/user "Get your own image badge on microbadger.com")

This service covers user account storage, to include cards and addresses

##Build
To build a local testing version:


```bash
make dockerruntest

```

To stop 

```bash
make clean

```

Accessible at localhost:8084

##Endpoints
/customers

/cards

/addresses

/login

/register

Test user account passwords can be found in the comments in users-db-test/scripts/customer-insert.js
