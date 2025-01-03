package seeder

import (
	db "MamangRust/paymentgatewaygrpc/pkg/database/schema"
	"MamangRust/paymentgatewaygrpc/pkg/logger"
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type transferSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewTransferSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *transferSeeder {
	return &transferSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *transferSeeder) Seed() error {
	cards, err := r.db.GetCards(r.ctx, db.GetCardsParams{
		Column1: "",
		Limit:   10,
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("failed to get card list", zap.Error(err))
		return fmt.Errorf("failed to get card list: %w", err)
	}

	if len(cards) < 2 {
		r.logger.Error("not enough cards available for transfer seeding")
		return fmt.Errorf("not enough cards available for transfer seeding")
	}

	for i := 0; i < 10; i++ {
		fromIndex := rand.Intn(len(cards))
		toIndex := rand.Intn(len(cards))

		for fromIndex == toIndex {
			toIndex = rand.Intn(len(cards))
		}

		transferFrom := cards[fromIndex].CardNumber
		transferTo := cards[toIndex].CardNumber

		transferAmount := int32(rand.Intn(1000000) + 50000)

		request := db.CreateTransferParams{
			TransferFrom:   transferFrom,
			TransferTo:     transferTo,
			TransferAmount: transferAmount,
		}

		_, err := r.db.CreateTransfer(r.ctx, request)
		if err != nil {
			r.logger.Error("failed to seed transfer", zap.Int("transfer", i+1), zap.Error(err))
			return fmt.Errorf("failed to seed transfer %d: %w", i+1, err)
		}
	}

	r.logger.Info("transfer seeded successfully")

	return nil
}
