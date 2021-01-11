package service

import (
	"context"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/clock"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/server/discount"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"strconv"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	RegisterTestingT(t)
	cfg := getMockConfig()
	servce, err := New(cfg)
	Expect(err).ShouldNot(HaveOccurred())
	servce.Run()
	defer servce.Shutdown()
	conn, err := grpc.Dial(":"+strconv.Itoa(cfg.ServerPort), grpc.WithInsecure(), grpc.WithBlock())
	Expect(err).ShouldNot(HaveOccurred())
	defer conn.Close()
	client := discount.NewDiscountClient(conn)
	t.Run("It calculates a discount to a user that haven't", func(t *testing.T) {
		req := &discount.Request{ProductId: 1, UserId: 1}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.Percentage).Should(BeZero())
		Expect(resp.ValueInCents).Should(BeZero())
	})
	t.Run("It calculates a black friday discount to a user", func(t *testing.T) {
		clock.SetFakeTimeNow(func() time.Time {
			blackFridayDate := time.Date(2021, 11, 25, 00, 00, 00, 00, time.UTC)
			return blackFridayDate
		})
		req := &discount.Request{ProductId: 1, UserId: 1}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.Percentage).Should(BeEquivalentTo(10.00))
		Expect(resp.ValueInCents).Should(BeEquivalentTo(15))
	})
	t.Run("It calculates a birthday discount to a user", func(t *testing.T) {
		clock.SetFakeTimeNow(func() time.Time {
			blackFridayDate := time.Date(2021, 06, 28, 00, 00, 00, 00, time.UTC)
			return blackFridayDate
		})
		req := &discount.Request{ProductId: 1, UserId: 1}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.Percentage).Should(BeEquivalentTo(05.00))
		Expect(resp.ValueInCents).Should(BeEquivalentTo(7))
	})
	t.Run("It calculates a birthday and black friday discount to a user", func(t *testing.T) {
		clock.SetFakeTimeNow(func() time.Time {
			blackFridayDate := time.Date(2021, 11, 25, 00, 00, 00, 00, time.UTC)
			return blackFridayDate
		})
		req := &discount.Request{ProductId: 1, UserId: 2}
		resp, err := client.Calculate(context.Background(), req)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.Percentage).Should(BeEquivalentTo(10.00))
		Expect(resp.ValueInCents).Should(BeEquivalentTo(15))
	})
}

func getMockConfig() *Config {
	return &Config{
		LogLevel: 6,
		Database: database.Config{
			Host:              "localhost",
			Port:              27017,
			Database:          "challenge",
			UserCollection:    "user",
			ProductCollection: "product",
			Username:          "root",
			Password:          "dummyPass",
		},
		ServerPort: 3006,
	}
}
