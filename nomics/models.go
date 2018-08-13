package nomics

import "time"

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
