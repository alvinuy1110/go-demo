#!/bin/bash

export GOINSECURE="proxy.golang.org/*,github.com,github.com/*"
export GONOSUMDB="proxy.golang.org/*,github.com,github.com/*"
export GOPRIVATE="proxy.golang.org/*,github.com,github.com/*"

go get .