package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"strings"
	"sync"
)

type BestSellers interface {
	GetBestSellers(c *gin.Context)
	Popular(c *gin.Context)
}

type BestSellerResponse struct {
	Lists []NYTListWithGoogleBooks `json:"lists"`
}

type NYTListWithGoogleBooks struct {
	List  *domain.List   `json:"list"`
	Books []*domain.Book `json:"books"`
}

type ISBNVolume struct {
	ISBN     string
	VolumeID string
}

var (
	popularISBNs = []*ISBNVolume{
		{
			ISBN: "9780525463467", // my side of the mountain
		},
		{
			ISBN: "9780385056199", // where the red fern grows
		},
		{
			ISBN: "9780330351690", // into the wild
		},
		{
			ISBN: "9780060115456", // old yeller
		},
		{
			ISBN: "9780007136599", // the fellowship of the ring
		},
		{
			ISBN: "9780691014647", // walden
		},
		{
			ISBN: "9780061233326", // pilgrim at tinker creek
		},
	}
	campingISBNS = []*ISBNVolume{
		{
			ISBN: "9781510722606", // Camping and Woodcraft,
		},
		{
			ISBN: "9781680511307", // Dirty Gormet,
		},
		{
			ISBN:     "9781452176994", // The Pendleton Field Guide to Camping
			VolumeID: "T3TTDwAAQBAJ",
		},
		{
			ISBN:     "9781512434286", // Camp so-and-so
			VolumeID: "H8ulDQAAQBAJ",
		},
		{
			ISBN: "9781612129013", //Backpack Explorer: On the Nature Trail
		},
		{
			ISBN: "9781612388052", // Moon Oregon Camping
		},
		{
			ISBN: "9781635861525", // Wilderness Adventure Camp
		},
		{
			ISBN: "9781452103822", // A Camping Spree with Mr. Magee
		},
		{
			ISBN: "9781590309551", //The Down and Dirty Guide to Camping with Kids
		},
	}
)

func (a *APIHandler) booksForIdentifiers(identifieres []*ISBNVolume) []*domain.Book {
	books := make([]*domain.Book, len(identifieres))
	wg := sync.WaitGroup{}
	for i, identifier := range identifieres {
		i := i

		wg.Add(1)
		go func(identifier *ISBNVolume) {
			defer wg.Done()
			var book *domain.Book
			var err error
			if identifier.VolumeID != "" {
				book, err = a.google.GetVolumeByID(identifier.VolumeID, identifier.ISBN, 0)
			} else {
				book, err = a.google.GetISBN(identifier.ISBN, 0)
			}

			if err != nil {
				panic(err)
			}
			books[i] = book
		}(identifier)
	}
	wg.Wait()
	return books
}

func (a *APIHandler) Category(c *gin.Context) {
	cat := strings.ToLower(strings.TrimSpace(c.Param("category")))
	if cat == "" {
		c.JSON(500, gin.H{"error": "You must provide a category"})
		return
	}
	switch cat {
	case "popular":
		books := a.booksForIdentifiers(popularISBNs)
		c.JSON(200, NYTListWithGoogleBooks{
			List: &domain.List{
				ListID:          0,
				ListName:        "",
				ListNameEncoded: "",
				DisplayName:     "Popular",
				Description:     "Our best selling collection of books imbued with the smell of campfire.  From camping to fantasy, this is what customers love the most.",
				Updated:         "",
				ListImage:       nil,
				ListImageWidth:  nil,
				ListImageHeight: nil,
				Books:           nil,
			},
			Books: books,
		})
		return
	case "camping":
		books := a.booksForIdentifiers(campingISBNS)
		c.JSON(200, NYTListWithGoogleBooks{
			List: &domain.List{
				ListID:          0,
				ListName:        "",
				ListNameEncoded: "",
				DisplayName:     "Camping",
				Description:     "Nothing like reading by the campfire in the woods, except when you can do it from home!  These books are oozing with that camping aesthetic and now with the smell of campfire itself.",
				Updated:         "",
				ListImage:       nil,
				ListImageWidth:  nil,
				ListImageHeight: nil,
				Books:           nil,
			},
			Books: books,
		})
		return
	}

	res, err := a.nyt.GetCategory(cat)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	ISBNs := []*ISBNVolume{}
	for _, book := range res.Results.Books {
		ISBNs = append(ISBNs, &ISBNVolume{ISBN: book.PrimaryIsbn13})
	}
	books := a.booksForIdentifiers(ISBNs)

	response := &NYTListWithGoogleBooks{List: &res.Results, Books: books}
	c.JSON(200, response)
}
