package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AmazonListing struct {
	URL          *url.URL `json:"url"`
	ISBN         string   `json:"isbn"`
	Type         string   `json:"type"`
	PriceInCents int32    `json:"price_in_cents"`
	CrawlDate    time.Time
}

type Amazon interface {
	ISBNToListings(ISBN string) ([]*AmazonListing, error)
	ListingToPriceInCents(listing *AmazonListing) error
}

type amazon struct {
	bookLinksText []string
}

func NewAmazon() Amazon {
	return &amazon{
		bookLinksText: []string{
			"paperback",
			"hardcover",
		},
	}
}

func (a amazon) ListingToPriceInCents(listing *AmazonListing) error {
	req, err := http.NewRequest("GET", listing.URL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "CampfireReads")
	req.Header.Add("Cookie", AmazonSession.String())
	if err != nil {
		return err
	}
	client := http.Client{}
	html, err := client.Do(req)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil {
		log.Fatal(err)
	}
	currentListPrice := 0
	var cents int
	doc.Find("*[data-a-color|=\"price\"]").Each(func(i int, selection *goquery.Selection) {
		prices := strings.Split(selection.Text(), "$")
		for _, price := range prices {
			if price == "" || strings.Count(price, ".") > 1 {
				continue
			}
			cents, err = strconv.Atoi(strings.ReplaceAll(price, ".", ""))
			if cents > currentListPrice {
				currentListPrice = cents
			}
		}
	})
	listing.PriceInCents = int32(currentListPrice)
	return err
}

func (a amazon) ISBNToListings(ISBN string) ([]*AmazonListing, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.amazon.com/s?k=%s&i=stripbooks", ISBN), nil)
	req.Header.Add("User-Agent", "CampfireReads")
	req.Header.Add("Cookie", AmazonSession.String())
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	html, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	AmazonSession.SetSessionDetailsFromResponse(html)

	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil {
		log.Fatal(err)
	}

	listings := []*AmazonListing{}

	doc.Find("a.a-size-base").Each(func(i int, selection *goquery.Selection) {
		for _, textToFind := range a.bookLinksText {
			if strings.ToLower(selection.Text()) != textToFind && !listingExists(strings.ToLower(selection.Text()), listings) {
				continue
			}

			listing := &AmazonListing{}
			listing.CrawlDate = time.Now()
			href, _ := selection.Attr("href")
			listing.ISBN = ISBN
			listing.Type = strings.ToLower(selection.Text())
			listing.URL, err = url.Parse(fmt.Sprintf("https://amazon.com/%s", href))
			if err != nil {
				return
			}
			listings = append(listings, listing)
		}
	})
	actualListings := []*AmazonListing{}
	for _, listing := range listings {
		if !listingExists(listing.Type, actualListings) {
			actualListings = append(actualListings, listing)
		}
	}

	return actualListings, err
}

func listingExists(listingType string, listings []*AmazonListing) bool {
	for _, listing := range listings {
		if listing.Type == listingType {
			return true
		}
	}
	return false
}
