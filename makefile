windows:
	mkdir -p export
	GOOS=windows GOARCH=386 go build -o export/windows.exe main.go

run:
	go run .
