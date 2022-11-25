package domain

import "time"

type BookWithListing struct {
	Book    *Book          `json:"book"`
	Listing *AmazonListing `json:"listing"`
}

type ShippingCost struct {
	AmountInCents int64  `json:"amountInCents"`
	MinDays       int64  `json:"minDays"`
	MaxDays       int64  `json:"maxDays"`
	Name          string `json:"name"`
}

type Receipt struct {
	OrderID           string            `json:"orderID"`
	OrderedOn         time.Time         `json:"orderedOn"`
	Books             []BookWithListing `json:"books"`
	Shipping          Address           `json:"shipping"`
	ShippingCost      ShippingCost      `json:"shippingCost"`
	Billing           Address           `json:"billing"`
	Customer          Customer          `json:"customer"`
	TotalInCents      int64             `json:"totalInCents"`
	PaymentType       string            `json:"paymentType"`
	PaymentIdentifier string            `json:"paymentIdentifier"`
}

type Customer struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type Address struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}
