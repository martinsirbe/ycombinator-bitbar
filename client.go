package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/MartinsIrbe/ycombinator-bitbar/models"
)

func InitialiseClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func GetTopStories(netClient *http.Client, apiUrl string) []int {
	url := fmt.Sprintf("%s/topstories.json?print=pretty", apiUrl)
	body := makeGetRequest(netClient, url)

	var stories []int
	err := json.Unmarshal(body, &stories)
	if err != nil {
		log.Errorf("failed to convert response body, err: %v", err)
	}

	return stories
}

func GetStoryDetails(netClient *http.Client, apiUrl string, stories []int) []models.Story {
	var storyDetails []models.Story
	for i, id := range stories {
		if i < storyCount {
			url := fmt.Sprintf("%s/item/%d.json?print=pretty", apiUrl, id)
			body := makeGetRequest(netClient, url)

			var story models.Story
			err := json.Unmarshal(body, &story)
			if err != nil {
				log.Fatalf("failed to convert response body, err: %v", err)
			}

			storyDetails = append(storyDetails, story)
		} else {
			break
		}
	}

	return storyDetails
}

func makeGetRequest(netClient *http.Client, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("failed to create new request, err: %v", err)
	}

	res, getErr := netClient.Do(req)
	if getErr != nil {
		log.Fatalf("failed to execute HTTP GET request, err: %v", err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalf("failed to read response body, err: %v", err)
	}
	return body
}
