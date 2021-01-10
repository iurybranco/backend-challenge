package database

import "github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"

type Client interface {
	GetUser(id int32) (*documents.User, error)
	GetProduct(id int32) (*documents.Product, error)
	Close() error
}
