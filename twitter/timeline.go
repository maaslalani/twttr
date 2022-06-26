package twitter

import (
	"encoding/json"
	"fmt"
	"time"
)

// Timeline is the returned JSON response from the Twitter API.
type Timeline struct {
	Tweets   []Tweet  `json:"data"`
	Includes Includes `json:"includes"`
	Meta     Meta     `json:"meta"`
}

// Tweet is a tweet, right?
type Tweet struct {
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Text      string    `json:"text"`
}

// Includes is the included data from the Twitter API.
type Includes struct {
	Users []User `json:"users"`
}

// Meta is the metadata from the Twitter API.
type Meta struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
}

// HomeTimeline returns the user's home timeline
func HomeTimeline(uid string) (Timeline, error) {
	response := Timeline{}
	bytes, err := Twurl(fmt.Sprintf("/2/users/%s/timelines/reverse_chronological?tweet.fields=created_at&expansions=author_id&exclude=replies,retweets", uid))
	if err != nil {
		return Timeline{}, err
	}
	json.Unmarshal(bytes, &response)
	return response, nil
}
