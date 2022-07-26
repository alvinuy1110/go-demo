GIN
===

https://go.dev/doc/tutorial/web-service-gin

## Bypass SSL

```


export GOMOD=""
# comma list of domains
export GOINSECURE="proxy.golang.org/*,github.com,github.com/*,golang.org/*,gopkg.in/"
export GONOSUMDB="proxy.golang.org/*,github.com,github.com/*,golang.org/*,gopkg.in/"
export GOPRIVATE="proxy.golang.org/*,github.com,github.com/*,golang.org/*,gopkg.in/*,google.golang.org/*"

## bypss all instead!!!! (works)
export GOINSECURE="*.org/*,*"
export GONOSUMDB="*.org/*,*"
export GOPRIVATE="*.org/*,*"

```

## Fetching Import
```go
go get .
```



## Get the certs
Note: May not be needed
echo -n | openssl s_client -connect proxy.golang.org:443  | sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > proxy.pem


sudo cp proxy.pem /usr/local/share/ca-certificates/proxy.pem
sudo update-ca-certificates

## Start

```go
go run .
or
go run main.go
```

## Curl

```
curl http://localhost:8080/albums/2
```

## Testing

https://circleci.com/blog/gin-gonic-testing/

```
go test
```

### Convention

create a file named main_test.go 

Note: Each test file within your project must end with _test.go and each test method must start with a Test prefix. This is a standard naming convention for a valid test.

