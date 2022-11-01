package domain

type AllListsBestSellers struct {
	Status     string  `json:"status"`
	Copyright  string  `json:"copyright"`
	NumResults int     `json:"num_results"`
	Results    Results `json:"results"`
}
type BuyLinks struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Books struct {
	AgeGroup           string     `json:"age_group"`
	AmazonProductURL   string     `json:"amazon_product_url"`
	ArticleChapterLink string     `json:"article_chapter_link"`
	Author             string     `json:"author"`
	BookImage          string     `json:"book_image"`
	BookImageWidth     int        `json:"book_image_width"`
	BookImageHeight    int        `json:"book_image_height"`
	BookReviewLink     string     `json:"book_review_link"`
	BookURI            string     `json:"book_uri"`
	Contributor        string     `json:"contributor"`
	ContributorNote    string     `json:"contributor_note"`
	CreatedDate        string     `json:"created_date"`
	Description        string     `json:"description"`
	FirstChapterLink   string     `json:"first_chapter_link"`
	Price              string     `json:"price"`
	PrimaryIsbn10      string     `json:"primary_isbn10"`
	PrimaryIsbn13      string     `json:"primary_isbn13"`
	Publisher          string     `json:"publisher"`
	Rank               int        `json:"rank"`
	RankLastWeek       int        `json:"rank_last_week"`
	SundayReviewLink   string     `json:"sunday_review_link"`
	Title              string     `json:"title"`
	UpdatedDate        string     `json:"updated_date"`
	WeeksOnList        int        `json:"weeks_on_list"`
	BuyLinks           []BuyLinks `json:"buy_links"`
}
type Lists struct {
	ListID          int         `json:"list_id"`
	ListName        string      `json:"list_name"`
	ListNameEncoded string      `json:"list_name_encoded"`
	DisplayName     string      `json:"display_name"`
	Updated         string      `json:"updated"`
	ListImage       interface{} `json:"list_image"`
	ListImageWidth  interface{} `json:"list_image_width"`
	ListImageHeight interface{} `json:"list_image_height"`
	Books           []Books     `json:"books"`
}
type Results struct {
	BestsellersDate          string  `json:"bestsellers_date"`
	PublishedDate            string  `json:"published_date"`
	PublishedDateDescription string  `json:"published_date_description"`
	PreviousPublishedDate    string  `json:"previous_published_date"`
	NextPublishedDate        string  `json:"next_published_date"`
	Lists                    []Lists `json:"lists"`
}
