package nomics

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Market struct {
	Exchange string `json:"exchange"`
	Market   string `json:"market"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
}

type MarketPrice struct {
	Exchange  string    `json:"exchange"`
	Quote     string    `json:"quote"`
	Price     string    `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type Interval struct {
	Exchange       string    `json:"exchange"`
	Quote          string    `json:"quote"`
	Volume         string    `json:"volume"`
	Open           string    `json:"open"`
	OpenTimestamp  time.Time `json:"open_timestamp"`
	Close          string    `json:"close"`
	CloseTimestamp time.Time `json:"close_timestamp"`
}

type ExchangeInterval struct {
	Exchange       string    `json:"exchange"`
	Base           string    `json:"base"`
	Quote          string    `json:"quote"`
	VolumeBase     string    `json:"volume_base"`
	OpenQuote      string    `json:"open_quote"`
	OpenTimestamp  time.Time `json:"open_timestamp"`
	CloseQuote     string    `json:"close_quote"`
	CloseTimestamp time.Time `json:"close_timestamp"`
}

type ExchangePrice struct {
	Exchange   string    `json:"exchange"`
	Base       string    `json:"base"`
	Quote      string    `json:"quote"`
	PriceQuote string    `json:"price_quote"`
	Timestamp  time.Time `json:"timestamp"`
}

type ExchangeRate struct {
	Currency  string    `json:"currency"`
	Rate      string    `json:"rate"`
	Timestamp time.Time `json:"timestamp"`
}

type ExchangeRateHistory struct {
	Timestamp time.Time `json:"timestamp"`
	Rate      string    `json:"rate"`
}

type ExchangeRateInterval struct {
	Currency       string    `json:"currency"`
	Open           string    `json:"open"`
	OpenTimestamp  time.Time `json:"open_timestamp"`
	Close          string    `json:"close"`
	CloseTimestamp time.Time `json:"close_timestamp"`
}

type Price struct {
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

type Dashboard struct {
	Currency          string    `json:"currency"`
	DayOpen           string    `json:"dayOpen"`
	DayVolume         string    `json:"dayVolume"`
	DayOpenVolume     string    `json:"dayOpenVolume"`
	WeekOpen          string    `json:"weekOpen"`
	WeekVolume        string    `json:"weekVolume"`
	WeekOpenVolume    string    `json:"weekOpenVolume"`
	MonthOpen         string    `json:"monthOpen"`
	MonthVolume       string    `json:"monthVolume"`
	MonthOpenVolume   string    `json:"monthOpenVolume"`
	YearOpen          string    `json:"yearOpen"`
	YearVolume        string    `json:"yearVolume"`
	YearOpenVolume    string    `json:"yearOpenVolume"`
	Close             string    `json:"close"`
	High              string    `json:"high"`
	HighTimestamp     time.Time `json:"highTimestamp"`
	HighExchange      string    `json:"highExchange"`
	HighQuoteCurrency string    `json:"highQuoteCurrency"`
	AvailableSupply   string    `json:"availableSupply"`
	MaxSupply         string    `json:"maxSupply"`
}

type High struct {
	Currency  string    `json:"currency"`
	Price     string    `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Exchange  string    `json:"exchange"`
	Quote     string    `json:"quote"`
}

type SupplyInterval struct {
	Currency       string    `json:"currency"`
	OpenAvailable  string    `json:"open_available"`
	OpenMax        string    `json:"open_max"`
	OpenTimestamp  time.Time `json:"open_timestamp"`
	CloseAvailable string    `json:"close_available"`
	CloseMax       string    `json:"close_max"`
	CloseTimestamp time.Time `json:"close_timestamp"`
}

type CurrencyInterval struct {
	Currency       string    `json:"currency"`
	Volume         string    `json:"volume"`
	Open           string    `json:"open"`
	OpenTimestamp  time.Time `json:"open_timestamp"`
	Close          string    `json:"close"`
	CloseTimestamp time.Time `json:"close_timestamp"`
}

type MarketCapHistory struct {
	Timestamp time.Time `json:"timestamp"`
	MarketCap string    `json:"market_cap"`
}

func (ac *ApiConfig) Markets() ([]Market, error) {

	//Quick Note - API does not support ending backslashes.
	requestURL := apiURL + "markets" + ac.KeyQueryString

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var markets []Market

	err = json.NewDecoder(resp.Body).Decode(&markets)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return markets, err

}

func (ac *ApiConfig) MarketIntervalsByHour(slug string, hours int) ([]Interval, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	hoursQuery := "hours=" + strconv.Itoa(hours)

	slugQuery := "currency=" + slug

	requestURL := apiURL + "markets/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + hoursQuery + "&" + slugQuery

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []Interval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) MarketIntervalsByTime(slug string, start, end time.Time) ([]Interval, error) {

	slugQuery := "currency=" + slug

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "markets/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	requestURL = requestURL + "&" + slugQuery

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []Interval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) MarketPrices(slug string) ([]MarketPrice, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	slugQuery := "currency=" + slug

	requestURL := apiURL + "markets/prices" + ac.KeyQueryString

	requestURL = requestURL + "&" + slugQuery

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var prices []MarketPrice

	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return prices, err

}

func (ac *ApiConfig) ExchangeMarketInterval(slug string, start, end time.Time) ([]ExchangeInterval, error) {

	slugQuery := "currency=" + slug

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "exchange-markets/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	requestURL = requestURL + "&" + slugQuery

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []ExchangeInterval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) ExchangeMarketPrices(slug string) ([]ExchangePrice, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	slugQuery := "currency=" + slug

	requestURL := apiURL + "exchange-markets/prices" + ac.KeyQueryString

	requestURL = requestURL + "&" + slugQuery

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var prices []ExchangePrice

	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return prices, err

}

func (ac *ApiConfig) Prices() ([]Price, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "prices" + ac.KeyQueryString

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var prices []Price

	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return prices, err

}

func (ac *ApiConfig) Dashboard() ([]Dashboard, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "dashboard" + ac.KeyQueryString

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var dashboards []Dashboard

	err = json.NewDecoder(resp.Body).Decode(&dashboards)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dashboards, err

}

func (ac *ApiConfig) Highs() ([]High, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "currencies/highs" + ac.KeyQueryString

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var highs []High

	err = json.NewDecoder(resp.Body).Decode(&highs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return highs, err

}

func (ac *ApiConfig) ExchangeRates() ([]ExchangeRate, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "exchange-rates" + ac.KeyQueryString

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var rates []ExchangeRate

	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rates, err

}

func (ac *ApiConfig) SuppliesInterval(start, end time.Time) ([]SupplyInterval, error) {

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "supplies/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []SupplyInterval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) CurrenciesInterval(start, end time.Time) ([]CurrencyInterval, error) {

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "currencies/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []CurrencyInterval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) ExchangeRateHistory(slug string, start, end time.Time) ([]ExchangeRateHistory, error) {

	slugQuery := "currency=" + slug

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "exchange-rates/history" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	requestURL = requestURL + "&" + slugQuery

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []ExchangeRateHistory

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) ExchangeRateIntervals(start, end time.Time) ([]ExchangeRateInterval, error) {

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "exchange-rates/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []ExchangeRateInterval

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}

func (ac *ApiConfig) MarketCapHistory(start, end time.Time) ([]MarketCapHistory, error) {

	//Start time and End time need to be escape for the URI
	startQuery, err := url.ParseQuery("start=" + start.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	endQuery, err := url.ParseQuery("end=" + end.Format(time.RFC3339))

	if err != nil {
		return nil, err
	}

	requestURL := apiURL + "market-cap/history" + ac.KeyQueryString

	requestURL = requestURL + "&" + startQuery.Encode() + "&" + endQuery.Encode()

	spew.Dump(requestURL)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	var intervals []MarketCapHistory

	err = json.NewDecoder(resp.Body).Decode(&intervals)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return intervals, err

}
