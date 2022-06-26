package twitter

import (
	"encoding/json"
)

// User represents a Twitter user.
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// UserResponse represents the returned JSON response from the Twitter API.
type UserResponse struct {
	Data User `json:"data"`
}

// Me returns the current user's profile
func Me() (User, error) {
	response := UserResponse{}
	bytes, err := Twurl("/2/users/me")
	if err != nil {
		return User{}, err
	}
	json.Unmarshal(bytes, &response)
	return response.Data, nil
}
