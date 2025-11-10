#!/bin/dash

mkdir "part$1"
cd "part$1"
go mod init "project$1"
