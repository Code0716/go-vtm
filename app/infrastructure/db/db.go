// Package db gorm init
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Code0716/go-vtm/app/util"
	"github.com/go-sql-driver/mysql"
	gormMySQLDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// NewDBConn initializes DB connection.
func NewDBConn(env util.Environment) (conn *sql.DB, close func() error, err error) {
	dsn, err := BuildMySQLConnectionString(env)
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	close = func() error {
		if err := sqlDB.Close(); err != nil {
			return err
		}
		return err
	}
	return sqlDB, close, nil

}

// NewDB initializes db
func NewDB(conn *sql.DB, env util.Environment) (*SQLHandler, error) {
	gormDB, err := gorm.Open(gormMySQLDriver.New(gormMySQLDriver.Config{
		Conn: conn,
	}), &gorm.Config{
		Logger: gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				LogLevel: logLevelToGormLogLevel(env.AppLogLevel),
			}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = gormDB
	return sqlHandler, nil
}

func logLevelToGormLogLevel(logLevel string) gormLogger.LogLevel {
	switch logLevel {
	case "debug":
		return gormLogger.Info
	case "warn":
		return gormLogger.Error
	default:
		return gormLogger.Info
	}
}

// BuildMySQLConnectionString builds mysql connection string.
func BuildMySQLConnectionString(env util.Environment) (string, error) {
	mysqlCfg := mysql.NewConfig()

	mysqlCfg.DBName = env.DBName
	mysqlCfg.Net = "tcp"
	mysqlCfg.Addr = fmt.Sprintf("%s:%s", env.DBHost, env.DBPort)
	mysqlCfg.User = env.DBAdmin
	mysqlCfg.Passwd = env.DBPassword
	mysqlCfg.ParseTime = true
	loc, err := time.LoadLocation(env.DBTimezone)
	if err != nil {
		return "", err
	}
	mysqlCfg.Loc = loc
	// mysqlCfg.Collation = env.DBCharset
	ret := mysqlCfg.FormatDSN()
	return ret, nil
}
