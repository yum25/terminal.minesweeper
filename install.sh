#!/bin/sh
set -e

REPO="yum25/terminal.minesweeper"
BINARY="tsweep"
INSTALL_DIR="/usr/local/bin"

VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' \
    | cut -d '"' -f 4)

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case $ARCH in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

TEMP_DIR=$(mktemp -d)
TEMP_BINARY="${TEMP_DIR}/${BINARY}"

URL="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY}_${OS}_${ARCH}.tar.gz"
echo "Downloading ${BINARY} ${VERSION} for ${OS}/${ARCH}..."
curl -sL "$URL" -o "$TEMP_BINARY"
tar xz -C "$TEMP_DIR" -f "${TEMP_BINARY}.tar.gz"


echo "Installing to ${INSTALL_DIR}/${BINARY}..."
mv "$TEMP_BINARY" "${INSTALL_DIR}/${BINARY}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo "Cleaning up..."
rm -rf "$TEMP_DIR"

echo "Time to start sweeping! Run '${BINARY}' to get started."