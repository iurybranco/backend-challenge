package documents

type Product struct {
	Id           int      `bson:"_id"`
	PriceInCents int      `bson:"price_in_cents"`
	Title        string   `bson:"title"`
	Description  string   `bson:"description"`
	Discount     Discount `bson:"discount"`
}

type Discount struct {
	Percentage   float32 `bson:"percentage"`
	ValueInCents int32   `bson:"value_in_cents"`
}
