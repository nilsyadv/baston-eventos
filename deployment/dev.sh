apk --no-cache add  tzdata git
apk add openssh

go get -u github.com/swaggo/swag/cmd/swag
go get github.com/cespare/reflex

reflex -r '\.go' -s -- sh -c 'go run /app/cmd/main.go -e development'