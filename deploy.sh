#!/bin/sh

COMMIT=$1
TAG=$2

rm -rf vendor
govendor init
govendor add +l
govendor add +e

if [ -n "$COMMIT" ]; then
echo Commit 为 $($COMMIT) , 即将push
git add -A
git commit -m ${COMMIT}
git push
git tag ${TAG}
git push origin ${TAG}
rm -rf vendor/
else
echo Commit 为 空 , 即将启动调试 ...
go install
message start -e alphapub
fi

