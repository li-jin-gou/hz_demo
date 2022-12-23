#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=hellotest
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}