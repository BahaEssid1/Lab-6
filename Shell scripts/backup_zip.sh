#!/bin/bash

# Folder to backup
FOLDER_TO_BACKUP="$HOME/lab6"  # Folder to backup
BACKUP_FILE="$HOME/lab6/backup.zip"  # Backup inside lab6

# Check if folder exists
if [ -d "$FOLDER_TO_BACKUP" ]; then
    echo "Creating backup of $FOLDER_TO_BACKUP..."
    zip -r "$BACKUP_FILE" "$FOLDER_TO_BACKUP"
    echo "Backup created successfully at $BACKUP_FILE"
else
    echo "Folder $FOLDER_TO_BACKUP does not exist."
fi

