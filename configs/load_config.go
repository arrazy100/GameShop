package configs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Configs struct {
	CorsConfig         cors.Config
	DatabaseConnection *gorm.DB
	Fiber              fiber.Config
	AppPort            int
}

func LoadConfig(filename string, all_models ...interface{}) Configs {
	app_config_file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	var app_config AppConfig

	err = yaml.Unmarshal(app_config_file, &app_config)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Cors config
	cors := CorsConfig(app_config.Cors)

	// Config DB
	db_conn, err := DBConfig(app_config.Database)

	if err != nil {
		panic(err)
	}

	// Fiber config
	fiber_config := FiberConfig(app_config.Fiber)

	// Migrations
	MigrationConfig(db_conn, app_config.Migrations, all_models...)

	return Configs{
		CorsConfig:         cors,
		DatabaseConnection: db_conn,
		Fiber:              fiber_config,
		AppPort:            app_config.Fiber.Port,
	}
}

func CorsConfig(cors_data Cors) cors.Config {
	// if all cors config is empty, return empty config
	if (len(cors_data.AllowMethods) == 0) && (len(cors_data.AllowOrigins) == 0) && (!cors_data.AllowCredentials) {
		return cors.Config{}
	}

	// if allow methods is empty, set default value
	if len(cors_data.AllowMethods) == 0 {
		cors_data.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	}

	// if allow origins is empty, set default value
	if len(cors_data.AllowOrigins) == 0 {
		cors_data.AllowOrigins = []string{"*"}
	}

	// if allow credentials is empty, set default value
	if cors_data.AllowCredentials {
		cors_data.AllowCredentials = true
	} else {
		cors_data.AllowCredentials = false
	}

	cors_config := cors.Config{
		AllowOrigins:     strings.Join(cors_data.AllowOrigins, " "),
		AllowMethods:     strings.Join(cors_data.AllowMethods, " "),
		AllowCredentials: cors_data.AllowCredentials,
	}

	return cors_config
}

func DBConfig(database DB) (*gorm.DB, error) {
	// if any of the database config is empty, return error
	if database.Host == "" || database.User == "" || database.Password == "" || database.DBName == "" || database.Port == 0 {
		return nil, errors.New("database config is not complete")
	}

	db_host := "host=" + database.Host
	db_user := "user=" + database.User
	db_password := "password=" + database.Password
	db_name := "dbname=" + database.DBName
	db_port := "port=" + strconv.Itoa(database.Port)

	var db_sslmode string

	if database.SSLMode {
		db_sslmode = "sslmode=enable"
	} else {
		db_sslmode = "sslmode=disable"
	}

	dsn := fmt.Sprintf("%s %s %s %s %s %s", db_host, db_user, db_password, db_name, db_port, db_sslmode)
	db_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	return db_conn, err
}

func FiberConfig(fiber_data Fiber) fiber.Config {
	return fiber.Config{
		Prefork:       fiber_data.Prefork,
		StrictRouting: fiber_data.StrictRouting,
		CaseSensitive: fiber_data.CaseSensitive,
	}
}

func MigrationConfig(db_connection *gorm.DB, migrations Migrations, all_models ...interface{}) {
	// if migrations mode is not set, return
	if migrations.Mode == "" {
		return
	}

	switch migrations.Mode {
	case "auto":
		db_connection.AutoMigrate(all_models...)
	case "drop":
		db_connection.Migrator().DropTable(all_models...)
	}
}
