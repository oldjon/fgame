#!/bin/sh

start()
{
  nohup $PWD/bin/gatewayserver -config=$PWD/bin/config.json &
}

stop()
{
  serverlist='gatewayserver'
  for ser in $serverlist
  do
    echo "stop $ser"
    ps aux|grep "/$ser"|sed -e "/grep/d"|awk '{print $2}'|xargs kill 2&>/dev/null
    while test -f run.sh
    do
      count=`ps x|grep -w $ser|sed -e '/grep/d'|wc -l`
      if [ $count -eq 0 ];then
        break
      fi
      sleep 1
    done
  done
  echo "running server:"`ps x|grep "server -c"|sed -e '/grep/d'|wc -l`
}

case $1 in
stop)
  stop
;;
start)
  start
;;
*)
  stop
  sleep
  start
esac