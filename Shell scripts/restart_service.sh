#!/bin/bash

# Name of the service to restart
SERVICE_NAME="cloudflare-warp"  # Make sure this is the exact service name

# Check if the service exists
if systemctl list-units --type=service | grep -q "$SERVICE_NAME"; then
    echo "Restarting $SERVICE_NAME..."
    sudo systemctl restart "$SERVICE_NAME"

    # Check status
    if systemctl is-active --quiet "$SERVICE_NAME"; then
        echo "[OK] $SERVICE_NAME restarted successfully."
    else
        echo "[ERROR] Failed to restart $SERVICE_NAME."
    fi
else
    echo "[ERROR] Service $SERVICE_NAME not found."
fi
