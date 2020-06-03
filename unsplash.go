package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// RandomImageResp is the JSON response of the Random Image
type RandomImageResp struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Links       Links
	CreatedAt   time.Time `json:"created_at"`
}

// Links is the JSON links structure
type Links struct {
	Download string `json:"download"`
}

func unsplash(topic string) *http.Response {

	unsplashAccessKey := "YourUnsplashAccessKey"

	baseURI := "https://api.unsplash.com"

	url := baseURI + "/photos/random"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept-Version", "v1")
	req.Header.Set("Authorization", "Client-ID "+unsplashAccessKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data RandomImageResp
	e := json.Unmarshal(body, &data)
	if e != nil {
		log.Fatal(e)
	}

	req2, _ := http.NewRequest("GET", data.Links.Download, nil)
	req2.Header.Set("Accept-Version", "v1")
	req2.Header.Set("Authorization", "Client-ID "+unsplashAccessKey)

	imageRes, err := client.Do(req2)

	return imageRes
}
