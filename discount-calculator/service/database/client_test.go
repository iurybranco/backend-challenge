package database

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestGettingExistentUser(t *testing.T) {
	RegisterTestingT(t)
	client, err := New(MockConfig())
	Expect(err).Should(BeNil())
	user, err := client.GetUser(1)
	Expect(err).Should(BeNil())
	Expect(user.Id).Should(BeEquivalentTo(1))
	Expect(user.FirstName).Should(BeEquivalentTo("iury"))
	Expect(user.LastName).Should(BeEquivalentTo("branco"))
	Expect(user.DateOfBirth).Should(BeEquivalentTo(time.Date(1997, 06, 28, 12, 00, 00, 00, time.UTC)))
}

func TestGettingANonexistentUser(t *testing.T) {
	RegisterTestingT(t)
	client, err := New(MockConfig())
	Expect(err).Should(BeNil())
	user, err := client.GetUser(5)
	Expect(err).ShouldNot(BeNil())
	Expect(err).Should(BeEquivalentTo(ErrNoDocuments))
	Expect(user).Should(BeNil())
}

func TestGettingAExistentProduct(t *testing.T) {
	RegisterTestingT(t)
	client, err := New(MockConfig())
	Expect(err).Should(BeNil())
	product, err := client.GetProduct(1)
	Expect(err).Should(BeNil())
	Expect(product.Id).Should(BeEquivalentTo(1))
	Expect(product.PriceInCents).Should(BeEquivalentTo(150))
	Expect(product.Title).Should(BeEquivalentTo("Soap"))
	Expect(product.Description).Should(BeEquivalentTo("A smelly soap to take a shower"))
	Expect(product.Discount.Percentage).Should(BeEquivalentTo(0))
	Expect(product.Discount.ValueInCents).Should(BeEquivalentTo(0))
}

func TestGettingANonexistentProduct(t *testing.T) {
	RegisterTestingT(t)
	client, err := New(MockConfig())
	Expect(err).Should(BeNil())
	product, err := client.GetProduct(5)
	Expect(err).ShouldNot(BeNil())
	Expect(err).Should(BeEquivalentTo(ErrNoDocuments))
	Expect(product).Should(BeNil())
}

func MockConfig() Config {
	return Config{
		Host:              "localhost",
		Port:              27017,
		Database:          "challenge",
		UserCollection:    "user",
		ProductCollection: "product",
		Username:          "root",
		Password:          "dummyPass",
	}
}
