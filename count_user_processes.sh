#!/bin/bash

# Ask for username
read -p "Enter the username: " USERNAME

# Count the number of processes for that user
PROCESS_COUNT=$(ps -u "$USERNAME" --no-headers | wc -l)

echo "User '$USERNAME' is running $PROCESS_COUNT process(es)."

