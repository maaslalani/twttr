package twitter

import (
	"os/exec"
)

// Twurl executes the twurl command.
func Twurl(args ...string) ([]byte, error) {
	c := exec.Command("twurl", args...)
	out, err := c.Output()
	return out, err
}
