#!/bin/sh
#export LD_LIBRARY_PATH=/opt/glibc2.14/lib:$LD_LIBRARY_PATH
#go build -a -o file_pool_api main.go
kill -9 $(pidof /var/gowww/src/file_pool_api/file_pool_api)
nohup /var/gowww/src/file_pool_api/file_pool_api -c /var/gowww/src/file_pool_api/conf/config.yaml -la /var/gowww/src/file_pool_api/conf/log-app.xml > /var/gowww/src/file_pool_api/sys.log  2>&1 &
