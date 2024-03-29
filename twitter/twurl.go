package twitter

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Used for swapping out `twurl` executable for a mock
// go run -ldflags="-X github.com/maaslalani/twttr/twitter.command=mock/twurl" .
var command = "twurl"

const twurlNotInstalled = `Looks like twurl is not installed on your machine or you have not run twurl authorize.

To use twttr, you'll need to:

  1. Install twurl (https://github.com/twitter/twurl)
  2. Get a twitter developer account (https://t.co/developer)
  3. twurl authorize --consumer-key <key> --consumer-secret <secret>
`

// Twurl executes the twurl command.
func Twurl(args ...string) ([]byte, error) {
	c := exec.Command(command, args...)
	out, err := c.Output()
	if err != nil {
		return out, errors.Errorf("%s\n%s", err, twurlNotInstalled)
	}
	return out, nil
}
