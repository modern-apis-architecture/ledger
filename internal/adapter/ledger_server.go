package adapter

import (
	"context"
	"github.com/google/uuid"
	"github.com/modern-apis-architecture/ledger/internal/api"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type LedgerServer struct {
	api.UnimplementedLedgerServiceServer
	svc *service.TransactionService
}

func NewLedgerServer(svc *service.TransactionService) *LedgerServer {
	return &LedgerServer{svc: svc}
}

func (ls *LedgerServer) Confirmation(ctx context.Context, rc *api.RequestConfirmation) (*api.TransactionResponse, error) {
	id, _ := uuid.NewUUID()
	t := &transaction.Transaction{
		Id:           id.String(),
		LocalTime:    time.Now(),
		RegisteredAt: rc.Transaction.LocalTime.AsTime(),
		CardId:       rc.Transaction.CardId,
		Data: transaction.TransactionData{
			ExternalId: rc.Transaction.ExternalId,
			Value:      rc.Transaction.Value,
			MerchantId: rc.Transaction.MerchantId,
		},
		Type: rc.Type.String(),
	}
	tr, err := ls.svc.Confirmation(t)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
	}
	rsp := &api.TransactionResponse{
		RegisterId: tr.Id,
		Result:     0,
	}
	return rsp, nil
}
func (ls *LedgerServer) Cancellation(ctx context.Context, rc *api.RequestCancellation) (*api.TransactionResponse, error) {
	id, _ := uuid.NewUUID()
	t := &transaction.Transaction{
		Id:           id.String(),
		LocalTime:    time.Now(),
		RegisteredAt: rc.Transaction.LocalTime.AsTime(),
		CardId:       rc.Transaction.CardId,
		Data: transaction.TransactionData{
			ExternalId: rc.Transaction.ExternalId,
			Value:      rc.Transaction.Value,
			MerchantId: rc.Transaction.MerchantId,
		},
		Type: rc.Type.String(),
	}
	tr, err := ls.svc.Cancellation(t)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
	}
	rsp := &api.TransactionResponse{
		RegisterId: tr.Id,
		Result:     0,
	}
	return rsp, nil
}
func (ls *LedgerServer) Reversal(ctx context.Context, rr *api.RequestReversal) (*api.TransactionResponse, error) {
	id, _ := uuid.NewUUID()
	t := &transaction.Transaction{
		Id:           id.String(),
		LocalTime:    time.Now(),
		RegisteredAt: rr.Transaction.LocalTime.AsTime(),
		CardId:       rr.Transaction.CardId,
		Data: transaction.TransactionData{
			ExternalId: rr.Transaction.ExternalId,
			Value:      rr.Transaction.Value,
			MerchantId: rr.Transaction.MerchantId,
		},
		Type: rr.Type.String(),
	}
	tr, err := ls.svc.Reversal(t)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
	}
	rsp := &api.TransactionResponse{
		RegisterId: tr.Id,
		Result:     0,
	}
	return rsp, nil
}
