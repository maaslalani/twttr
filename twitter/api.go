package twitter

import (
	"os/exec"
)

// Twurl executes the twurl command.
func Twurl(request string) ([]byte, error) {
	c := exec.Command("twurl", request)
	out, err := c.Output()
	return out, err
}
