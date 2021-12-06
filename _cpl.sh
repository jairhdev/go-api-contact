#!/bin/bash
clear

# Test
echo "** Testing..."
go test -cover .././...
echo "** Finish!" 

# Compile
nameSys="app"

echo
echo "** Compiling with race..."
go build -race -o $nameSys
echo "** Finish!" 

# Check
echo
echo "** Checking..."
staticcheck ./...
echo "** Finish!"

# Run
echo
echo "** Starting..."
echo
./$nameSys
