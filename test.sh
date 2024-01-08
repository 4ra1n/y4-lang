#!/bin/bash

for file in examples/*.y4; do
    echo "running script for $file"
    go run cmd/main.go -f $file
    if [ $? -ne 0 ]; then
        echo "error occurred with $file"
        exit 1
    fi
done

echo "all scripts executed successfully"