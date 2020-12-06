#!/bin/sh

echo "Clean cache..."
go clean -modcache

exec "$@"
