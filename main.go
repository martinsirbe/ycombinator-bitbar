package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"

)

var (
	ycombinatorApiUrl = "https://hacker-news.firebaseio.com"
	version           = "v0"
	storyCount        = 5 // Number of stories to display
)

func main() {
	netClient := InitialiseClient()
	apiUrl := fmt.Sprintf("%s/%s", ycombinatorApiUrl, version)

	stories := GetTopStories(netClient, apiUrl)
	storyDetails := GetStoryDetails(netClient, apiUrl, stories)
	log.Info(storyDetails)
}
