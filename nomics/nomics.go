package nomics

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
