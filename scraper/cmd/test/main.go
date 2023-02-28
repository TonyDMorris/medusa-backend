package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "https://www.castlefineart.com/art/alex-ross/batman-dark-knight-detective1", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Language", "en-US,en;q=0.9")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Host", "www.castlefineart.com")
	request.Header.Add("Pragma", "no-cache")
	request.Header.Add("Referer", "https://www.castlefineart.com/")
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-GPC", "1")
	request.Header.Add("TE", "trailers")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("response.html", bytes, 0644)
	if err != nil {
		panic(err)
	}

}
