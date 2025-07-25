package database

import (
	"fmt"
	"log/slog"
	"minemetrics_golang/internal/database/entity"
	"time"

	"minemetrics_golang/internal/config"
	"minemetrics_golang/internal/systemlog"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB
var cfg *config.SQLConfig

func NewDB(sqlConfig *config.SQLConfig) (*gorm.DB, error) {
	cfg = sqlConfig

	connection, err := initDB()

	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		return nil, err
	}

	return connection, err
}

func initDB() (*gorm.DB, error) {
	slog.Info("Initializing database connection...")

	gormConfig := initGormConfig()
	var gormDialector gorm.Dialector

	switch cfg.Driver {
	case "postgres", "postgresql":
		gormDialector = postgres.Open(
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
				cfg.Host,
				cfg.Port,
				cfg.User,
				cfg.Password,
				cfg.DBName,
				cfg.SSL,
				cfg.TimeZone,
			),
		)
	case "mysql", "mariadb":
		gormDialector = mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.User,
				cfg.Password,
				cfg.Host,
				cfg.Port,
				cfg.DBName,
			),
		)
	default:
		return nil, fmt.Errorf("Unsupported database driver: %s", cfg.Driver)
	}

	connect, err := gorm.Open(
		gormDialector,
		gormConfig,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v\n", err)
	}

	database, err := connect.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get database instance: %v\n", err)
	}

	connMaxLifetime, err := time.ParseDuration(cfg.MaxConnLifetime)
	if err != nil {
		slog.Warn("Invalid MaxConnLifetime, using default", "error", err)
		connMaxLifetime = time.Hour
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxLifetime(connMaxLifetime)

	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v\n", err)
	}

	slog.Info("Database connection established successfully")

	connect.AutoMigrate(
		&entity.ClientEntity{},
		//TODO Auto migrate in entities func
		&entity.GPU{},
		&entity.ChunkPosition{},
		//&entity.ServerEntity{},
	)

	return connect, nil
}

func initGormConfig() *gorm.Config {
	var logMode logger.LogLevel

	switch systemlog.GetLevel() {
	case slog.LevelInfo:
		logMode = logger.Info
	case slog.LevelWarn:
		logMode = logger.Warn
	case slog.LevelError:
		logMode = logger.Error
	default:
		logMode = logger.Warn
	}

	gormLogger := logger.Default.LogMode(logMode)

	return &gorm.Config{
		Logger: gormLogger,
	}
}

func GetConnection() *gorm.DB {
	if Connection == nil || Connection.Error != nil {
		connection, err := initDB()

		Connection = connection

		if err != nil {
			slog.Error("Failed to initialize database", "error", err)
			panic(err)
		}
	}

	return Connection
}

func GetTransaction() *gorm.DB {
	transaction := GetConnection().Begin()

	defer func() {
		if recove := recover(); recove != nil {
			transaction.Rollback()
		}
	}()

	return transaction
}
