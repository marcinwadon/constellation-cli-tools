build:
	go build -o bin/cl-tools main.go

run:
	go run main.go

deps:
	go get github.com/devfacet/gocmd

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/cl-tools-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/cl-tools-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/cl-tools-freebsd-386 main.go