package domain

type BookSearchResult struct {
	Kind       string   `json:"kind"`
	TotalItems int      `json:"totalItems"`
	Items      []Result `json:"items"`
}

type BookSearch struct {
	Query string
}

type Result struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo struct {
		Country                string `json:"country"`
		Viewability            string `json:"viewability"`
		Embeddable             bool   `json:"embeddable"`
		PublicDomain           bool   `json:"publicDomain"`
		TextToSpeechPermission string `json:"textToSpeechPermission"`
		Epub                   struct {
			IsAvailable  bool   `json:"isAvailable"`
			AcsTokenLink string `json:"acsTokenLink,omitempty"`
		} `json:"epub"`
		Pdf struct {
			IsAvailable  bool   `json:"isAvailable"`
			AcsTokenLink string `json:"acsTokenLink,omitempty"`
		} `json:"pdf"`
		WebReaderLink       string `json:"webReaderLink"`
		AccessViewStatus    string `json:"accessViewStatus"`
		QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
	} `json:"accessInfo"`
	SearchInfo struct {
		TextSnippet string `json:"textSnippet"`
	} `json:"searchInfo"`
}

type VolumeInfo struct {
	Title               string   `json:"title"`
	Subtitle            string   `json:"subtitle,omitempty"`
	Authors             []string `json:"authors"`
	Publisher           string   `json:"publisher,omitempty"`
	PublishedDate       string   `json:"publishedDate"`
	Description         string   `json:"description"`
	IndustryIdentifiers []struct {
		Type       string `json:"type"`
		Identifier string `json:"identifier"`
	} `json:"industryIdentifiers"`
	ReadingModes struct {
		Text  bool `json:"text"`
		Image bool `json:"image"`
	} `json:"readingModes"`
	PageCount           int      `json:"pageCount"`
	PrintType           string   `json:"printType"`
	Categories          []string `json:"categories"`
	AverageRating       float64  `json:"averageRating,omitempty"`
	RatingsCount        int      `json:"ratingsCount,omitempty"`
	MaturityRating      string   `json:"maturityRating"`
	AllowAnonLogging    bool     `json:"allowAnonLogging"`
	ContentVersion      string   `json:"contentVersion"`
	PanelizationSummary struct {
		ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
		ContainsImageBubbles bool `json:"containsImageBubbles"`
	} `json:"panelizationSummary,omitempty"`
	ImageLinks struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
	Language            string `json:"language"`
	PreviewLink         string `json:"previewLink"`
	InfoLink            string `json:"infoLink"`
	CanonicalVolumeLink string `json:"canonicalVolumeLink"`
}

type SaleInfo struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
	ListPrice   struct {
		Amount       float64 `json:"amount"`
		CurrencyCode string  `json:"currencyCode"`
	} `json:"listPrice,omitempty"`
	RetailPrice struct {
		Amount       float64 `json:"amount"`
		CurrencyCode string  `json:"currencyCode"`
	} `json:"retailPrice,omitempty"`
	BuyLink string `json:"buyLink,omitempty"`
	Offers  []struct {
		FinskyOfferType int `json:"finskyOfferType"`
		ListPrice       struct {
			AmountInMicros int    `json:"amountInMicros"`
			CurrencyCode   string `json:"currencyCode"`
		} `json:"listPrice"`
		RetailPrice struct {
			AmountInMicros int    `json:"amountInMicros"`
			CurrencyCode   string `json:"currencyCode"`
		} `json:"retailPrice"`
		Giftable bool `json:"giftable"`
	} `json:"offers,omitempty"`
}

type SearchResult struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title  string `json:"title"`
		Facets [][]struct {
			Anchor      string `json:"anchor"`
			Label       string `json:"label"`
			LabelWithOp string `json:"label_with_op"`
		} `json:"facets"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		CacheID          string `json:"cacheId"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Pagemap          struct {
			CseThumbnail []struct {
				Src    string `json:"src"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"cse_thumbnail"`
			Metatags []struct {
				Viewport string `json:"viewport"`
			} `json:"metatags"`
			CseImage []struct {
				Src string `json:"src"`
			} `json:"cse_image"`
		} `json:"pagemap,omitempty"`
		Labels []struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			LabelWithOp string `json:"label_with_op"`
		} `json:"labels,omitempty"`
	} `json:"items"`
}
