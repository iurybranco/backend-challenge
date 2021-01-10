package controller

import (
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	"time"
)

type Controller interface {
	Calculate(currentDate time.Time, userId, productId int32) (*documents.Discount, error)
}
