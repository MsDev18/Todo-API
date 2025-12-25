# Todo App API
The Rest Full Api for todo application  

* write with golang 
* using **gin** fream work 
* implement **docker-compose** file for database 
* develop migration tool 

## Migration Tool
install dependensie
``` bash
$ go get -u  github.com/go-sql-driver/mysql
```
migration command :
``` bash
$ sql-migrate up -config="./repository/migrator/dbconfig.yml" -env="production"
$ sql-migrate down -config="./repository/migrator/dbconfig.yml" -env="production" -limit=1
$ sql-migrate status -config="./repository/migrator/dbconfig.yml" -env="production"
```