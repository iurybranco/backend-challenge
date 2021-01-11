# Backend-challenge

**Recruiter**: July Demenjon

**Applicant**: Iury Dias

**Contact**: iiurydias@hotmail.com

#### Requirements

* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/install/)

#### Project Architecture

![alt text](https://github.com/iurybranco/backend-challenge/blob/main/arch.png?raw=true "Project Architecture")

**api**: Responsible for receiving requests to get products list. Calling discount-calculator service through gRPC, to applicable discounts.

**discount-calculator**: A service responsible for serving RPC service to check possible product discount based on own rules from requested user.

#### Running project
After cloning the project into your machine and with all requirements, you must execute the following command on root path:

```shell script
    sh run.sh
```

*Ps: Api settings is configurable on docker-composer file and discount-calculator settings on config.json inside the directory and if edited should happen before building.* 

#### Routes

##### Get product list

```GET /product```

**Request**

Params (*Via Query*)

* **userId**: User identification (optional)

Example:

```curl -XGET 'localhost:3001/product?userId=1'```

##### Response

+ **Success** 201

```json
{
  "status": "success",
  "data": [{
    "_id": 2,
    "price_in_cents": 700,
    "title": "Soda",
    "description": "The best soda ever",
    "discount": {
      "percentage": 5,
      "value_in_cents": 35
    }
  }, {
    "_id": 1,
    "price_in_cents": 150,
    "title": "Soap",
    "description": "A smelly soap to take a shower",
    "discount": {
      "percentage": 5,
      "value_in_cents": 7
    }
  }]
} 
``` 
+ **Fail** 500

```json
{
  "status": "error",
  "message": "data has been lost on server"
}   
``` 

*Ps: If user does not exist or not informed any discount will be applied. Users and products dump can be find [here](https://github.com/iurybranco/backend-challenge/blob/master/dump).* 

#### Running tests

**Requirement**

* [Golang](https://golang.org/doc/install) 1.15
* [Node](https://nodejs.org/en/)
* [Yarn](https://classic.yarnpkg.com/en/docs/install)

**Command**

You must enter into each service repository and execute the following commands.

To api:
```shell script
    yarn test
```
To discount-calculator:
```shell script
    make run-tests
```

#### Technologies


* **Express**: The most popular node web framework which provides mechanisms to build our api.
* **TypeORM**: A popular ORM which make easy the database use (used on the api).
* **Mongo**: A non relational database, easy to work, with no schema, flexible and high performance.

