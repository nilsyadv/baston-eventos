apk --no-cache add  tzdata git
apk add openssh

go env -w GOPRIVATE=github.com/lin-sel/*
echo "StrictHostKeyChecking no " > /root/.ssh/config
printf "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig

go get -u github.com/swaggo/swag/cmd/swag
go get github.com/cespare/reflex

reflex -r '\.go' -s -- sh -c 'go run /app/cmd/main.go -e development'