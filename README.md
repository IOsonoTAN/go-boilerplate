# A Go Boilerplate
Hi there, here is a go boilerplate which is using [MongoDB](https://github.com/mongodb/mongo-go-driver) to be database, [Gin](https://github.com/gin-gonic/gin) to be an api framework and use a docker-compose to run as local with Docker.

## What are using in the project?
- [Dot ENV](https://github.com/joho/godotenv)
- [Go Module](https://blog.golang.org/using-go-modules)
- [Go DotEnv](https://github.com/joho/godotenv)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Mongo Go Driver](https://github.com/mongodb/mongo-go-driver)

## Required
- [Docker](https://www.docker.com/) to running this project.
- MongoDB client ([Robo 3T](https://robomongo.org/), [MongoDB Compass](https://www.mongodb.com/products/compass) or what ever you want to use) to access database.

## Clone and run with Docker
```bash
$ git clone https://github.com/IOsonoTAN/go-boilerplate.git;
$ cd go-boilerplate;
$ docker-compose up;
```

## Docker ports expose
1. Go's application will expose and listen to port [5000](#)
2. MongoDB's database will expose to port [29017](#)
