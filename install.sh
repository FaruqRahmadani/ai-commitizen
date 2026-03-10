#!/usr/bin/env bash
set -euo pipefail

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

log_info() {
    printf "${BLUE}[INFO]${NC} %s\n" "$1"
}

log_success() {
    printf "${GREEN}[SUCCESS]${NC} %s\n" "$1"
}

log_warn() {
    printf "${YELLOW}[WARN]${NC} %s\n" "$1"
}

log_error() {
    printf "${RED}[ERROR]${NC} %s\n" "$1"
}

# Default install dir, can be overridden with env AI_COMMITIZEN_INSTALL_DIR
INSTALL_DIR="${AI_COMMITIZEN_INSTALL_DIR:-$HOME/.local/bin}"

printf "${CYAN}\n"
cat << "EOF"
    ___    ____     ______                          _ __  _                
   /   |  /  _/    / ____/___  ____ ___  ____ ___  (_) /_(_)___  ___  ____ 
  / /| |  / /_____/ /   / __ \/ __ `__ \/ __ `__ \/ / __/ /_  / / _ \/ __ \
 / ___ |_/ /_____/ /___/ /_/ / / / / / / / / / / / / /_/ / / /_/  __/ / / /
/_/  |_/___/     \____/\____/_/ /_/ /_/_/ /_/ /_/_/\__/_/ /___/\___/_/ /_/ 
                                                                           
EOF
printf "${NC}\n"
echo "---------------------------------------------------------"
echo "           AI-Powered Commitizen Installer               "
echo "---------------------------------------------------------"
echo

# Check requirements
if ! command -v go &> /dev/null; then
    log_error "Go is not installed. Please install Go first."
    exit 1
fi

if ! command -v git &> /dev/null; then
    log_error "Git is not installed. Please install Git first."
    exit 1
fi

log_info "Installing to directory: ${INSTALL_DIR}"
mkdir -p "$INSTALL_DIR"

# Move to the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

log_info "Building binary..."
if go build -o "$INSTALL_DIR/ai-commitizen" ./cmd; then
    log_success "Build successful!"
else
    log_error "Build failed."
    exit 1
fi

log_info "Binary installed at: ${INSTALL_DIR}/ai-commitizen"

# Set git alias so it can be used as `git cz`
log_info "Configuring git alias 'cz'..."
if git config --global alias.cz '!ai-commitizen'; then
    log_success "Git alias 'cz' configured."
else
    log_warn "Failed to configure git alias. You might need to do it manually."
fi

echo
echo "---------------------------------------------------------"
log_success "Installation Complete! 🎉"
echo "---------------------------------------------------------"
echo
printf "${CYAN}Usage Instructions:${NC}\n"
echo
echo "  1) Ensure ${YELLOW}${INSTALL_DIR}${NC} is in your PATH."
echo "     Example (zsh):"
echo "       echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.zshrc"
echo "       source ~/.zshrc"
echo
echo "  2) Inside a git repository, stage your changes:"
printf "       ${GREEN}git add .${NC}\n"
echo
echo "  3) Run the AI Commitizen:"
printf "       ${GREEN}git cz${NC}\n"
echo
echo "     Workflow:"
echo "       📝 Input Ticket Number (Jira integration)"
echo "       🏷️  Select Commit Type (feat, fix, chore, etc.)"
echo "       🤖 AI Generates Commit Message based on diff"
echo "       ✏️  Review & Edit Message (opens your default EDITOR)"
echo "       🚀 Confirm & Commit"
echo
