package seeder

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type saldoSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewSaldoSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *saldoSeeder {
	return &saldoSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *saldoSeeder) Seed() error {
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
		r.logger.Error("no cards available to seed saldo")
		return fmt.Errorf("no cards available to seed saldo")
	}

	for _, card := range cards {
		request := db.CreateSaldoParams{
			CardNumber:   card.CardNumber,
			TotalBalance: int32(rand.Intn(10000000) + 1000000),
		}

		_, err := r.db.CreateSaldo(r.ctx, request)
		if err != nil {
			r.logger.Error("failed to seed saldo for card", zap.String("card", card.CardNumber), zap.Error(err))

			return fmt.Errorf("failed to seed saldo for card %s: %w", card.CardNumber, err)
		}
	}

	r.logger.Info("saldo seeded successfully")

	return nil
}
