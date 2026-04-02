#!/bin/sh
set -e

# Start nginx in the background
nginx

# Run the Go app in the foreground — container lifetime tied to this process
exec /go/bin/app
