package seeder

import (
	apikey "MamangRust/paymentgatewaygrpc/pkg/api-key"
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type merchantSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewMerchantSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *merchantSeeder {
	return &merchantSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *merchantSeeder) Seed() error {
	merchantNames := []string{
		"Merchant One",
		"Merchant Two",
		"Merchant Three",
		"Merchant Four",
		"Merchant Five",
	}

	for i, name := range merchantNames {
		apiKey := apikey.GenerateApiKey()

		req := db.CreateMerchantParams{
			Name:   name,
			UserID: int32(rand.Intn(10) + 1),
			Status: "active",
			ApiKey: apiKey,
		}

		_, err := r.db.CreateMerchant(r.ctx, req)
		if err != nil {
			r.logger.Error("failed to seed merchant", zap.Int("merchant", i+1), zap.Error(err))

			return fmt.Errorf("failed to seed merchant %d: %w", i+1, err)
		}
	}

	r.logger.Info("merchant seeded successfully")

	return nil
}
