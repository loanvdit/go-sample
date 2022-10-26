## How to run
go mod vendor

### Required

- Mysql

### Ready

Create a **fuji_challenge database** and import [SQL](docs/sql/db.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = fuji_challenge
TablePrefix = fuji_
...
```

### Run
```
$ cd $GOPATH/src/fuji-challenge

$ go run main.go 
```

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis