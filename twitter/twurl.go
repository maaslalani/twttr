package twitter

import (
	"os/exec"

	"github.com/pkg/errors"
)

const twurlNotInstalled = `Looks like twurl is not installed on your machine.
To use twttr, you'll need to:
  1. Install twurl (https://github.com/twitter/twurl)
  2. Get a twitter developer account (https://developer.twitter.com/apply-for-access)`

// Twurl executes the twurl command.
func Twurl(args ...string) ([]byte, error) {
	c := exec.Command("twurl", args...)
	out, err := c.Output()
	if err != nil {
		return out, errors.Wrap(err, twurlNotInstalled)
	}
	return out, nil
}
