#! /bin/env bash

set -eu pipefail

# get os
get_os() {
  if [ "$(uname -s)" = "Darwin" ]; then
    echo darwin | tr -cd '[:print:]'
  elif [ "$(uname -s)" = "Linux" ]; then
    echo linux | tr -cd '[:print:]'
  else
    echo "MacOS and Linux are supported as of now." >&2
    exit 1
  fi
}

# get arch
get_arch() {
  if [ "$(uname -m)" = "x86_64" ]; then
    echo amd64 | tr -cd '[:print:]'
  elif [ "$(uname -m)" = "arm64" ]; then
    echo arm64 | tr -cd '[:print:]'
  else
    echo "x86_64 and arm64 architectures are supported as of now." >&2
    exit 1
  fi
}

# get current version
get_current_version() {
  curl -sIq https://github.com/SymmetricalAI/symctl/releases/latest | grep location | cut -d" " -f2 | rev | cut -d/ -f1 | rev | tr -cd '[:print:]'
}

# prepare temp directory
prepare_temp_dir() {
  mktemp -d /tmp/symctl.XXXXXX
}

# prepare url to download
prepare_url() {
  local version="$1"
  local os="$2"
  local arch="$3"
  echo "https://github.com/SymmetricalAI/symctl/releases/download/$version/symctl-$version-$os-$arch"
}

#download to temp directory
download_to_temp_dir() {
  local url="$1"
  local tmp_dir="$2"
  curl -sLo "$tmp_dir/symctl" "$url"
  chmod +x "$tmp_dir/symctl"
}

# ensure ~/.symctl directory with underlying bin and plugins directories exists
ensure_symctl_dir_exists() {
  mkdir -p ~/.symctl/bin
  mkdir -p ~/.symctl/plugins
}

# backup old version of binary in ~/.symctl/bin/symctl.old if ~/.symctl/bin/symctl exists
backup_old_version() {
  if [ -f ~/.symctl/bin/symctl ]; then
    mv ~/.symctl/bin/symctl ~/.symctl/bin/symctl.old
  fi
}

# move symctl to ~/.symctl/bin
move_symctl_to_bin() {
  local tmp_dir="$1"
  mv "$tmp_dir/symctl" ~/.symctl/bin/symctl
}

if [ -z "$BASH_VERSION" ]; then
    echo "Error: This script requires Bash." >&2
    exit 1
fi

echo "Installing symctl..."

if [ -n "${DEBUG:-}" ]; then
  echo "DEBUG is enabled."
fi

if [ -n "${DEBUG:-}" ]; then
  echo "will get OS"
fi
OS=$(get_os)
if [ -n "${DEBUG:-}" ]; then
  echo "got OS: $OS"
  echo "will get ARCH"
fi
ARCH=$(get_arch)
if [ -n "${DEBUG:-}" ]; then
  echo "got ARCH: $ARCH"
  echo "will get VERSION"
fi
VERSION=$(get_current_version)
if [ -n "${DEBUG:-}" ]; then
  echo "got VERSION: $VERSION"
  echo "will prepare URL"
fi
URL=$(prepare_url "$VERSION" "$OS" "$ARCH")
if [ -n "${DEBUG:-}" ]; then
  echo "got URL: $URL"
  echo "will prepare temp dir"
fi
TMP_DIR=$(prepare_temp_dir)
if [ -n "${DEBUG:-}" ]; then
  echo "got TMP_DIR: $TMP_DIR"
  echo "will ensure ~/.symctl dir exists"
fi
ensure_symctl_dir_exists
if [ -n "${DEBUG:-}" ]; then
  echo "ensured ~/.symctl dir exists"
  echo "will download binary to temp dir"
fi
download_to_temp_dir "$URL" "$TMP_DIR"
if [ -n "${DEBUG:-}" ]; then
  echo "downloaded binary to temp dir"
  echo "will backup old version"
fi
backup_old_version
if [ -n "${DEBUG:-}" ]; then
  echo "backed up old version"
  echo "will move symctl to ~/.symctl/bin"
fi
move_symctl_to_bin "$TMP_DIR"
if [ -n "${DEBUG:-}" ]; then
  echo "moved symctl to ~/.symctl/bin"
fi

echo "symctl $VERSION has been installed successfully."
echo "You can now use symctl command to interact with SymmetricalAI platform."
echo "Please make sure ~/.symctl/bin is in your PATH."
echo "You can also use symctl --help to see the available commands."
echo "For more information, please visit: https://github.com/SymmetricalAI/symctl"
echo "Happy hacking!"
