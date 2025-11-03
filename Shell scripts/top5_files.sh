#!/bin/bash

echo "Top 5 largest files in $(pwd):"
find . -type f -exec du -h {} + | sort -rh | head -n 5
