package protomapper

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
)

type cardProtoMapper struct {
}

func NewCardProtoMapper() *cardProtoMapper {
	return &cardProtoMapper{}
}

func (s *cardProtoMapper) ToResponseCard(card *response.CardResponse) *pb.CardResponse {
	return &pb.CardResponse{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

func (s *cardProtoMapper) ToResponsesCard(cards []*response.CardResponse) []*pb.CardResponse {
	responses := make([]*pb.CardResponse, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseCard(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseCardDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt {
	return &pb.CardResponseDeleteAt{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    card.DeletedAt,
	}
}

func (s *cardProtoMapper) ToResponsesCardDeletedAt(cards []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt {
	responses := make([]*pb.CardResponseDeleteAt, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseCardDeleteAt(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseDashboardCard(dash *response.DashboardCard) *pb.CardResponseDashboard {
	return &pb.CardResponseDashboard{
		TotalBalance:     *dash.TotalBalance,
		TotalWithdraw:    *dash.TotalWithdraw,
		TotalTopup:       *dash.TotalTopup,
		TotalTransfer:    *dash.TotalTransfer,
		TotalTransaction: *dash.TotalTransaction,
	}
}

func (s *cardProtoMapper) ToResponseDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber {
	return &pb.CardResponseDashboardCardNumber{
		TotalBalance:          *dash.TotalBalance,
		TotalWithdraw:         *dash.TotalWithdraw,
		TotalTopup:            *dash.TotalTopup,
		TotalTransferSend:     *dash.TotalTransferSent,
		TotalTransferReceiver: *dash.TotalTransferReceiver,
		TotalTransaction:      *dash.TotalTransaction,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyBalance(cards *response.CardResponseMonthBalance) *pb.CardResponseMonthlyBalance {
	return &pb.CardResponseMonthlyBalance{
		Month:        cards.Month,
		TotalBalance: cards.TotalBalance,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyBalances(cards []*response.CardResponseMonthBalance) []*pb.CardResponseMonthlyBalance {
	responses := make([]*pb.CardResponseMonthlyBalance, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyBalance(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyBalance(cards *response.CardResponseYearlyBalance) *pb.CardResponseYearlyBalance {
	return &pb.CardResponseYearlyBalance{
		Year:         cards.Year,
		TotalBalance: cards.TotalBalance,
	}
}

func (s *cardProtoMapper) ToResponseYearlyBalances(cards []*response.CardResponseYearlyBalance) []*pb.CardResponseYearlyBalance {
	responses := make([]*pb.CardResponseYearlyBalance, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyBalance(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseMonthlyTopupAmount(cards *response.CardResponseMonthTopupAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyTopupAmounts(cards []*response.CardResponseMonthTopupAmount) []*pb.CardResponseMonthlyAmount {
	responses := make([]*pb.CardResponseMonthlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyTopupAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyTopupAmount(cards *response.CardResponseYearlyTopupAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseYearlyTopupAmounts(cards []*response.CardResponseYearlyTopupAmount) []*pb.CardResponseYearlyAmount {
	responses := make([]*pb.CardResponseYearlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyTopupAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseMonthlyWithdrawAmount(cards *response.CardResponseMonthWithdrawAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyWithdrawAmounts(cards []*response.CardResponseMonthWithdrawAmount) []*pb.CardResponseMonthlyAmount {
	responses := make([]*pb.CardResponseMonthlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyWithdrawAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyWithdrawAmount(cards *response.CardResponseYearlyWithdrawAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseYearlyWithdrawAmounts(cards []*response.CardResponseYearlyWithdrawAmount) []*pb.CardResponseYearlyAmount {
	responses := make([]*pb.CardResponseYearlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyWithdrawAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseMonthlyTransactionAmount(cards *response.CardResponseMonthTransactionAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyTransactionAmounts(cards []*response.CardResponseMonthTransactionAmount) []*pb.CardResponseMonthlyAmount {
	responses := make([]*pb.CardResponseMonthlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyTransactionAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyTransactionAmount(cards *response.CardResponseYearlyTransactionAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseYearlyTransactionAmounts(cards []*response.CardResponseYearlyTransactionAmount) []*pb.CardResponseYearlyAmount {
	responses := make([]*pb.CardResponseYearlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyTransactionAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseMonthlyTransferSenderAmount(cards *response.CardResponseMonthTransferAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyTransferSenderAmounts(cards []*response.CardResponseMonthTransferAmount) []*pb.CardResponseMonthlyAmount {
	responses := make([]*pb.CardResponseMonthlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyTransferSenderAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyTransferSenderAmount(cards *response.CardResponseYearlyTransferAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseYearlyTransferSenderAmounts(cards []*response.CardResponseYearlyTransferAmount) []*pb.CardResponseYearlyAmount {
	responses := make([]*pb.CardResponseYearlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyTransferSenderAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseMonthlyTransferReceiverAmount(cards *response.CardResponseMonthTransferAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseMonthlyTransferReceiverAmounts(cards []*response.CardResponseMonthTransferAmount) []*pb.CardResponseMonthlyAmount {
	responses := make([]*pb.CardResponseMonthlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseMonthlyTransferReceiverAmount(card))
	}
	return responses
}

func (s *cardProtoMapper) ToResponseYearlyTransferReceiverAmount(cards *response.CardResponseYearlyTransferAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

func (s *cardProtoMapper) ToResponseYearlyTransferReceiverAmounts(cards []*response.CardResponseYearlyTransferAmount) []*pb.CardResponseYearlyAmount {
	responses := make([]*pb.CardResponseYearlyAmount, 0, len(cards))
	for _, card := range cards {
		responses = append(responses, s.ToResponseYearlyTransferReceiverAmount(card))
	}
	return responses
}
