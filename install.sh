#!/bin/bash

# Define variables
BIN_DIR="/usr/local/bin"
BIN_NAME="c-lient"
SRC_BIN="./${BIN_NAME}"

# Check if the binary exists
if [ ! -f "$SRC_BIN" ]; then
  echo "Error: Binary not found at $SRC_BIN"
  exit 1
fi

# Copy the binary to /usr/local/bin
echo "Installing $BIN_NAME to $BIN_DIR"
sudo cp "$SRC_BIN" "$BIN_DIR/"
sudo chmod +x "$BIN_DIR/$BIN_NAME"

# Verify installation
if command -v "$BIN_NAME" > /dev/null 2>&1; then
  echo "$BIN_NAME installed successfully!"
else
  echo "Installation failed."
  exit 1
fi