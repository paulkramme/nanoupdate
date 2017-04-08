all:
	go build nanoupdate.go
clean:
	rm nanoupdate
windows:
	env GOOS=windows GOARCH=amd64 go build nanoupdate.go
