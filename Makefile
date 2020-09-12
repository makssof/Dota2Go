build:
	go build -i -o ./bin/Dota2Go.exe -ldflags -H=windowsgui
	copy scheme bin\scheme
