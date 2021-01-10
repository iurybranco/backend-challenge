package discount

import (
	"context"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/controller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
	cntroller controller.Controller
}

func NewDiscountServer(cntroller controller.Controller) *Server {
	return &Server{cntroller: cntroller}
}

// GRPC SERVICE TO CALCULATE A PRODUCT DISCOUNT
func (s *Server) Calculate(ctx context.Context, req *Request) (*Response, error) {
	discount, err := s.cntroller.Calculate(time.Now(), req.UserId, req.ProductId)
	if err != nil {
		if err == controller.ErrUserNotFound || err == controller.ErrProductNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &Response{
		Percentage:   discount.Percentage,
		ValueInCents: discount.ValueInCents,
	}, nil
}
