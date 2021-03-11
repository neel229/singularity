package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// YTStats holds the YT data retrieved for the channel
type YTStats struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

// Items stores the ID + Statistics for
// a given channel
type Items struct {
	Kind  string `json:"kind"`
	ID    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

// Stats stores the following information
// total views, total subscribers
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

// GetData fetches a YouTube channel's data provided the ID
func GetData(channelID string, apiKey string) (Items, error) {
	var response YTStats
	// We want to craft a new GET request that will include the query parameters we want
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	// here we define the query parameters and their respective values
	q := req.URL.Query()
	// notice how I'm using os.Getenv() to pick up the environment
	// variables that we defined earlier. No hard coded credentials here
	q.Add("key", apiKey)
	q.Add("id", channelID)
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// finally we make the request to the URL that we have just
	// constructed
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Items{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	// we then read in all of the body of the
	// JSON response
	body, _ := ioutil.ReadAll(resp.Body)
	// and finally unmarshal it into an Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, err
	}
	// we only care about the first Item in our
	// Items array, so we just send that back
	return response.Items[0], nil
}

