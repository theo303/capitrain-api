# capitrain-api
Tested on Debian 10. It will not work on windows because it uses the "traceroute" command which is not available on windows.

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/zgegonline/capitrain-api
cd $GOPATH/src/github.com/zgegonline/capitrain-api
go build
./capitrain-api
```

