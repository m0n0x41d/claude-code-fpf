#!/usr/bin/env bash
set -euo pipefail

REPO="m0n0x41d/claude-code-fpf"
SKILL_DIR="${HOME}/.claude/skills/fpf"
REF_DIR="${SKILL_DIR}/references"

# Detect OS and arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64) ARCH="arm64" ;;
  arm64)   ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

BINARY_NAME="fpf-rag-${OS}-${ARCH}"

echo "Installing claude-code-fpf..."
echo "  Platform: ${OS}/${ARCH}"

# Get latest release URL
RELEASE_URL=$(curl -sfL "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep "browser_download_url.*${BINARY_NAME}" \
  | head -1 \
  | cut -d '"' -f 4)

if [ -z "$RELEASE_URL" ]; then
  echo "Error: No release found for ${BINARY_NAME}"
  echo "Check https://github.com/${REPO}/releases for available binaries."
  exit 1
fi

# Create directories
mkdir -p "$REF_DIR"

# Download binary
echo "  Downloading ${BINARY_NAME}..."
curl -sfL "$RELEASE_URL" -o "${REF_DIR}/fpf-rag"
chmod +x "${REF_DIR}/fpf-rag"

# Download skill file
echo "  Downloading skill file..."
curl -sfL "https://raw.githubusercontent.com/${REPO}/main/skill/fpf.md" \
  -o "${SKILL_DIR}/fpf.md"

echo ""
echo "Installed successfully!"
echo "  Skill: ${SKILL_DIR}/fpf.md"
echo "  Binary: ${REF_DIR}/fpf-rag"
echo ""
echo "The /fpf skill is now available in Claude Code."
