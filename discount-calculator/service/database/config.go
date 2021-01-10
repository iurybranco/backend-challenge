package database

type Config struct {
	Host              string `json:"host"`
	Port              int    `json:"port"`
	Database          string `json:"database"`
	UserCollection    string `json:"userCollection"`
	ProductCollection string `json:"productCollection"`
	Username          string `json:"username"`
	Password          string `json:"password"`
}
