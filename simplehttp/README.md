# `gsimpleapi`

## what-is-it-?
* get response-json from api
* parse and insert into mongo

## how-to-build-?
```sh
$ docker-compose up
```

## how-to-run-?
### - run-to-get-api-response
```sh
$ curl http://localhost:8080/api/v1/getstds
```

### - run-insert-into-db
```sh
$ docker exec -it simplehttp_client_1 bash -c "go run insert.go"
```

## example-output
```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                      NAMES
795e0d6095f6        simplehttp_client   "/bin/sleep 360000"      About an hour ago   Up About an hour                               simplehttp_client_1
632ab7883177        simplehttp_api      "go run main.go"         About an hour ago   Up About an hour    0.0.0.0:8080->8080/tcp     simplehttp_api_1
adba2d29d9e0        mongo               "docker-entrypoint.s…"   About an hour ago   Up About an hour    0.0.0.0:27017->27017/tcp   simplehttp_mongo_1
```
```sh
$ docker exec -it simplehttp_client_1 bash -c "go run insert.go"
Connected to MongoDB!
Inserted document:  ObjectID("5d9c23a692ef923710663601")
Inserted document:  ObjectID("5d9c23a692ef923710663602")
```
```sh
$ curl http://localhost:8080/api/v1/getstds
{"users":[{"name":"tin","type":"student","Age":10,"social":"https://facebook.com"},{"name":"thang","type":"teacher","Age":10,"social":"https://facebook.com"}]}
```
