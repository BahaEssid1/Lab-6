#!/bin/bash

# Threshold for free space (percentage)
THRESHOLD=10

# Get the root filesystem usage percentage
USAGE=$(df / | awk 'NR==2 {print $5}' | sed 's/%//')

if [ "$USAGE" -ge $((100 - THRESHOLD)) ]; then
    echo "[ALERT] Disk space low! Used: $USAGE%"
else
    echo "[OK] Disk space is sufficient. Used: $USAGE%"
fi
