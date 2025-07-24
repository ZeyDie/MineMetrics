package database

import (
	"context"
	"fmt"
	"log/slog"
	"minemetrics_golang/internal/loggers"

	"time"

	"minemetrics_golang/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

func InitDB(sqlConfig *config.SQLConfig) (*gorm.DB, error) {
	slog.Info("Initializing database connection...")

	gormConfig := initGormConfig()
	var gormDialector gorm.Dialector

	switch sqlConfig.Driver {
	case "postgres", "postgresql":
		gormDialector = postgres.Open(
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
				sqlConfig.Host,
				sqlConfig.Port,
				sqlConfig.User,
				sqlConfig.Password,
				sqlConfig.DBName,
				sqlConfig.SSL,
				sqlConfig.TimeZone,
			),
		)
	case "mysql", "mariadb":
		gormDialector = mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				sqlConfig.User,
				sqlConfig.Password,
				sqlConfig.Host,
				sqlConfig.Port,
				sqlConfig.DBName,
			),
		)
	default:
		return nil, fmt.Errorf("Unsupported database driver: %s", sqlConfig.Driver)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	connMaxLifetime, err := time.ParseDuration(sqlConfig.MaxConnLifetime)
	if err != nil {
		slog.Warn("Invalid MaxConnLifetime, using default", "error", err)
		connMaxLifetime = time.Hour
	}

	database.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	database.SetMaxIdleConns(sqlConfig.MaxIdleConns)
	database.SetConnMaxLifetime(connMaxLifetime)

	pingCtx, pingCancel := context.WithTimeout(ctx, 5*time.Second)
	defer pingCancel()

	if err := database.PingContext(pingCtx); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v\n", err)
	}

	slog.Info("Database connection established successfully")

	Connection = connect

	return connect, nil
}

func initGormConfig() *gorm.Config {
	var logMode logger.LogLevel

	switch loggers.GetLevel() {
	case slog.LevelInfo:
		logMode = logger.Info
	case slog.LevelWarn:
		logMode = logger.Warn
	case slog.LevelError:
		logMode = logger.Error
	default:
		logMode = logger.Info
	}

	gormLogger := logger.Default.LogMode(logMode)

	return &gorm.Config{
		Logger: gormLogger,
	}
}

func GetConnection() *gorm.DB {
	return Connection
}
