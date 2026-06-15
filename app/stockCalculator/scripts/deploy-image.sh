#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="${IMAGE_NAME:-stockcalculator-webapp:latest}"
TARGET_PLATFORM="${TARGET_PLATFORM:-linux/amd64}"
SERVER_USER="${SERVER_USER:-root}"
SERVER_PORT="${SERVER_PORT:-22}"
REMOTE_IMAGE_NAME="${REMOTE_IMAGE_NAME:-$IMAGE_NAME}"

if [[ -z "${SERVER_IP:-}" ]]; then
  echo "ERROR: SERVER_IP is required" >&2
  exit 1
fi

if [[ -z "${SERVER_PASSWORD:-}" ]]; then
  echo "ERROR: SERVER_PASSWORD is required" >&2
  exit 1
fi

if ! command -v docker >/dev/null 2>&1; then
  echo "ERROR: docker is not installed or not in PATH" >&2
  exit 1
fi

if ! command -v sshpass >/dev/null 2>&1; then
  echo "ERROR: sshpass is required for password-based SSH" >&2
  echo "Install on macOS: brew install hudochenkov/sshpass/sshpass" >&2
  echo "Install on Debian/Ubuntu: apt-get install sshpass" >&2
  exit 1
fi

echo "Building $IMAGE_NAME for $TARGET_PLATFORM..." >&2
docker build --platform "$TARGET_PLATFORM" -t "$IMAGE_NAME" .

REMOTE="${SERVER_USER}@${SERVER_IP}"
SSH_BASE=(
  sshpass -p "$SERVER_PASSWORD"
  ssh
  -p "$SERVER_PORT"
  -o StrictHostKeyChecking=accept-new
  -o UserKnownHostsFile="$HOME/.ssh/known_hosts"
)

echo "Pushing $IMAGE_NAME to $REMOTE as $REMOTE_IMAGE_NAME..."
docker save "$IMAGE_NAME" \
  | gzip \
  | "${SSH_BASE[@]}" "$REMOTE" "gzip -d | docker load"

if [[ "$REMOTE_IMAGE_NAME" != "$IMAGE_NAME" ]]; then
  "${SSH_BASE[@]}" "$REMOTE" "docker tag '$IMAGE_NAME' '$REMOTE_IMAGE_NAME'"
fi

echo "Done."
