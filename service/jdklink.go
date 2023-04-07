package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

func GetJdk(version string, ch chan []byte, os string) {

	resp, err := http.Get(fmt.Sprintf("https://api.adoptium.net/v3/assets/latest/%s/hotspot?os=%s&image_type=jdk", version, os))

	if err != nil {
		color.Red("%v", err)
	}

	defer resp.Body.Close()

	resp_str, _ := io.ReadAll(resp.Body)

	ch <- resp_str

}
