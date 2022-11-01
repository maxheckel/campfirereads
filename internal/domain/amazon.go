package domain

import (
	"net/url"
	"time"
)

type AmazonListings struct {
	Listings []*AmazonListing `json:"listings"`
}

type AmazonListing struct {
	URL          *url.URL `json:"url"`
	ISBN         string   `json:"isbn"`
	Type         string   `json:"type"`
	PriceInCents int32    `json:"price_in_cents"`
	CrawlDate    time.Time
}
