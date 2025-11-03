#!/bin/bash

# Website to check
URL="https://www.netacad.com/"  # Change to the website you want

# Send a HEAD request and get HTTP status code
STATUS=$(curl -o /dev/null -s -w "%{http_code}" "$URL")

if [ "$STATUS" -eq 200 ]; then
    echo "[OK] $URL is accessible (HTTP 200)"
else
    echo "[ERROR] $URL is not accessible (HTTP $STATUS)"
fi
