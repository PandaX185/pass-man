#!/bin/sh

ENV_PATH="/usr/local/bin/pass-man.env"

KEY=$(head -c 16 /dev/urandom | base64)

echo "ENCRYPTION_KEY=$KEY" > "$ENV_PATH"
echo "Key written to $ENV_PATH"