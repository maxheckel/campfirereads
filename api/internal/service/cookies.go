package service

import (
	"fmt"
	"net/http"
)

type AmazonSessionDetails struct {
	SessionID   string
	SessionTime string
	UBID        string
}

var AmazonSession = &AmazonSessionDetails{}

func (a *AmazonSessionDetails) String() string {
	if a.SessionID == "" {
		err := a.EstablishSession()
		if err != nil {
			panic(err)
		}
	}
	return fmt.Sprintf("i18n-prefs=USD; session-id=%s; session-id-time=%s; skin=noskin; ubid-main=%s", a.SessionID, a.SessionTime, a.UBID)
}

func (a *AmazonSessionDetails) EstablishSession() error {
	fmt.Println("Establishing session")
	req, _ := http.NewRequest("GET", "https://amazon.com", nil)
	req.Header.Set("User-Agent", ":Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	client := http.Client{}
	html, err := client.Do(req)
	if err != nil {
		return err
	}
	a.SetSessionDetailsFromResponse(html)
	return nil
}

func (a *AmazonSessionDetails) SetSessionDetailsFromResponse(html *http.Response) {
	for _, c := range html.Cookies() {
		switch c.Name {
		case "session-id":
			AmazonSession.SessionID = c.Value
		case "session-id-time":
			AmazonSession.SessionTime = c.Value
		case "ubid-main":
			AmazonSession.UBID = c.Value
		}
	}
}
