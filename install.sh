#!/bin/bash
set -e

REPO="cumulus13/config-read"
VERSION="${VERSION:-latest}"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_color() {
    printf "${1}${2}${NC}\n"
}

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case $ARCH in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        armv7l|armv6l)
            ARCH="arm"
            ;;
        *)
            print_color "$RED" "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
    
    case $OS in
        linux|darwin)
            ;;
        msys_nt|mingw*|cygwin*)
            OS="windows"
            ;;
        *)
            print_color "$RED" "Unsupported OS: $OS"
            exit 1
            ;;
    esac
    
    print_color "$GREEN" "Detected: $OS/$ARCH"
}

# Download and install binary
install_binary() {
    if [ "$VERSION" = "latest" ]; then
        DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/config-read_${OS}_${ARCH}.tar.gz"
    else
        DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/config-read_${OS}_${ARCH}.tar.gz"
    fi
    
    # Add .exe extension for Windows
    if [ "$OS" = "windows" ]; then
        DOWNLOAD_URL="${DOWNLOAD_URL}.zip"
    fi
    
    print_color "$YELLOW" "Downloading config-read $VERSION..."
    
    TMP_DIR=$(mktemp -d)
    cd "$TMP_DIR"
    
    # Download with error handling
    if ! curl -sSfL "$DOWNLOAD_URL" -o "config-read.tar.gz"; then
        print_color "$RED" "Failed to download from $DOWNLOAD_URL"
        exit 1
    fi
    
    # Extract
    if [ "$OS" = "windows" ]; then
        unzip -q "config-read.tar.gz"
    else
        tar xzf "config-read.tar.gz"
    fi
    
    # Install
    chmod +x "config-read"
    if [ "$OS" = "windows" ]; then
        mv "config-read.exe" "$INSTALL_DIR/"
    else
        sudo mv "config-read" "$INSTALL_DIR/"
    fi
    
    cd - > /dev/null
    rm -rf "$TMP_DIR"
    
    print_color "$GREEN" "✓ config-read installed successfully to $INSTALL_DIR"
}

# Verify installation
verify_installation() {
    if command -v config-read > /dev/null 2>&1; then
        print_color "$GREEN" "✓ config-read is ready to use!"
        config-read version
    else
        print_color "$RED" "✗ Installation failed"
        exit 1
    fi
}

# Main installation process
main() {
    print_color "$YELLOW" "=== config-read Installer ==="
    
    detect_platform
    
    # Check if already installed
    if command -v config-read > /dev/null 2>&1; then
        CURRENT_VERSION=$(config-read version 2>/dev/null | awk '{print $3}')
        print_color "$YELLOW" "config-read $CURRENT_VERSION is already installed"
        read -p "Do you want to overwrite? [y/N] " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 0
        fi
    fi
    
    # Check prerequisites
    if ! command -v curl > /dev/null 2>&1; then
        print_color "$RED" "curl is required but not installed"
        exit 1
    fi
    
    install_binary
    verify_installation
    
    print_color "$GREEN" "=== Installation Complete ==="
}

main "$@"
