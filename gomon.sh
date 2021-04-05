#!/bin/bash
while inotifywait $(find . -name '*.go') -e MODIFY;
do
    killall main
    go run main.go &
done
