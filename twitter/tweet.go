package twitter

// CreateTweetResponseData represents the data returned from a successful tweet
// creation.
type CreateTweetResponseData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// CreateTweetResponse represents a response after successfully creating a new tweet.
type CreateTweetResponse struct {
	Data CreateTweetResponseData `json:"data"`
}

// CreateTweet creates a new tweet on behalf of the authenticated user.
func CreateTweet(text string) error {
	_, err := Twurl("/2/tweets", "-X", "POST", "-d", "text="+text)
	return err
}
