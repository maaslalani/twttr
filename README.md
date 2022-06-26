# twttr

Use Twitter from the comfort of your terminal.

```
twttr
```
### Timeline

When you first launch `twttr`, you will see your home timeline. Press `→/n/j/l` and `←/p/k/h` to traverse through your timeline tweet by tweet.

![Home Timeline](https://user-images.githubusercontent.com/42545625/175827521-898f777a-2f3d-4ad7-90a2-5f51c0cc4320.png)

### Keybindings

Press `?` to access help which shows you the keybindings to navigate around.

![Help Page](https://user-images.githubusercontent.com/42545625/175827520-8e6adf19-5e5d-41d3-9f37-8a3308ba434e.png)

### Compose Tweet

Press `c` to compose a new tweet and type what's on your mind, hit `enter` to share with the world right from your terminal.

![Compose Tweet](https://user-images.githubusercontent.com/42545625/175827518-6e82b5e2-faa7-4392-9d9e-023adfab4824.png)
## Installation

Install `twttr` with Go.

```bash
go install github.com/maaslalani/twttr@latest
```

**Important**: You will also need [Twitter's twurl](https://github.com/twitter/twurl) for `twttr` to function properly.

Install `twurl` with any package manager.

```bash
# Brew
brew install twurl

# Gem
gem install twurl

# Arch Linux (btw)
yay -S ruby-twurl

# Nix
nix-env -iA nixpkgs.twurl
```
## Setup

To run `twttr`, you will need a Twitter development account which can be created from the [Twitter Developer's Page](https://developer.twitter.com/en/apply-for-access).
```
https://developer.twitter.com/en/apply-for-access
```

After you have that access you can create a Twitter app and generate a consumer key and secret.

When you have your consumer key and its secret you authorize your Twitter account to make API requests with that consumer key and secret.

```
twurl authorize --consumer-key key \
                --consumer-secret secret
```

Essentially, if you can run the following command with `twurl` successfully, `twttr` should work without issue.
```
twurl /2/users/me
```

See [Twurl's Getting Started](https://github.com/twitter/twurl#getting-started) section for more details.


