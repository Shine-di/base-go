#!/usr/bin/expect -f
set user root
set host 120.77.39.56
set password 2015@Shinedi
set timeout -1


git add -A
git commit -m  "update"
git push

ssh  $user@$host
expect "*:*"
send "$password\r"
interact
expect eof

cd python/chat

git pull

docker build -t everyday_wechat:v1 .
docker run everyday_wechat:v1  -d