.SILENT: make mock
.PHONY: mock

make:
	go run .

mock:
	go run -ldflags="-X github.com/maaslalani/twttr/twitter.command=mock/twurl" .
