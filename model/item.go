package model

type Item struct {
	Id                int64       `json:"id"`
	Title             string      `json:"title"`
	SellerId          int64       `json:"sellerId"`
	Description       Description `json:"description"`
	RemainingQuantity int64       `json:"remainingQuantity"`
	SoldQuantity      int64       `json:"soldQuantity"`
	Price             float32     `json:"price"`
}

type Description struct {
	PlainText string `json:"plaintext"`
	HTML      string `json:"html"`
}
