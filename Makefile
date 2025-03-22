build:
	go build -o dist/api cmd/main.go

run: build
	./dist/api

clean:
	rm -rf dist