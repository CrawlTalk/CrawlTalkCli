env GOOS=windows GOARCH=amd64 go build -o ./build/CrawlTalkCli64.exe CrawlTalkCli.go structs.go ui.go network.go
env GOOS=windows GOARCH=386 go build -o ./build/CrawlTalkCli.exe CrawlTalkCli.go structs.go ui.go network.go
env GOOS=linux GOARCH=386 go build -o ./build/CrawlTalkCli CrawlTalkCli.go structs.go ui.go network.go
