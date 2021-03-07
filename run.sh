#!/bin/bash
# @Author: gongwen
# @Date:   2019-02-01 20:12:03
# @Last Modified by:   gongwen
# @Last Modified time: 2019-02-01 20:12:03


case $1 in
start)
nohup ./ApiManager 2>&1 >> run.log 2>&1 /dev/null &
echo "服务已启动..."
sleep 1
;;
stop)
killall ApiManager
echo "服务已停止..."
sleep 1
;;
restart)
killall ApiManager
sleep 1
nohup ./ApiManager 2>&1 >> run.log 2>&1 /dev/null &
echo "服务已重启..."
sleep 1
;;
*)
echo "$0 {start|stop|restart}"
exit 4
;;
esac