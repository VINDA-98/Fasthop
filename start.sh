#!/bin/bash
# 重启jenkins的脚本web服务

closeport=10008

name=$(lsof -i:$closeport|tail -1|awk '"$1"!=""{print $2}')
if [ -z $name ]
then
      echo "No process can be used to killed!"
      #重新启动服务  后台挂起工具服务
      nohup ./Fasthop >./out.log 2>&1 &
      n=1
      a=`uname  -a`
      b="Darwin"
      echo "正在检测服务是否启动"
      while true
      do
        if test $n -gt 10
        then
                echo "服务启动失败"
                break
        fi
        # 每三秒检测工具服务是否启动
        sleep 3
        n=$(($n+1))
        if [[ $a =~ $b ]]
        then
                echo "mac正在监听服务启动端口"
                port=`netstat -anvp tcp |grep 10008`   #端口
        else
                echo "正在监听服务启用端口"
                port=`netstat -antp | grep 10008`      #端口
        fi

        if [ ${#port} -gt 3 ]
        then
                echo "服务已经重新启动"
                break;
        fi
      done

else
    echo "服务已经启动"
fi