#!/bin/bash

# Configuration
TO_EMAIL="essidbaha18@gmail.com"   # Recipient
SUBJECT="ALERT: Test Email"
MESSAGE="This is a test alert message from your script."

# Send the email using msmtp
if command -v msmtp &> /dev/null; then
    echo -e "Subject: $SUBJECT\n\n$MESSAGE" | msmtp "$TO_EMAIL"
    echo "[OK] Email sent to $TO_EMAIL"
else
    echo "[ERROR] 'msmtp' command not found. Please install msmtp and configure it."
fi

