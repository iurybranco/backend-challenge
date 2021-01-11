package documents

type Product struct {
	Id           int      `json:"_id" bson:"_id"`
	PriceInCents int      `json:"price_in_cents" bson:"price_in_cents"`
	Title        string   `bson:"title"`
	Description  string   `bson:"description"`
	Discount     Discount `bson:"discount"`
}

type Discount struct {
	Percentage   float32 `bson:"percentage"`
	ValueInCents int32   `json:"value_in_cents" bson:"value_in_cents"`
}
