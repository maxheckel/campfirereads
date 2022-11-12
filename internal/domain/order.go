package domain

import "time"

type BookWithListing struct {
	Book    *Book          `json:"book"`
	Listing *AmazonListing `json:"listing"`
}

type Receipt struct {
	OrderedOn time.Time
	Books     []BookWithListing `json:"books"`
	Shipping  Address           `json:"shipping"`
	Customer  Customer
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
