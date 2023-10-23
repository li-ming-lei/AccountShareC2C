#!/bin/sh
DTTERM=`pgrep AccountShareC2C`       #appName为进程名
mkdir -p /root/AccountShareC2C/log
DATE=`date '+%Y%m%d-%H%M%S'`
echo $DATE  >> /root/AccountShareC2C/log/watch_dog.log
 if [ -n "$DTTERM" ]
 then  
    echo "app service is ok" >> /root/AccountShareC2C/log/watch_dog.log
 #正确输入信息到日志文件
 else
    echo "app servicie not exist, try start it" >> /root/AccountShareC2C/log/watch_dog.log
    cd /root/AccountShareC2C
    ./bin/AccountShareC2C & > /root/AccountShareC2C/log/service.log
 fi