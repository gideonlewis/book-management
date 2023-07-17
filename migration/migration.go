package migration

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}
	migrations := &migrate.FileMigrationSource{
		Dir: "/Users/quaan2hand/Desktop/workSpace/teqnological/source_code/teq-book-manage/migration",
	}

	n, err := migrate.Exec(getDB, "mysql", migrations, migrate.Up)
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}
	fmt.Println(n)
	teqlogger.GetLogger().Info("Up done!")

}

//nolint:revive
// import (
// 	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
// 	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/mysql"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	"github.com/pkg/errors"
// 	"gorm.io/gorm"

// 	"git.teqnological.asia/teq-go/teq-echo/config"
// )

// func Up(db *gorm.DB) {
// 	getDB, err := db.DB()
// 	if err != nil {
// 		teqsentry.Fatal(err)
// 		teqlogger.GetLogger().Fatal(err.Error())
// 	}

// 	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
// 	if err != nil {
// 		teqsentry.Fatal(err)
// 		teqlogger.GetLogger().Fatal(err.Error())
// 	}

// 	m, err := migrate.NewWithDatabaseInstance("file://./migration", config.GetConfig().MySQL.DBName, driver)
// 	if err != nil {
// 		teqsentry.Fatal(err)
// 		teqlogger.GetLogger().Fatal(err.Error())
// 	}

// 	err = m.Up()
// 	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		teqsentry.Fatal(err)
// 		teqlogger.GetLogger().Fatal(err.Error())
// 	}

// 	teqlogger.GetLogger().Info("Up done!")
// }
