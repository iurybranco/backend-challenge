package controller

import (
	"github.com/iurybranco/backend-challenge/discount-calculator/service/clock"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	"github.com/pkg/errors"
	"time"
)

var ErrUserNotFound = errors.New("user not found")
var ErrProductNotFound = errors.New("product not found")

type controller struct {
	dbClient database.Client
}

func New(dbClient database.Client) Controller {
	return &controller{
		dbClient: dbClient,
	}
}

// IT CALCULATES A PRODUCT DISCOUNT
func (c *controller) Calculate(userId, productId int32) (*documents.Discount, error) {
	user, err := c.dbClient.GetUser(userId)
	if err != nil {
		if err == database.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "failed to get user from database")
	}
	product, err := c.dbClient.GetProduct(productId)
	if err != nil {
		if err == database.ErrNoDocuments {
			return nil, ErrProductNotFound
		}
		return nil, errors.Wrap(err, "failed to get product from database")
	}
	currentDate := clock.Now()
	if c.isBlackFriday(currentDate) {
		return NewDiscount(10.00, product.PriceInCents), nil
	}
	if c.isUserBirthday(currentDate, user) {
		return NewDiscount(05.00, product.PriceInCents), nil
	}
	return &product.Discount, nil
}

func (c *controller) isUserBirthday(currentDate time.Time, user *documents.User) bool {
	return user.DateOfBirth.Month() == currentDate.Month() && user.DateOfBirth.Day() == currentDate.Day()
}

func (c *controller) isBlackFriday(currentDate time.Time) bool {
	return currentDate.Month() == time.November && currentDate.Day() == 25
}

func NewDiscount(percentage float32, totalPrice int) *documents.Discount {
	return &documents.Discount{
		Percentage:   percentage,
		ValueInCents: int32(float32(totalPrice) * (percentage / 100)),
	}
}
