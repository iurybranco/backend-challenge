package controller

import (
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
)

type Controller interface {
	Calculate(userId, productId int32) (*documents.Discount, error)
}
