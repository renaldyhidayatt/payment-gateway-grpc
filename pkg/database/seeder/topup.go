package seeder

import (
	"context"
	"fmt"
	"math/rand"

	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"

	"go.uber.org/zap"
)

type topupSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewTopupSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *topupSeeder {
	return &topupSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *topupSeeder) Seed() error {
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
		r.logger.Error("no cards available for topup seeding")
		return fmt.Errorf("no cards available for topup seeding")
	}

	topupMethods := []string{"bri", "mandiri", "bni"}

	for _, card := range cards {
		cardNumber := card.CardNumber
		if len(cardNumber) < 5 {
			r.logger.Error("card number is too short", zap.String("card", cardNumber))
			return fmt.Errorf("card number %s is too short", cardNumber)
		}

		topupNo := fmt.Sprintf("TOPUP-%s", cardNumber[len(cardNumber)-5:])

		request := db.CreateTopupParams{
			CardNumber:  cardNumber,
			TopupNo:     topupNo,
			TopupAmount: int32(rand.Intn(10000000) + 1000000),
			TopupMethod: topupMethods[rand.Intn(len(topupMethods))],
		}

		_, err := r.db.CreateTopup(r.ctx, request)
		if err != nil {
			r.logger.Error("failed to seed topup for card", zap.String("card", cardNumber), zap.Error(err))
			return fmt.Errorf("failed to seed topup for card %s: %w", cardNumber, err)
		}
	}

	r.logger.Info("topup seeded successfully")

	return nil
}
