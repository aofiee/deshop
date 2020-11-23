package discount

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//API map[string]string
var API = map[string]string{
	"JP":    "https://ec.nintendo.com/api/JP/ja/search/sales?count=30&offset=0",
	"US":    "https://ec.nintendo.com/api/US/en/search/sales?count=30&offset=0",
	"GB":    "https://ec.nintendo.com/api/GB/en/search/sales?count=10&offset=0",
	"CA":    "https://ec.nintendo.com/api/CA/en/search/sales?count=30&offset=0#Canada",
	"AU":    "https://ec.nintendo.com/api/AU/en/search/sales?count=10&offset=0#Australia",
	"NZ":    "https://ec.nintendo.com/api/NZ/en/search/sales?count=10&offset=0#NewZealand",
	"CZ":    "https://ec.nintendo.com/api/CZ/en/search/sales?count=10&offset=0#Czech",
	"DK":    "https://ec.nintendo.com/api/DK/en/search/sales?count=10&offset=0#Denmark",
	"FI":    "https://ec.nintendo.com/api/FI/en/search/sales?count=10&offset=0#Finland",
	"GR":    "https://ec.nintendo.com/api/GR/en/search/sales?count=10&offset=0#Greece",
	"HU":    "https://ec.nintendo.com/api/HU/en/search/sales?count=10&offset=0#Hungary",
	"NO":    "https://ec.nintendo.com/api/NO/en/search/sales?count=10&offset=0#Norway",
	"PL":    "https://ec.nintendo.com/api/PL/en/search/sales?count=10&offset=0#Poland",
	"ZA":    "https://ec.nintendo.com/api/ZA/en/search/sales?count=10&offset=0#SouthAfrica",
	"SE":    "https://ec.nintendo.com/api/SE/en/search/sales?count=10&offset=0#Sweden",
	"DE":    "https://ec.nintendo.com/api/DE/de/search/sales?count=10&offset=0",
	"CH":    "https://ec.nintendo.com/api/CH/de/search/sales?count=10&offset=0#Switzerland",
	"FR":    "https://ec.nintendo.com/api/FR/fr/search/sales?count=10&offset=0",
	"BE-fr": "https://ec.nintendo.com/api/BE/fr/search/sales?count=10&offset=0#Belgium",
	"IT":    "https://ec.nintendo.com/api/IT/it/search/sales?count=10&offset=0",
	"NL":    "https://ec.nintendo.com/api/NL/nl/search/sales?count=10&offset=0",
	"BE-nl": "https://ec.nintendo.com/api/BE/nl/search/sales?count=10&offset=0#Belgium",
	"RU":    "https://ec.nintendo.com/api/RU/ru/search/sales?count=30&offset=0",
	"ES":    "https://ec.nintendo.com/api/ES/es/search/sales?count=30&offset=0",
	"MX":    "https://ec.nintendo.com/api/MX/es/search/sales?count=30&offset=0#Mexico",
	"CO":    "https://ec.nintendo.com/api/CO/es/search/sales?count=10&offset=0#Columbia",
	"AR":    "https://ec.nintendo.com/api/AR/es/search/sales?count=10&offset=0#Argentina",
	"CL":    "https://ec.nintendo.com/api/CL/es/search/sales?count=10&offset=0#Chile",
	"PE":    "https://ec.nintendo.com/api/PE/es/search/sales?count=10&offset=0#Peru",
	"PT":    "https://ec.nintendo.com/api/PT/pt/search/sales?count=30&offset=0",
	"BR":    "https://ec.nintendo.com/api/BR/pt/search/sales?count=10&offset=0",
	"HK":    "https://ec.nintendo.com/api/HK/zh/search/sales?count=10&offset=0",
	"KR":    "https://ec.nintendo.com/api/KR/ko/search/sales?count=10&offset=0",
}

//Contents struct
type Contents struct {
	Contents []Content `json:"contents"`
	Length   int       `json:"length"`
	Offset   int       `json:"offset"`
	Total    int       `json:"total"`
}

