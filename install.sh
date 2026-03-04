#!/usr/bin/env bash
set -euo pipefail

REPO="m0n0x41d/claude-code-fpf"
SKILL_DIR="${HOME}/.claude/skills/fpf"
REF_DIR="${SKILL_DIR}/references"
BINARY="${REF_DIR}/fpf-rag"
VERSION_FILE="${REF_DIR}/.version"

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

# Get latest release info
RELEASE_JSON=$(curl -sfL "https://api.github.com/repos/${REPO}/releases/latest")
LATEST_TAG=$(echo "$RELEASE_JSON" | grep '"tag_name"' | head -1 | cut -d '"' -f 4)
RELEASE_URL=$(echo "$RELEASE_JSON" | grep "browser_download_url.*${BINARY_NAME}" | head -1 | cut -d '"' -f 4)

if [ -z "$RELEASE_URL" ] || [ -z "$LATEST_TAG" ]; then
  echo "Error: No release found for ${BINARY_NAME}"
  echo "Check https://github.com/${REPO}/releases for available binaries."
  exit 1
fi

# Check if already up to date
CURRENT_VERSION=""
if [ -f "$VERSION_FILE" ]; then
  CURRENT_VERSION=$(cat "$VERSION_FILE")
fi

if [ "$CURRENT_VERSION" = "$LATEST_TAG" ] && [ -x "$BINARY" ]; then
  echo "Already up to date (${LATEST_TAG})."
  exit 0
fi

if [ -n "$CURRENT_VERSION" ]; then
  echo "Updating claude-code-fpf: ${CURRENT_VERSION} → ${LATEST_TAG}"
else
  echo "Installing claude-code-fpf ${LATEST_TAG}"
fi
echo "  Platform: ${OS}/${ARCH}"

# Create directories
mkdir -p "$REF_DIR"

# Download binary
echo "  Downloading ${BINARY_NAME}..."
curl -sfL "$RELEASE_URL" -o "${BINARY}.tmp"
chmod +x "${BINARY}.tmp"
mv -f "${BINARY}.tmp" "$BINARY"

# Download skill file
echo "  Downloading skill file..."
curl -sfL "https://raw.githubusercontent.com/${REPO}/main/skill/fpf/SKILL.md" \
  -o "${SKILL_DIR}/SKILL.md"

# Clean up legacy skill file name
rm -f "${SKILL_DIR}/fpf.md"

# Record installed version
echo "$LATEST_TAG" > "$VERSION_FILE"

echo ""
echo "Done! ${LATEST_TAG}"
echo "  Skill:  ${SKILL_DIR}/SKILL.md"
echo "  Binary: ${BINARY}"
echo ""
echo "The /fpf skill is now available in Claude Code."
