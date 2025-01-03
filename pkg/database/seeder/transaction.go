package seeder

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type transactionSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewTransactionSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *transactionSeeder {
	return &transactionSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *transactionSeeder) Seed() error {
	paymentMethods := []string{"mandiri", "bri", "bni"}

	cards, err := r.db.GetCards(r.ctx, db.GetCardsParams{
		Column1: "",
		Limit:   10,
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("failed to get card list", zap.Error(err))
		return fmt.Errorf("failed to get card list: %w", err)
	}

	if len(cards) == 0 {
		r.logger.Error("no cards available for transaction seeding")
		return fmt.Errorf("no cards available for transaction seeding")
	}

	merchants, err := r.db.GetMerchants(r.ctx, db.GetMerchantsParams{
		Column1: "",
		Limit:   10,
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("failed to get merchant list", zap.Error(err))
		return fmt.Errorf("failed to get merchant list: %w", err)
	}

	if len(merchants) == 0 {
		r.logger.Error("no merchants available for transaction seeding")
		return fmt.Errorf("no merchants available for transaction seeding")
	}

	for i := 0; i < 10; i++ {
		selectedCard := cards[rand.Intn(len(cards))]
		selectedMerchant := merchants[rand.Intn(len(merchants))]

		selectedPaymentMethod := paymentMethods[rand.Intn(len(paymentMethods))]

		transactionAmount := int32(rand.Intn(1000000-50000) + 50000)

		request := db.CreateTransactionParams{
			CardNumber:      selectedCard.CardNumber,
			Amount:          transactionAmount,
			PaymentMethod:   selectedPaymentMethod,
			MerchantID:      selectedMerchant.MerchantID,
			TransactionTime: time.Now(),
		}

		_, err := r.db.CreateTransaction(r.ctx, request)
		if err != nil {
			r.logger.Error("failed to seed transaction", zap.Int("transaction", i+1), zap.Error(err))
			return fmt.Errorf("failed to seed transaction %d: %w", i+1, err)
		}
	}

	r.logger.Info("transaction seeded successfully")

	return nil
}
