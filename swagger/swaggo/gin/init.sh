export GOPATH=$GOPATH
export GOROOT=$GOROOT
export PATH=$PATH:$GOROOT:$GOPATH
swag init -g main.go -o docs
