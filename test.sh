#!/bin/bash

for file in y4-examples/*.y4; do
    echo "running script for $file"
    go run cmd/main.go -quiet $file
    if [ $? -ne 0 ]; then
        echo "error occurred with $file"
        exit 1
    fi
done

echo "all scripts executed successfully"
