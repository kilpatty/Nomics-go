package nomics

import (
	"encoding/json"
	"log"
	"net/http"
)

//Notes - Need a function that shows all the markets that are served. I think this might exist, but something that we can iterate through. Double check on this before committing.

var (
	apiURL = "https://api.nomics.com/v1/"
)

type ApiConfig struct {
	Key            string
	KeyQueryString string
}

func New(key string) ApiConfig {

	kqs := "?key=" + key

	ac := ApiConfig{
		Key:            key,
		KeyQueryString: kqs,
	}

	return ac

}

func doRequest(requestURL string, object interface{}) error {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(object)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
