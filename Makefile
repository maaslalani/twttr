.SILENT: make mock error
.PHONY: mock

PKG="github.com/maaslalani/twttr"

build:
	go build -o twttr

install:
	go install

# Runs the program normally
make:
	go run .

# Runs the program with mock twurl. See mock/twurl
mock:
	go run -ldflags="-X ${PKG}/twitter.command=mock/twurl" .

# Runs the program with a command that doesn't exist to simulate a user not
# having the twurl command installed.
error:
	go run -ldflags="-X ${PKG}/twitter.command=commanddoesnotexist" .
