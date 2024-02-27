package postgres

import (
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewClient(logger logger.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
	)

	con, err := sql.Open("postgres", connStr)

	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, errors.New("failed to connect to database")
	}

	return con, nil
}
