# echo-xorm

My own toy example with

- HTTP server: [labstack-echo](https://gihtub.com/labstack/echo)
- Database-driver: [go-sqlite3](https://github.com/mattn/go-sqlite3)
- ORM: [xorm](https://github.com/go-xorm/xorm)
- Authorization: [JSON Web Tokens](https://github.com/dgrijalva/jwt-go)


# Installation
## Prerequisites

```bash
go get -u github.com/labstack/echo
go get -u github.com/go-xorm/xorm
go get -u github.com/mattn/go-sqlite3
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/Sirupsen/logrus
go get -u golang.org/x/crypto/bcrypt
```

## Application
```bash
go get -u github.com/corvinusy/mygolangtoys/echo-xorm
```

## Database

Currently using *sqlite3*-database, located at '/tmp/echo-xorm-test.sqlite.db'

## Vendoring

TODO

## Testing
- Test Framework: [Goconvey](https://github.com/smartystreets/goconvey)
- HTTP-Client: [Go-resty](https://github.com/go-resty/resty)


#License

MIT
