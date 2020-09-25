APP_NAME ="csvplay"

hello:
	echo "Hello"

build:
	go build -o bin/$(APP_NAME) main.go

install: build
	cp bin/$(APP_NAME) /usr/local/bin

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/$(APP_NAME)-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/$(APP_NAME)-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/$(APP_NAME)-windows-386 main.go
