all: unix windows osx


build:
	go build -o bin/server . 

unix:
	GOOS=linux GOARCH=amd64 go build -o bin/server_linux .

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/server_windows .

osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/server_osx . 