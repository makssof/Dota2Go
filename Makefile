build:
	go build -i -o ./bin/Dota2Go.exe -ldflags -H=windowsgui ./cmd/Dota2Go/main.go
	copy assets\scheme bin\scheme
run:
	./bin/Dota2Go.exe
