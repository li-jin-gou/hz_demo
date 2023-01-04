#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=test
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}