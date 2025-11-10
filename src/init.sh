#!/bin/dash

mkdir "part$1"
cd "part$1"
go mod init "project$1"
touch main.go
cat "../format.txt" > main.go

# go run main.go
