package database_test

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Code0716/go-vtm/app/infrastructure/db"
	database "github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/util"
)

var (
	testCtx context.Context
)
var env = util.Env()

var testEnv = util.Environment{
	DBHost:     env.DBHost,
	DBPort:     env.DBPort,
	DBName:     env.DBName,
	DBUser:     "root",
	DBPassword: env.DBAdminPassword,
	DBTimezone: env.DBTimezone,
}

func getTestDB(t *testing.T, seeds []any) (db *db.SQLHandler, close func(), err error) {
	t.Helper()
	teardownFuncs := []func(){}
	close = func() {
		for i := len(teardownFuncs) - 1; i > 0; i-- {
			teardownFuncs[i]()
		}
	}

	tmpDBName := fmt.Sprintf("%v_%d", testEnv.DBName, rand.Int63()) // #nosec G404

	fmt.Printf("test '%s' DB name: %s\n", t.Name(), tmpDBName)

	testEnvCopy := testEnv
	testEnvCopy.DBName = ""

	dsn0, err := database.BuildMySQLConnectionString(testEnvCopy)
	if err != nil {
		return nil, nil, err
	}

	sqlDB0, err := sql.Open("mysql", dsn0)
	if err != nil {
		return nil, nil, err
	}

	teardownFuncs = append(teardownFuncs, func() {
		if err := sqlDB0.Close(); err != nil {
			fmt.Println(err)
		}
	})

	file0, err := os.ReadFile("../../../_init_sql/0000_create_vtm_database.sql")
	if err != nil {
		return nil, nil, err
	}

	stmt0 := strings.ReplaceAll(string(file0), testEnv.DBName, tmpDBName)
	_, err = sqlDB0.Exec(stmt0)

	if err != nil {
		return nil, nil, err
	}

	teardownFuncs = append(teardownFuncs, func() {
		stmt := fmt.Sprintf("DROP DATABASE IF EXISTS `%v`;", tmpDBName)
		_, err = sqlDB0.Exec(stmt)
		if err != nil {
			fmt.Println(err)
		}
	})

	testEnvCopy.DBName = tmpDBName

	dsn1, err := database.BuildMySQLConnectionString(testEnvCopy)

	if err != nil {
		return nil, nil, err
	}

	sqlDB1, err := sql.Open("mysql", dsn1)
	if err != nil {
		return nil, nil, err
	}
	teardownFuncs = append(teardownFuncs, func() {
		if err := sqlDB1.Close(); err != nil {
			fmt.Println(err)
		}
	})

	createTableSQLFiles := []string{}

	err = filepath.WalkDir("../../../_init_sql/", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, "create_") && strings.HasSuffix(path, ".sql") {
			createTableSQLFiles = append(createTableSQLFiles, path)
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	createTable := func(file string) error {
		file1, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		stmt1 := strings.ReplaceAll(string(file1), testEnv.DBName, tmpDBName)
		_, err = sqlDB1.Exec(stmt1)
		if err != nil {
			return err
		}

		return nil
	}

	for _, f := range createTableSQLFiles {
		err := createTable(f)
		if err != nil {
			close()
			return nil, nil, err
		}
	}

	gormDB, err := database.NewDB(sqlDB1, testEnvCopy)
	if err != nil {
		close()
		return nil, nil, err
	}
	teardownFuncs = append(teardownFuncs)

	for _, s := range seeds {
		if err := gormDB.Conn.Create(s).Error; err != nil {
			close()
			return nil, nil, err
		}
	}

	return gormDB, close, nil
}
