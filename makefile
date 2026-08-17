windows:
	mkdir -p export
	GOOS=windows GOARCH=386 go build -o export/git_setup.exe main.go

run:
	go run .
