package server

import (
	"context"
	"errors"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/controller"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/server/discount"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterTestingT(t)
	cntroller := &MockController{}
	srv, err := New(4500, cntroller)
	Expect(err).ShouldNot(HaveOccurred())
	srv.Run()
	defer srv.Close()
	conn, err := grpc.Dial(":4500", grpc.WithInsecure(), grpc.WithBlock())
	Expect(err).ShouldNot(HaveOccurred())
	defer conn.Close()
	client := discount.NewDiscountClient(conn)
	t.Run("It calculates a discount", func(t *testing.T) {
		cntroller.calculateFunc = func(userId, productId int32) (*documents.Discount, error) {
			Expect(userId).Should(BeEquivalentTo(1))
			Expect(productId).Should(BeEquivalentTo(1))
			return &documents.Discount{
				Percentage:   10,
				ValueInCents: 100,
			}, nil
		}
		req := &discount.Request{
			ProductId: 1,
			UserId:    1,
		}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.Percentage).Should(BeEquivalentTo(10))
		Expect(resp.ValueInCents).Should(BeEquivalentTo(100))
	})
	t.Run("It calculates a discount with user not found", func(t *testing.T) {
		cntroller.calculateFunc = func(userId, productId int32) (*documents.Discount, error) {
			Expect(userId).Should(BeEquivalentTo(1))
			Expect(productId).Should(BeEquivalentTo(1))
			return nil, controller.ErrUserNotFound
		}
		req := &discount.Request{
			ProductId: 1,
			UserId:    1,
		}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).Should(HaveOccurred())
		Expect(resp).Should(BeNil())
	})
	t.Run("It calculates a discount with product not found", func(t *testing.T) {
		cntroller.calculateFunc = func(userId, productId int32) (*documents.Discount, error) {
			Expect(userId).Should(BeEquivalentTo(1))
			Expect(productId).Should(BeEquivalentTo(1))
			return nil, controller.ErrProductNotFound
		}
		req := &discount.Request{
			ProductId: 1,
			UserId:    1,
		}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).Should(HaveOccurred())
		Expect(resp).Should(BeNil())
	})
	t.Run("It calculates a discount with failure", func(t *testing.T) {
		cntroller.calculateFunc = func(userId, productId int32) (*documents.Discount, error) {
			Expect(userId).Should(BeEquivalentTo(1))
			Expect(productId).Should(BeEquivalentTo(1))
			return nil, errors.New("failed")
		}
		req := &discount.Request{
			ProductId: 1,
			UserId:    1,
		}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).Should(HaveOccurred())
		Expect(resp).Should(BeNil())
	})
}

type MockController struct {
	calculateFunc func(userId, productId int32) (*documents.Discount, error)
}

func (m *MockController) Calculate(userId, productId int32) (*documents.Discount, error) {
	return m.calculateFunc(userId, productId)
}
