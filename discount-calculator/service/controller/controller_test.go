package controller

import (
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestCalculatingAndGettingBirthdayDiscount(t *testing.T) {
	RegisterTestingT(t)
	cntroller := New(&MockDbClient{
		getProductFunc: func(id int32) (*documents.Product, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.Product{
				Id:           1,
				PriceInCents: 1000,
				Title:        "product test",
				Description:  "description test",
				Discount: documents.Discount{
					Percentage:   0,
					ValueInCents: 0,
				},
			}, nil
		},
		getUserFunc: func(id int32) (*documents.User, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.User{
				Id:          1,
				FirstName:   "iury",
				LastName:    "branco",
				DateOfBirth: time.Now(),
			}, nil
		},
	})
	discount, err := cntroller.Calculate(time.Now(), 1, 1)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(*discount).Should(BeEquivalentTo(documents.Discount{
		Percentage:   05.00,
		ValueInCents: 50,
	}))
}

func TestCalculatingAndGettingBlackFridayDiscount(t *testing.T) {
	RegisterTestingT(t)
	cntroller := New(&MockDbClient{
		getProductFunc: func(id int32) (*documents.Product, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.Product{
				Id:           1,
				PriceInCents: 1000,
				Title:        "product test",
				Description:  "description test",
				Discount: documents.Discount{
					Percentage:   0,
					ValueInCents: 0,
				},
			}, nil
		},
		getUserFunc: func(id int32) (*documents.User, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.User{
				Id:          1,
				FirstName:   "iury",
				LastName:    "branco",
				DateOfBirth: time.Now(),
			}, nil
		},
	})
	blackFridayDate := time.Date(2021, 11, 25, 00, 00, 00, 00, time.UTC)
	discount, err := cntroller.Calculate(blackFridayDate, 1, 1)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(*discount).Should(BeEquivalentTo(documents.Discount{
		Percentage:   10.00,
		ValueInCents: 100,
	}))
}

func TestCalculatingAndGettingNoDiscount(t *testing.T) {
	RegisterTestingT(t)
	cntroller := New(&MockDbClient{
		getProductFunc: func(id int32) (*documents.Product, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.Product{
				Id:           1,
				PriceInCents: 1000,
				Title:        "product test",
				Description:  "description test",
				Discount: documents.Discount{
					Percentage:   0,
					ValueInCents: 0,
				},
			}, nil
		},
		getUserFunc: func(id int32) (*documents.User, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.User{
				Id:          1,
				FirstName:   "iury",
				LastName:    "branco",
				DateOfBirth: time.Date(2021, 11, 28, 00, 00, 00, 00, time.UTC),
			}, nil
		},
	})
	discount, err := cntroller.Calculate(time.Now(), 1, 1)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(*discount).Should(BeEquivalentTo(documents.Discount{
		Percentage:   00.00,
		ValueInCents: 0,
	}))
}

func TestCalculatingWithANotFoundProduct(t *testing.T) {
	RegisterTestingT(t)
	cntroller := New(&MockDbClient{
		getProductFunc: func(id int32) (*documents.Product, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return nil, database.ErrNoDocuments
		},
		getUserFunc: func(id int32) (*documents.User, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return &documents.User{
				Id:          1,
				FirstName:   "iury",
				LastName:    "branco",
				DateOfBirth: time.Date(2021, 11, 28, 00, 00, 00, 00, time.UTC),
			}, nil
		},
	})
	discount, err := cntroller.Calculate(time.Now(), 1, 1)
	Expect(err).Should(HaveOccurred())
	Expect(err).Should(BeEquivalentTo(ErrProductNotFound))
	Expect(discount).Should(BeNil())
}

func TestCalculatingWithANotFoundUser(t *testing.T) {
	RegisterTestingT(t)
	cntroller := New(&MockDbClient{
		getUserFunc: func(id int32) (*documents.User, error) {
			Expect(id).Should(BeEquivalentTo(1))
			return nil, database.ErrNoDocuments
		},
	})
	discount, err := cntroller.Calculate(time.Now(), 1, 1)
	Expect(err).Should(HaveOccurred())
	Expect(err).Should(BeEquivalentTo(ErrUserNotFound))
	Expect(discount).Should(BeNil())
}

type MockDbClient struct {
	getUserFunc    func(id int32) (*documents.User, error)
	getProductFunc func(id int32) (*documents.Product, error)
}

func (m *MockDbClient) GetUser(id int32) (*documents.User, error) {
	return m.getUserFunc(id)
}
func (m *MockDbClient) GetProduct(id int32) (*documents.Product, error) {
	return m.getProductFunc(id)
}
func (m *MockDbClient) Close() error {
	return nil
}
