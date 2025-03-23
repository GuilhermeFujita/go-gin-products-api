build:
	go build -o dist/api cmd/main.go

run: build
	GIN_MODE=release ./dist/api

clean:
	rm -rf dist