cd user

go mod tidy

go env -w GOOS=linux
go env -w GOARCH=arm

go build -o user.elf

go env -w GOOS=windows
go env -w GOARCH=amd64

docker rmi -f golangdemomicroserviceuser:v0.0.1
docker build -t golangdemomicroserviceuser:v0.0.1 .