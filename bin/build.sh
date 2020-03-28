#!/bin/bash
cd ../src
go build

if [ $? != 0 ]; then
    exit $?
fi

mv ./src ../bin/ludwig
echo "Compiled:"
find . -name '*.go' | xargs wc -l
