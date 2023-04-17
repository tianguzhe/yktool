package util

import (
	"io"
	"net/http"
	"net/url"
)

func GetWithParams(baseUrl string, data url.Values) ([]byte, error) {

	u, _ := url.ParseRequestURI(baseUrl)
	u.RawQuery = data.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.0.0")

	// fmt.Println(u.String())

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	resp_str, _ := io.ReadAll(resp.Body)

	return resp_str, nil
}
