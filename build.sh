#!/bin/bash

echo "Building binary"

GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating zip file"

zip deployment.zip main

echo "cleanup!"

rm main
