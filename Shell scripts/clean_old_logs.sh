#!/bin/bash

# Directory containing the logs
LOG_DIR="$HOME/lab6/logs"

# Check if the directory exists
if [ -d "$LOG_DIR" ]; then
    echo "Cleaning logs older than 7 days in $LOG_DIR..."
    find "$LOG_DIR" -type f -name "*.log" -mtime +7 -exec rm -f {} \;
    echo "Cleanup completed successfully."
else
    echo "Directory $LOG_DIR does not exist."
fi
