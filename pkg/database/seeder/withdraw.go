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

type withdrawSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewWithdrawSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *withdrawSeeder {
	return &withdrawSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *withdrawSeeder) Seed() error {
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
		r.logger.Error("no cards available for withdraw seeding")
		return fmt.Errorf("no cards available for withdraw seeding")
	}

	for i := 0; i < 40; i++ {
		selectedCard := cards[rand.Intn(len(cards))]

		request := db.CreateWithdrawParams{
			CardNumber:     selectedCard.CardNumber,
			WithdrawAmount: 100000,
			WithdrawTime:   time.Now(),
		}

		withdraw, err := r.db.CreateWithdraw(r.ctx, request)
		if err != nil {
			r.logger.Error("failed to seed withdraw", zap.Int("withdraw", i+1), zap.Error(err))
			return fmt.Errorf("failed to seed withdraw %d: %w", i+1, err)
		}

		if i < 20 {
			err = r.db.TrashWithdraw(r.ctx, withdraw.WithdrawID)
			if err != nil {
				r.logger.Error("failed to trash withdraw", zap.Int("withdraw", i+1), zap.Error(err))
				return fmt.Errorf("failed to trash withdraw %d: %w", i+1, err)
			}
		}
	}

	r.logger.Info("withdraw seeded successfully")

	return nil
}
