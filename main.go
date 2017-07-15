package main

import (
	"fmt"
	"time"
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

	fmt.Println("🔥\n---")
	for _, story := range storyDetails {
		//fmt.Printf("#{%s} | href=#{%s} color=white", story.Title, story.Url)
		fmt.Printf("%s (%d) | href=%s color=#a7cc53 size=13\n", story.Title, len(story.Kids), story.Url)
		fmt.Printf("%s | color=#b7ed38 size=9\n", time.Unix(story.Time, 0).Format("2006-01-02 15:04:05"))
		fmt.Println("---")
	}
}
