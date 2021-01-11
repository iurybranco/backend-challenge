package main

import (
	"encoding/json"
	"fmt"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database/documents"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

//IT REQUIRES ALL CONTAINERS UP
func TestIntegration(t *testing.T) {
	RegisterTestingT(t)
	user := getMockUser()
	mongoClient := getMongoClient()
	createUser(mongoClient, user)
	defer deleteUser(mongoClient, user.Id)
	t.Run("It calculates a discount to a user that haven't", func(t *testing.T) {
		resp := requestProductsToUser(1)
		Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body := getResponseBody(resp.Body)
		productList := unmarshallProductList(body)
		for _, value := range productList {
			assertProduct(value, false)
		}
	})
	t.Run("It calculates a birthday discount to a user", func(t *testing.T) {
		resp := requestProductsToUser(user.Id)
		Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body := getResponseBody(resp.Body)
		productList := unmarshallProductList(body)
		for _, value := range productList {
			assertProduct(value, true)
		}
	})
	t.Run("It calculates a discount to a user nonexistent user", func(t *testing.T) {
		resp := requestProductsToUser(13)
		Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body := getResponseBody(resp.Body)
		productList := unmarshallProductList(body)
		for _, value := range productList {
			assertProduct(value, false)
		}
	})
}

func unmarshallProductList(body []byte) []documents.Product {
	var jsend Jsend
	var productList []documents.Product
	jsend.Data = &productList
	err := json.Unmarshal(body, &jsend)
	Expect(err).ShouldNot(HaveOccurred())
	return productList
}

func assertProduct(product documents.Product, withDiscount bool) {
	switch product.Id {
	case 1:
		Expect(product.Title).Should(BeEquivalentTo("Soap"))
		Expect(product.Description).Should(BeEquivalentTo("A smelly soap to take a shower"))
		Expect(product.PriceInCents).Should(BeEquivalentTo(150))
		if withDiscount {
			Expect(product.Discount.Percentage).Should(BeEquivalentTo(05.00))
			Expect(product.Discount.ValueInCents).Should(BeEquivalentTo(7))
			return
		}
		Expect(product.Discount.Percentage).Should(BeZero())
		Expect(product.Discount.ValueInCents).Should(BeZero())
	case 2:
		Expect(product.Title).Should(BeEquivalentTo("Soda"))
		Expect(product.Description).Should(BeEquivalentTo("The best soda ever"))
		Expect(product.PriceInCents).Should(BeEquivalentTo(700))
		if withDiscount {
			Expect(product.Discount.Percentage).Should(BeEquivalentTo(05.00))
			Expect(product.Discount.ValueInCents).Should(BeEquivalentTo(35))
			return
		}
		Expect(product.Discount.Percentage).Should(BeZero())
		Expect(product.Discount.ValueInCents).Should(BeZero())
	}
}

func requestProductsToUser(id int) *http.Response {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3001/product?userId=%d", id))
	Expect(err).ShouldNot(HaveOccurred())
	return resp
}

func getResponseBody(reader io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(reader)
	Expect(err).ShouldNot(HaveOccurred())
	return body
}

func getMongoClient() *mongo.Client {
	dbClient, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", "localhost", 27017)).
		SetAuth(options.Credential{
			Username: "root",
			Password: "dummyPass",
		}),
	)
	Expect(err).Should(BeNil())
	return dbClient
}

func createUser(mongoClient *mongo.Client, user documents.User) {
	_, err := mongoClient.Database("challenge").Collection("user").InsertOne(nil, user)
	Expect(err).Should(BeNil())
}

func deleteUser(mongoClient *mongo.Client, id int) {
	_, err := mongoClient.Database("challenge").Collection("user").DeleteOne(nil, bson.M{"_id": id})
	Expect(err).Should(BeNil())
}

func getMockUser() documents.User {
	return documents.User{
		Id:          4,
		FirstName:   "alguem",
		LastName:    "silva",
		DateOfBirth: time.Now().UTC(),
	}
}

type Jsend struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
