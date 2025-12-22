#!/bin/sh

# Ensure we are in the /app folder
cd /app || exit

# If we aren't running as root, just exec the CMD
if [ "$(id -u)" -ne 0 ] ; then
  exec "$@"
fi

DOCPORT_UID=${DOCPORT_UID:-1000}
DOCPORT_GID=${DOCPORT_GID:-1000}

# Validate that DOCPORT_UID and DOCPORT_GID are numeric
case "$DOCPORT_UID" in
  ''|*[!0-9]*)
    echo "invalid DOCPORT_UID: must be a numeric UID" >&2
    exit 1
    ;;
  *)
    # UID is valid
    ;;
esac
case "$DOCPORT_GID" in
  ''|*[!0-9]*)
    echo "invalid DOCPORT_GID: must be a numeric GID" >&2
    exit 1
    ;;
  *)
    # GID is valid
    ;;
esac

# Check if the group exists; if not, create it
if ! getent group docport-io-group > /dev/null 2>&1; then
  echo "creating group $DOCPORT_GID..."
  addgroup -g "$DOCPORT_GID" docport-io-group
fi

# Check if the user exists; if not, create it
if ! id -u docport-io > /dev/null 2>&1; then
  if ! getent passwd "$DOCPORT_UID" > /dev/null 2>&1; then
    echo "creating user $DOCPORT_UID..."
    if ! adduser -D -u "$DOCPORT_UID" -G docport-io-group docport-io > /dev/null; then
      echo "failed to create user with UID $DOCPORT_UID in group docport-io-group" >&2
      exit 1
    fi
  else
    existing_user=$(getent passwd "$DOCPORT_UID" | cut -d: -f1)
    echo "using existing user: $existing_user"
  fi
fi

mkdir -p /app/data

# Change ownership of the /app/data directory
find /app/data \( ! -group "${DOCPORT_GID}" -o ! -user "${DOCPORT_UID}" \) -exec chown "${DOCPORT_UID}:${DOCPORT_GID}" {} +

# Switch to the non-root user
exec su-exec "$DOCPORT_UID:$DOCPORT_GID" "$@"
