package discount

import (
	"context"
	"fmt"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/controller"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cntroller controller.Controller
}

func NewDiscountServer(cntroller controller.Controller) *Server {
	return &Server{cntroller: cntroller}
}

// GRPC SERVICE TO CALCULATE A PRODUCT DISCOUNT
func (s *Server) Calculate(ctx context.Context, req *Request) (*Response, error) {
	log.Info(fmt.Sprintf("new request to calculate discount of product %d to user %d", req.ProductId, req.UserId))
	discount, err := s.cntroller.Calculate(req.UserId, req.ProductId)
	if err != nil {
		log.Error(err)
		if err == controller.ErrUserNotFound || err == controller.ErrProductNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info(fmt.Sprintf("discount of product %d calculated successfully to user %d: %.2f%% i.e. R$%.2f", req.ProductId, req.UserId, discount.Percentage, float32(discount.ValueInCents)/100))
	return &Response{
		Percentage:   discount.Percentage,
		ValueInCents: discount.ValueInCents,
	}, nil
}
