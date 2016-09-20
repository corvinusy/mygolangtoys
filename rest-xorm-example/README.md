# rest-xorm-example

My own toy example with
http://github.com/ant0ine/go-json-rest as rest server
http://github.com/go-xorm/xorm as a xorm database ORM

# Installation
## Requirements
* go 1.6 or 1.7
* sqlite3

## Vendored dependencies
```bash
go get -u github.com/govend/govend
govend
```

## Application
```bash
go install github.com/corvinusy/rest-xorm-example
```

## Database
Currently is using *sqlite3*-database, located in file /tmp/rest-xorm.sqlite.db

# Tests
```bash
cd $GOPATH/src/github.com/corvinusy/rest-xorm-example/restxorm
go test
```

# License
BSD
