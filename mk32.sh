#!/bin/sh

pwd |grep -q 'github.com/ziutek/matrix$' || exit 1

rm -rf matrix32
mkdir matrix32
cp *.go matrix32
cd matrix32
sed -i 's/float64/float32/g' *.go
sed -i 's/ackage matrix/ackage matrix32/g' *.go
sed -i 's/matrix:/matrix32:/g' *.go