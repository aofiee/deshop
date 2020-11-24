package eshop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//GameDetail struct
type GameDetail struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Banner string `json:"hero_banner_url"`
	Prices Prices `json:"prices,omitempty"`
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

func buildResult(contents Contents, country string) []byte {
	gameList := []GameDetail{}
	for _, game := range contents.Contents {
		price := getPriceWithCountryAndGameID(country, game.ID)
		FormalName := []rune(game.FormalName)
		gameFetch := GameDetail{
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

//GetEndpoint func
func GetEndpoint(country string, service string) []byte {
	var lang = `en`
	if country == `JP` {
		lang = `ja`
	}
	var path string
	switch service {
	case `sale`:
		path = `sales?count=30&offset=0`
		break
	case `new`:
		path = `new?count=30&offset=0`
		break
	case `ranking`:
		path = `ranking?count=10&offset=0`
		break
	}
	endpoint := `https://ec.nintendo.com/api/` + country + `/` + lang + `/search/` + path
	return getEndpointData(country, endpoint)
}

//getEndpointData func
func getEndpointData(country string, endpoint string) []byte {
	data, err := http.Get(endpoint)
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
	return buildResult(contents, country)
}
