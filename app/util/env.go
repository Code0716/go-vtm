package util

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// Environment has application env variables.
type Environment struct {
	EnvCode     string
	AppPort     string
	AppHost     string
	AppEnv      string
	AppLocale   string
	AppLogLevel string
	AppDebug    bool

	DBDialect       string
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPassword      string
	DBAdminPassword string
	DBCharset       string
	DBTimezone      string

	Signingkey string
}

var env Environment
var envMux sync.Mutex
var envOnce sync.Once

const defaultPort = "8080"

// Env returns env variables.
func Env() Environment {
	envMux.Lock()
	defer envMux.Unlock()

	envOnce.Do(func() {
		env = loadEnv()
	})

	return env
}

func loadEnv() Environment {
	var e Environment

	// APP side
	e.AppHost = os.Getenv("APP_HOST")
	e.AppPort = os.Getenv("APP_LISTEN_PORT")
	if e.AppPort == "" {
		e.AppPort = defaultPort
	}

	e.AppEnv = os.Getenv("APP_ENV")
	e.AppLocale = os.Getenv("APP_LOCALE")
	e.AppLogLevel = os.Getenv("APP_LOG_LEVEL")
	isDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		fmt.Println(err)
	}
	e.AppDebug = isDebug
	// DB side
	e.DBDialect = os.Getenv("MYSQL_DIALECT")
	e.DBHost = os.Getenv("MYSQL_HOST")
	e.DBPort = os.Getenv("MYSQL_PORT")
	e.DBName = os.Getenv("MYSQL_DBNAME")
	e.DBUser = os.Getenv("MYSQL_USER")
	e.DBPassword = os.Getenv("MYSQL_PASSWORD")
	e.DBAdminPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	e.DBCharset = os.Getenv("MYSQL_CHARSET")
	e.DBTimezone = os.Getenv("MYSQL_TIMEZONE")
	e.Signingkey = os.Getenv("SIGNINGKEY")
	e.EnvCode = os.Getenv("ENV_CODE")
	return e
}

// GetAPIPath get path
// TODO:環境によって読み替えが必要だが、後ほど。
func GetAPIPath(e Environment) string {
	addr := fmt.Sprintf("%s:%s", e.AppHost, e.AppPort)
	return addr
}
