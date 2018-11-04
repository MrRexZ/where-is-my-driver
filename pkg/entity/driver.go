package entity

type Driver struct {
	Id       string  `json:"id" bson:"id"`
	Lat      float64 `json:"latitude" bson:"latitude"`
	Long     float64 `json:"longitude" bson:"longitude"`
	Accuracy float64 `json:"accuracy" bson:"accuracy"`
}
