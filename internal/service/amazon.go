package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Amazon interface {
	ISBNToListings(ISBN string) ([]*domain.AmazonListing, error)
	ListingToPriceInCents(listing *domain.AmazonListing) error
	ISBNToPrices(ISBN string) ([]*domain.AmazonListing, error)
}

type amazon struct {
	bookLinksText []string
	cache         cache.Cache
}

func NewAmazon(cache cache.Cache) Amazon {
	return &amazon{
		cache: cache,
		bookLinksText: []string{
			"paperback",
			"hardcover",
		},
	}
}

func (a amazon) ListingToPriceInCents(listing *domain.AmazonListing) error {
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
		if len(prices) == 1 {

		}
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
	if currentListPrice == 0 {
		doc.Find("#price").Each(func(i int, selection *goquery.Selection) {
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
	}
	if currentListPrice == 0 {
		currentListPrice = -1
	}
	listing.PriceInCents = int32(currentListPrice)
	return err
}

func (a amazon) ISBNToListings(ISBN string) ([]*domain.AmazonListing, error) {
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

	listings := []*domain.AmazonListing{}

	doc.Find("a.a-size-base").Each(func(i int, selection *goquery.Selection) {
		for _, textToFind := range a.bookLinksText {
			if strings.ToLower(selection.Text()) != textToFind && !listingExists(strings.ToLower(selection.Text()), listings) {
				continue
			}

			listing := &domain.AmazonListing{}
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
	actualListings := []*domain.AmazonListing{}
	for _, listing := range listings {
		if !listingExists(listing.Type, actualListings) {
			actualListings = append(actualListings, listing)
		}
	}

	return actualListings, err
}

func (a amazon) ISBNToPrices(ISBN string) ([]*domain.AmazonListing, error) {
	cacheKey := fmt.Sprintf("listings-%s", ISBN)
	cacheVal, err := a.cache.Read(cacheKey)
	if err != nil {
		return nil, err
	}
	if listings, ok := cacheVal.(*domain.AmazonListings); ok {
		return listings.Listings, nil
	}
	res, err := a.ISBNToListings(ISBN)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for index := range res {
		wg.Add(1)
		go func(index int) {
			err := a.ListingToPriceInCents(res[index])
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(index)
	}
	wg.Wait()
	err = a.cache.Write(cacheKey, &domain.AmazonListings{
		Listings: res,
	}, 24*60*60)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func listingExists(listingType string, listings []*domain.AmazonListing) bool {
	for _, listing := range listings {
		if listing.Type == listingType {
			return true
		}
	}
	return false
}
