# echo "Compiling for every OS and Platform"
# GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
# GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
# GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

clean: 
	rm -r ./bin

build:
	echo "Compiling for every OS and Platform"
	mkdir ./bin
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin ./...
	
run:
	go run cmd/main.go