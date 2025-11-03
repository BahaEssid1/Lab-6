#!/bin/bash

# Get currently logged-in users
USERS=$(who | awk '{print $1}' | sort | uniq)

if [ -n "$USERS" ]; then
    echo "[OK] Currently logged-in users:"
    echo "$USERS"
else
    echo "[INFO] No users are currently logged in."
fi
