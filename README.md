# Todo App API

The Rest Full Api for todo application

- write with golang
- using **gin** fream work
- implement **docker-compose** file for database
- develop migration tool

## Migration Tool

install dependensie

1. get sql-driver for connecting to database
2. get thear party package for simple handling migration

```bash
$ go get -u  github.com/go-sql-driver/mysql
$ go install github.com/rubenv/sql-migrate/...@latest
```

migration command :

```bash
$ sql-migrate up -config="./repository/migrator/dbconfig.yml" -env="production"
$ sql-migrate down -config="./repository/migrator/dbconfig.yml" -env="production" -limit=1
$ sql-migrate status -config="./repository/migrator/dbconfig.yml" -env="production"
```

<!-- > you can use package migrator in this project  -->

3. handel migrations with migrator package in this project

```go
migrate := migrator.New(DBCfg, "mysql")
migrate.Up()
migrate.Down()
```
