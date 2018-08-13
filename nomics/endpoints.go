package nomics

import (
	"net/url"
	"strconv"
	"time"
)

func (ac *ApiConfig) Markets() ([]Market, error) {

	//Quick Note - API does not support ending backslashes.
	requestURL := apiURL + "markets" + ac.KeyQueryString

	var markets []Market

	err := doRequest(requestURL, &markets)

	return markets, err

}

func (ac *ApiConfig) MarketIntervalsByHour(slug string, hours int) ([]Interval, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	hoursQuery := "hours=" + strconv.Itoa(hours)

	slugQuery := "currency=" + slug

	requestURL := apiURL + "markets/interval" + ac.KeyQueryString

	requestURL = requestURL + "&" + hoursQuery + "&" + slugQuery

	var intervals []Interval

	err := doRequest(requestURL, &intervals)

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

	var intervals []Interval

	err := doRequest(requestURL, &intervals)

	return intervals, err

}

func (ac *ApiConfig) MarketPrices(slug string) ([]MarketPrice, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	slugQuery := "currency=" + slug

	requestURL := apiURL + "markets/prices" + ac.KeyQueryString

	requestURL = requestURL + "&" + slugQuery

	var prices []MarketPrice

	err := doRequest(requestURL, &prices)

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

	var intervals []ExchangeInterval

	err := doRequest(requestURL, &intervals)

	return intervals, err

}

func (ac *ApiConfig) ExchangeMarketPrices(slug string) ([]ExchangePrice, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	slugQuery := "currency=" + slug

	requestURL := apiURL + "exchange-markets/prices" + ac.KeyQueryString

	requestURL = requestURL + "&" + slugQuery

	var prices []ExchangePrice

	err := doRequest(requestURL, &prices)

	return prices, err

}

func (ac *ApiConfig) Prices() ([]Price, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "prices" + ac.KeyQueryString

	var prices []Price

	err := doRequest(requestURL, &prices)

	return prices, err

}

func (ac *ApiConfig) Dashboard() ([]Dashboard, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "dashboard" + ac.KeyQueryString

	var dashboards []Dashboard

	err := doRequest(requestURL, &dashboards)

	return dashboards, err

}

func (ac *ApiConfig) Highs() ([]High, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "currencies/highs" + ac.KeyQueryString

	var highs []High

	err := doRequest(requestURL, &highs)

	return highs, err

}

func (ac *ApiConfig) ExchangeRates() ([]ExchangeRate, error) {

	//note to self here - the slug NEEDS TO BE CAPITALIZED

	requestURL := apiURL + "exchange-rates" + ac.KeyQueryString

	var rates []ExchangeRate

	err := doRequest(requestURL, &rates)

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

	var intervals []SupplyInterval

	err := doRequest(requestURL, &intervals)

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

	var intervals []CurrencyInterval

	err := doRequest(requestURL, &intervals)

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

	var intervals []ExchangeRateHistory

	err := doRequest(requestURL, &intervals)

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

	var intervals []ExchangeRateInterval

	err := doRequest(requestURL, &intervals)

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

	var intervals []MarketCapHistory

	err := doRequest(requestURL, &intervals)

	return intervals, err

}
