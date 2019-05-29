source go-bootstrap.sh
eval "$(GIMME_GO_VERSION=1.11 gimme)"
export GOPATH=$HOME/go
go run main.go
