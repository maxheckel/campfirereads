package domain

type BookWithListing struct {
	Book    *Book          `json:"book"`
	Listing *AmazonListing `json:"listing"`
}
