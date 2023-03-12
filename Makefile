all: build

pack:
	esc -o pkg/Dota2Catcher/static.go -pkg Dota2Catcher assets
build:
	esc -o pkg/Dota2Catcher/static.go -pkg Dota2Catcher assets
	export GOOS=windows && \
	export GOARCH=amd64 && \
	go build -o ./bin/Dota2Go.exe -ldflags="-s -w -H=windowsgui" ./cmd/Dota2Go/main.go
run:
	./bin/Dota2Go.exe
