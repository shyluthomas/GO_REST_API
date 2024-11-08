# GO_REST_API
REST Web api


to run test
cmd = go test ./...

to run go files
go run filename

to clean dependancy link updates
go mod tidy

download all dependency
go mod download

Build the go app
go build
then goto terminal and run
./buildfilename


setup the orm and postgress driver

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv


setup swagger openapi
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
to setup swag command in terminal
go install github.com/swaggo/swag/cmd/swag@latest

add all annotationn and run command "swag init"

in mac add below steps add swag command

export PATH=$PATH:$(go env GOPATH)/bin

source ~/.zshrc

check swagger
swag --version



