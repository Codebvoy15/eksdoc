cat <<'EOF' > install.sh
#!/usr/bin/env bash
set -e

REPO="Codebvoy15/eksdoc"
BINARY="eksdoctor"

OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" ]]; then
  ARCH="arm64"
fi

URL="https://github.com/$REPO/releases/latest/download/${BINARY}-${OS}-${ARCH}"

echo "Downloading $URL"
curl -fL "$URL" -o "$BINARY"

chmod +x "$BINARY"
sudo mv "$BINARY" /usr/local/bin/

echo "eksdoctor installed successfully"
EOF