//Content struct
type Content struct {
	ContentType        string     `json:"content_type"`
	DominantColors     []string   `json:"dominant_colors"`
	FormalName         string     `json:"formal_name"`
	HeroBannerURL      string     `json:"hero_banner_url"`
	ID                 int        `json:"id"`
	IsNew              bool       `json:"is_new"`
	MembershipRequired bool       `json:"membership_required"`
	PublicStatus       string     `json:"public"`
	RatingInfo         RatingInfo `json:"rating_info"`
	ReleaseDateOnEshop string     `json:"release_date_on_eshop"`
	Screenshots        []Images   `json:"screenshots"`
	Tags               []string   `json:"tags"`
	TargetTitles       []string   `json:"target_titles"`
}

//RatingInfo struct
type RatingInfo struct {
	ContentDescriptors []ContentDescriptors `json:"content_descriptors"`
	Rating             Rating               `json:"rating"`
	RatingSystem       RatingSystem         `json:"rating_system"`
}

//ContentDescriptors struct
type ContentDescriptors struct {
	ID          string `json:"id"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	SvgImageURL string `json:"svg_image_url"`
	Type        string `json:"type"`
}

//Rating struct
type Rating struct {
	Age         string `json:"age"`
	ID          string `json:"id"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Provisional bool   `json:"provisional"`
	SvgImageURL string `json:"svg_image_url"`
}

// RatingSystem struct
type RatingSystem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Images struct
type Images struct {
	Images []URL `json:"images"`
}

//URL struct
type URL struct {
	URL string `json:"url"`
}

//GameDiscount struct
type GameDiscount struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Banner string `json:"hero_banner_url"`
	Prices Prices `json:"prices"`
}

//Prices struct
type Prices struct {
	Personalized bool         `json:"personalized"`
	Country      string       `json:"country"`
	Prices       []PricesList `json:"prices"`
}

//PricesList stuct
type PricesList struct {
	TitleID       int      `json:"title_id"`
	SaleStatus    string   `json:"sales_status"`
	RegularPrice  Regular  `json:"regular_price"`
	DiscountPrice Discount `json:"discount_price"`
}

//Regular struct
type Regular struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	RawValue string `json:"raw_value"`
}

//Discount struct
type Discount struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	RawValue      string `json:"raw_value"`
	StartDateTime string `json:"start_datetime"`
	EndDateTime   string `json:"end_datetime"`
}

//GetDiscountGameFrom func
func GetDiscountGameFrom(api string) []byte {
	data, err := http.Get(api)
	if err != nil {
		return nil
	}
	if data.StatusCode != 200 {
		return nil
	}
	body, err := ioutil.ReadAll(data.Body)
	data.Body.Close()
	if err != nil {
		return nil
	}
	contents := Contents{}
	json.Unmarshal([]byte(body), &contents)
	gameList := []GameDiscount{}

	u, _ := url.Parse(api)
	countryCode := parserURL(u)

	for _, game := range contents.Contents {
		price := getPriceWithCountryAndGameID(countryCode, game.ID)
		FormalName := []rune(game.FormalName)
		gameFetch := GameDiscount{
			ID:     game.ID,
			Name:   string(FormalName),
			Banner: game.HeroBannerURL,
			Prices: price,
		}
		gameList = append(gameList, gameFetch)
	}
	result, _ := json.MarshalIndent(gameList, "", "    ")
	return result

}

//parserURL func
func parserURL(u *url.URL) string {
	country := strings.Split(u.Path, "/")
	return country[2]
}

func getPriceWithCountryAndGameID(code string, gameID int) Prices {
	priceURL := `https://api.ec.nintendo.com/v1/price?country=` + code + `&ids=` + fmt.Sprintf("%v", gameID) + `&lang=us`
	data, err := http.Get(priceURL)
	if err != nil {
		return Prices{}
	}
	if data.StatusCode != 200 {
		return Prices{}
	}
	body, err := ioutil.ReadAll(data.Body)
	data.Body.Close()
	if err != nil {
		return Prices{}
	}
	prices := Prices{}
	json.Unmarshal([]byte(body), &prices)
	return prices
}
