#!/bin/bash

# Ensure a file argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <file>"
    exit 1
fi

FILE=$1

# Check if the file exists
if [ ! -f "$FILE" ]; then
    echo "File not found: $FILE"
    exit 1
fi

# Initialize the file hash
FHASH=$(md5 -q "$FILE")

while true; do
    # Compute the new hash
    NHASH=$(md5 -q "$FILE")

    # Compare the hashes and run the CLI tool if they differ
    if [ "$NHASH" != "$FHASH" ]; then
        ./mdp -file "$FILE"
        FHASH=$NHASH
    fi

    # Sleep for a while before checking again
    sleep 2
done
