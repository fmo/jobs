#!/bin/zsh

URL="$1"

if [ -z "$URL" ]; then
  echo "Usage: $0 <url>"
  exit 1
fi

# Use curl to measure timing directly
OUTPUT=$(curl -s -o /dev/null -w "%{http_code} %{time_total}" "$URL")
HTTP_CODE=$(echo "$OUTPUT" | awk '{print $1}')
DURATION=$(echo "$OUTPUT" | awk '{printf "%.0f", $2 * 1000}')  # in ms

echo "[$(date)] Pinged $URL - Status: $HTTP_CODE - Duration: ${DURATION}ms"
