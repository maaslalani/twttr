package twitter

import (
	"os/exec"

	"github.com/pkg/errors"
)

const twurlNotInstalled = `Looks like twurl is not installed on your machine.
To use twttr, you'll need to install twurl (https://github.com/twitter/twurl)
and follow the instructions to get setup with a developer account`

// Twurl executes the twurl command.
func Twurl(args ...string) ([]byte, error) {
	c := exec.Command("twurl", args...)
	out, err := c.Output()
	if err != nil {
		return out, errors.Wrap(err, twurlNotInstalled)
	}
	return out, nil
}
