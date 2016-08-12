#User Service

This service covers user account storage, to include cards and addresses

##Build
To build a local testing version:


```bash
docker build -t userservice .

```

To run

```bash
docker run -p 8084:8084 -i -t userservice
```
Accessible at localhost:8084

##Endpoints
/customers
/cards
/addresses
/login
/register

Test user account passwords can be found in the comments in scripts/customer-insert.js
