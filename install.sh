#!/bin/bash

set -e

# Meba CLI Installation Script

REPO="meba-cli/meba"
BINARY_NAME="meba"
INSTALL_DIR="/usr/local/bin"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case $OS in
    linux) PLATFORM="linux-$ARCH" ;;
    darwin) PLATFORM="darwin-$ARCH" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Get latest release
echo "🔍 Getting latest release..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Failed to get latest release"
    exit 1
fi

echo "📦 Latest version: $LATEST_RELEASE"

# Download binary
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/meba-$PLATFORM"
TEMP_FILE="/tmp/$BINARY_NAME"

echo "⬇️  Downloading $BINARY_NAME..."
curl -L -o "$TEMP_FILE" "$DOWNLOAD_URL"

if [ ! -f "$TEMP_FILE" ]; then
    echo "❌ Download failed"
    exit 1
fi

# Make executable
chmod +x "$TEMP_FILE"

# Install binary
echo "📦 Installing to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
    mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
else
    sudo mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
fi

# Verify installation
if command -v $BINARY_NAME >/dev/null 2>&1; then
    echo "✅ Meba CLI installed successfully!"
    echo "🚀 Run 'meba --help' to get started"
    $BINARY_NAME --version
else
    echo "❌ Installation failed"
    exit 1
fi