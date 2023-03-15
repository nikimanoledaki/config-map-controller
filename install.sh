#!/bin/sh -e

WORKSPACE="/usr/local"
TAG=${version:=v0.3.0-rc.1}
VERSION=$(echo "$TAG" | tr -d v)
REPO=${repo:=open-component-model/ocm}
PLATFORM=linux-amd64
ARCHIVESUFFIX=.tar.gz
ARCHIVEFILE="ocm-$VERSION-$PLATFORM$ARCHIVESUFFIX"
URL="https://github.com/$REPO/releases/download/$TAG/$ARCHIVEFILE"
TARGET=${WORKSPACE}/bin/ocm

cd /tmp
echo "Install Open Component Model CLI Tool version $TAG from $URL"
rm -f ocm-cli.tgz
mkdir -p "$(dirname "$TARGET")"
curl -L -o ocm-cli.tar.gz "$URL"
tar -xvzf ocm-cli.tar.gz
cp ocm "$TARGET"
chmod a+x "$TARGET"
echo "ocm installed into $TARGET"
echo "ocm-path=$TARGET" >> "$GITHUB_OUTPUT"