#!/bin/bash
cd ../src
go build

if [ $? != 0 ]; then
    exit $?
fi

sudo mv ./src /usr/bin/ludwig
echo "Compiled:"
find . -name '*.go' | xargs wc -l

cp -r ../lib ~/.kgo
