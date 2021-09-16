package util

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Environment has application env variables.
type Environment struct {
	AppPort     string
	AppHost     string
	AppEnv      string
	AppLocale   string
	AppLogLevel string
	AppDebug    bool

	DBDialect  string
	DBHost     string
	DBPort     string
	DBName     string
	DBAdmin    string
	DBPassword string
	DBCharset  string
	DBTimezone string

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
	envPath := os.Getenv("GO_ENV")
	err := godotenv.Load(fmt.Sprintf("%s.env", envPath))
	if err != nil {
		panic(err)
	}
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
	e.DBAdmin = os.Getenv("MYSQL_USER")
	e.DBPassword = os.Getenv("MYSQL_PASSWORD")
	e.DBCharset = os.Getenv("MYSQL_CHARSET")
	e.DBTimezone = os.Getenv("MYSQL_TIMEZONE")
	e.Signingkey = os.Getenv("SIGNINGKEY")
	return e
}

// GetAPIPath get path
// TODO:環境によって読み替えが必要だが、後ほど。
func GetAPIPath(e Environment) string {
	addr := fmt.Sprintf("%s:%s", e.AppHost, e.AppPort)
	return addr
}
