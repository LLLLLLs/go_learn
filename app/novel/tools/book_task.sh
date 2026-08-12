#!/usr/bin/env sh

set -u

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
BOOK_CREATE="${SCRIPT_DIR}/book_create"
BOOK_UPDATE="${SCRIPT_DIR}/book_update"
DOWNLOADER="${SCRIPT_DIR}/TomatoNovelDownloader"

usage() {
  echo "Usage: $0 {create|update}"
}

ensure_files() {
  [ -f "$BOOK_CREATE" ] || touch "$BOOK_CREATE"
  [ -f "$BOOK_UPDATE" ] || touch "$BOOK_UPDATE"
}

process_create() {
  tmp_remaining=$(mktemp)
  tmp_update=$(mktemp)

  cp "$BOOK_UPDATE" "$tmp_update"

  while IFS= read -r raw_id || [ -n "$raw_id" ]; do
    id=$(printf "%s" "$raw_id" | tr -d '\r')
    [ -n "$id" ] || continue

    if "$DOWNLOADER" --download "$id"; then
      printf "%s\n" "$id" >> "$tmp_update"
      echo "moved: $id (book_create -> book_update)"
    else
      printf "%s\n" "$id" >> "$tmp_remaining"
      echo "download failed, keep in book_create: $id" >&2
    fi
  done < "$BOOK_CREATE"

  awk 'NF && !seen[$0]++' "$tmp_update" > "$BOOK_UPDATE"
  mv "$tmp_remaining" "$BOOK_CREATE"
  rm -f "$tmp_update"
}

process_update() {
  while IFS= read -r raw_id || [ -n "$raw_id" ]; do
    id=$(printf "%s" "$raw_id" | tr -d '\r')
    [ -n "$id" ] || continue

    "$DOWNLOADER" --update "$id"
  done < "$BOOK_UPDATE"
}

main() {
  [ -x "$DOWNLOADER" ] || {
    echo "downloader not found or not executable: $DOWNLOADER" >&2
    exit 1
  }

  ensure_files

  case "${1:-}" in
    create)
      process_create
      ;;
    update)
      process_update
      ;;
    *)
      usage
      exit 1
      ;;
  esac
}

main "$@"
