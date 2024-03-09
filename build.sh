go env -w GOOS=linux
go env -w GOARCH=amd64

go build -ldflags="-s -w -v" -o main .

go env -w GOOS=windows
go env -w GOARCH=amd64