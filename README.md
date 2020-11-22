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

You'll need [redis](https://redis.io/) to be installed to use all features.

## Configuration

The file config.json can be edited to set the port that you want for the API and the settings of the REDIS database. If you put -1 in REDIS_PORT it will disable the database features.

## Features

### Traceroute
Main feature is a traceroute made by the API, you can request by using this endpoint : /traceroute via POST method, you need to specify the address in json 
```json
{"address":"imt-atlantique.fr"}
```
After doing the traceroute, the API will find locations of the different IPs using [ip-api](https://ip-api.com/).
Each location will be stored in the redis database under the following key : ip + "/loc"
```
  
```

The API will return a json containing the address of the traceroute, all hops that it contains and theirs locations. This json is stored in the Redis database using the address as key.

### Get all routes stored in database

Using the GET endpoint /all-routes will return a json containing all routes that have been stored in the redis database.
