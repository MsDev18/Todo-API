package migrator

import (
	"database/sql"
	"fmt"
	"todo-api/repository/mysql"

	"github.com/rubenv/sql-migrate"
)

type Migrator struct {
	migrations migrate.FileMigrationSource
	dbCfg      mysql.Config
	dialect    string
}

func New(dbCfg mysql.Config, dialect string) Migrator {
	migrations := migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}
	return Migrator{
		migrations: migrations,
		dbCfg:      dbCfg,
		dialect:    dialect,
	}
}

func (m Migrator) Up() {
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbCfg.Username, m.dbCfg.Password, m.dbCfg.Host, m.dbCfg.Port, m.dbCfg.DBName)

	db, err := sql.Open(m.dialect, dataSource)
	if err != nil {
		fmt.Println("error in conecction to database in migration Up section : ", err)

	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	fmt.Printf("Applied %d Migration Up\n", n)
}

func (m Migrator) Down() {
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbCfg.Username, m.dbCfg.Password, m.dbCfg.Host, m.dbCfg.Port, m.dbCfg.DBName)

	db, err := sql.Open(m.dialect, dataSource)
	if err != nil {
		fmt.Println("error in conecction to database in migration Down section : ", err)

	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	fmt.Printf("Applied %d Migration Down\n", n)
}