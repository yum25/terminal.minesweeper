#!/bin/sh
set -e

APP_NAME="terminal.minesweeper"
REPO="yum25/$APP_NAME"
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
TEMP_APP="${TEMP_DIR}/${APP_NAME}"

cleanup() {
    rm -rf "$TEMP_DIR"
    echo "Cleaning up..."
}
trap cleanup EXIT

URL="https://github.com/${REPO}/releases/download/${VERSION}/${APP_NAME}_${OS}_${ARCH}.tar.gz"
echo "Downloading ${APP_NAME} ${VERSION} for ${OS}/${ARCH} from $URL..."
curl -sL "$URL" -o "${TEMP_APP}.tar.gz"
tar -xz -C "$TEMP_DIR" -f "${TEMP_APP}.tar.gz"


echo "Installing ${BINARY} to ${INSTALL_DIR}..."
mv "${TEMP_DIR}/${BINARY}" "${INSTALL_DIR}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo "Time to start sweeping! Run '${BINARY}' to get started."