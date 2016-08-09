package users

type Card struct {
	LongNum string `json:"longNum" bson:"longNum"`
	Expires string `json:"expires" bson:"expires"`
	CCV     string `json:"ccv" bson:"ccv"`
}
