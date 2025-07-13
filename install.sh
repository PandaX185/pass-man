#!/bin/sh

set -e

BINARY_NAME="passman"
INSTALL_DIR="/usr/local/bin"
ARCHIVE_URL="https://github.com/PandaX185/pass-man/releases/download/v1.1-beta/passman"
TMP_DIR="/tmp/$BINARY_NAME-install"

echo "📦 Starting installation of $BINARY_NAME..."
mkdir -p "$TMP_DIR"

echo "🌐 Downloading binary..."
curl -L "$ARCHIVE_URL" -o "$TMP_DIR/$BINARY_NAME"

if [[ ! -f "$TMP_DIR/$BINARY_NAME" ]]; then
  echo "❌ Download failed. Exiting."
  exit 1
fi

chmod +x "$TMP_DIR/$BINARY_NAME"

if [[ "$EUID" -ne 0 ]]; then
  echo "🔐 Sudo required to install in $INSTALL_DIR"
  sudo mv "$TMP_DIR/$BINARY_NAME" "$INSTALL_DIR/"
else
  mv "$TMP_DIR/$BINARY_NAME" "$INSTALL_DIR/"
fi

rm -rf "$TMP_DIR"
mkdir -p "$HOME/.config/passman"

echo "🔑 Generating encryption key..."
curl -sL https://raw.githubusercontent.com/PandaX185/pass-man/refs/tags/v1.0-beta/gen-key.sh | bash
echo "🔑 Encryption key generated and stored in /usr/local/bin/pass-man.env"
echo "🚀 Installation complete!"
echo "➡️  You can now run it using: $BINARY_NAME"
