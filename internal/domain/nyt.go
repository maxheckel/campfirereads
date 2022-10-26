package domain

type BestSellers struct {
	Status     string       `json:"status"`
	Copyright  string       `json:"copyright"`
	NumResults int          `json:"num_results"`
	Results    []BestSeller `json:"results"`
}
type Isbns struct {
	Isbn10 string `json:"isbn10"`
	Isbn13 string `json:"isbn13"`
}
type RanksHistory struct {
	PrimaryIsbn10   string `json:"primary_isbn10"`
	PrimaryIsbn13   string `json:"primary_isbn13"`
	Rank            int    `json:"rank"`
	ListName        string `json:"list_name"`
	DisplayName     string `json:"display_name"`
	PublishedDate   string `json:"published_date"`
	BestsellersDate string `json:"bestsellers_date"`
	WeeksOnList     int    `json:"weeks_on_list"`
	RankLastWeek    int    `json:"rank_last_week"`
	Asterisk        int    `json:"asterisk"`
	Dagger          int    `json:"dagger"`
}
type Reviews struct {
	BookReviewLink     string `json:"book_review_link"`
	FirstChapterLink   string `json:"first_chapter_link"`
	SundayReviewLink   string `json:"sunday_review_link"`
	ArticleChapterLink string `json:"article_chapter_link"`
}
type BestSeller struct {
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Contributor     string         `json:"contributor"`
	Author          string         `json:"author"`
	ContributorNote string         `json:"contributor_note"`
	Price           string         `json:"price"`
	AgeGroup        string         `json:"age_group"`
	Publisher       string         `json:"publisher"`
	Isbns           []Isbns        `json:"isbns"`
	RanksHistory    []RanksHistory `json:"ranks_history"`
	Reviews         []Reviews      `json:"reviews"`
}
