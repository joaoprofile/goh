package environment

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joaocprofile/goh/core"
	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type dbconnection struct {
	DBDriver         string
	ConnectionString string
	LimitsConnetion  bool
	MaxOpenConns     int
	MaxIdleConns     int
}

type cachedb struct {
	CacheDriver      string
	ConnectionString string
	ExpirationCache  time.Duration
	Username         string
	Password         string
}

type security struct {
	JWTSecret  []byte
	expiration time.Time
}

type environment struct {
	APIPort      int
	Security     *security
	DBConnection *dbconnection
	CacheDB      *cachedb
}

var _instance *environment

func Inicialize() {
	if _instance == nil {
		lock.Lock()
		_instance = createEnvironment()
		defer lock.Unlock()
	}
}

func Get() *environment {
	if _instance == nil {
		Inicialize()
	}
	return _instance
}

func createEnvironment() *environment {
	path, _ := os.Getwd()
	var api_port int
	if err := godotenv.Load(filepath.Join(path, ".env")); err != nil {
		log.Fatalf(core.Red("Error loading environment variables: %v"), err)
	}

	if api_port, _ = strconv.Atoi(os.Getenv("API_PORT")); api_port == 0 {
		api_port = 4000 // default port
	}

	security := readSecurityVariables()
	dbConnection := readDBConnectionVariables()
	cacheDb := readCacheDBVariables()
	env := &environment{
		APIPort:      api_port,
		Security:     security,
		DBConnection: dbConnection,
		CacheDB:      cacheDb,
	}

	fmt.Println(core.Green(core.LOGO))
	fmt.Println(core.Red(core.VERSION))
	log.Println(core.Yellow("Environment variables started..."))
	return env
}

func readSecurityVariables() *security {
	var jwtSecrete string
	jwtSecrete = os.Getenv("JWT_SECRET")
	return &security{
		JWTSecret:  []byte(jwtSecrete),
		expiration: time.Now().Add(24 * time.Hour),
	}
}

func readDBConnectionVariables() *dbconnection {
	connectionString := os.Getenv("DSN")
	limitsConnetion, _ := strconv.ParseBool(os.Getenv("CONNECTION_LIMITS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))

	return &dbconnection{
		DBDriver:         "postgres",
		ConnectionString: connectionString,
		LimitsConnetion:  limitsConnetion,
		MaxOpenConns:     maxOpenConns,
		MaxIdleConns:     maxIdleConns,
	}
}

func readCacheDBVariables() *cachedb {
	cacheDB := os.Getenv("CACHEDB")
	cacheDBPassword := os.Getenv("CACHEDB_PASSWORD")

	return &cachedb{
		CacheDriver:      "Redis",
		ConnectionString: cacheDB,
		ExpirationCache:  time.Second * 20,
		Username:         "",
		Password:         cacheDBPassword,
	}
}
