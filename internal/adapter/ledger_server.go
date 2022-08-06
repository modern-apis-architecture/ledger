package adapter

import (
	"context"
	"github.com/modern-apis-architecture/ledger/internal/api"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LedgerServer struct {
	api.UnimplementedLedgerServiceServer
	svc *service.TransactionService
}

func NewLedgerServer(svc *service.TransactionService) *LedgerServer {
	return &LedgerServer{svc: svc}
}

func (ls *LedgerServer) Confirmation(context.Context, *api.RequestConfirmation) (*api.TransactionResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
}
func (ls *LedgerServer) Cancellation(context.Context, *api.RequestCancellation) (*api.TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancellation not implemented")
}
func (ls *LedgerServer) Reversal(context.Context, *api.RequestReversal) (*api.TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reversal not implemented")
}
