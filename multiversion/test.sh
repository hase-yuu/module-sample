#!/bin/bash

echo "go get rsc.io/sampler@v1.0.0"
go get rsc.io/sampler@v1.0.0
go build main.go
./main

echo "go get rsc.io/sampler@latest"
go get -u rsc.io/sampler
go build main.go
./main

