#!/usr/bin/env bash
set -euo pipefail

# Default install dir, can be overridden with env AI_COMMITIZEN_INSTALL_DIR
INSTALL_DIR="${AI_COMMITIZEN_INSTALL_DIR:-$HOME/.local/bin}"

echo "[ai-commitizen] Install directory: $INSTALL_DIR"
mkdir -p "$INSTALL_DIR"

# Move to the script directory (scripts/commitizen)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "[ai-commitizen] Building binary..."
go build -o "$INSTALL_DIR/ai-commitizen" ./cmd

echo "[ai-commitizen] Binary installed at: $INSTALL_DIR/ai-commitizen"

# Set git alias so it can be used as `git cz`
echo "[ai-commitizen] Configuring git alias 'cz'..."
git config --global alias.cz '!ai-commitizen'

echo
echo "[ai-commitizen] Done."
echo
echo "Usage:"
echo "  1) Make sure $INSTALL_DIR is in your PATH."
echo "     Example (zsh):"
echo "       echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.zshrc"
echo "       source ~/.zshrc"
echo
echo "  2) Inside a git repo, after git add, run:"
echo "       git cz"
echo
echo "     The flow is as follows:"
echo "       - Asked for ticket number"
echo "       - Choose commit type"
echo "       - Read staged diff"
echo "       - Generate commit message with AI"
echo "       - Confirm, then automatically git commit"
