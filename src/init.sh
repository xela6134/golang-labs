#!/bin/dash

name="part$1"

if [ -d "$name" ]
then
    echo "Folder exists"
    exit 1
fi

mkdir "$name"
cd "$name"
go mod init "project$1"
touch main.go
cat "../format.txt" > main.go

# go run main.go
