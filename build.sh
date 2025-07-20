#!/bin/bash

# ========================================
# 🔐 VEKO DOME - Stealth Mode Build Script
# ========================================

echo ""
echo "🔨 Building VEKO DOME - STEALTH MODE..."
echo ""

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

# Initialize Go module if not exists
if [ ! -f "go.mod" ]; then
    echo "📦 Initializing Go module..."
    go mod init veko-dome
fi

# Create builds directory
mkdir -p builds

# Get required dependencies
echo "⬇️  Downloading dependencies..."
go get github.com/refraction-networking/utls@latest

# Build for multiple platforms
echo ""
echo "🏗️  Cross-compiling for multiple platforms..."

# Linux (64-bit)
echo "🟢 Linux x64..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o builds/veko-dome-linux main.go

# Windows (64-bit)
echo "🟦 Windows x64..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o builds/veko-dome-windows.exe main.go

# macOS (Intel)
echo "🍎 macOS x64..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o builds/veko-dome-macos-x64 main.go

# macOS (Apple Silicon)
echo " macOS ARM64..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o builds/veko-dome-macos-arm64 main.go

# Local build
echo " Building for current platform..."
go build -ldflags="-s -w" -o veko-dome main.go

# Create sample files if not exist
echo ""
echo " Creating sample config and data files..."

# Sample config.json
if [ ! -f "config.json" ]; then
cat > config.json << 'EOF'
{
  "proxy_file": "proxylist.txt",
  "rotation_interval_seconds": 30,
  "tor_enabled": true,
  "tor_proxy": "socks5://127.0.0.1:9050",
  "doh_server": "https://1.1.1.1/dns-query",
  "useragent_file": "useragents.txt",
  "tls_fingerprint": "chrome"
}
EOF
echo " config.json created"
fi

# Sample proxylist.txt
if [ ! -f "proxylist.txt" ]; then
cat > proxylist.txt << 'EOF'
# VEKO DOME Proxy List
# Format: IP:PORT or IP:PORT:USER:PASS
# Prefix with 'socks5://' for SOCKS5 proxies

# Example:
# 198.51.100.42:8080
# 198.51.100.43:3128:user:pass
# socks5://127.0.0.1:9050
EOF
echo " proxylist.txt created"
fi

# Sample useragents.txt
if [ ! -f "useragents.txt" ]; then
cat > useragents.txt << 'EOF'
# VEKO DOME User Agent List
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) Gecko/20100101 Firefox/114.0
Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36
Mozilla/5.0 (iPhone; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1
EOF
echo " useragents.txt created"
fi

# Done
echo ""
echo " VEKO DOME build completed successfully."
echo " Output saved in ./builds/"
echo ""
