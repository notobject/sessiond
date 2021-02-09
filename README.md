# 会话管理服务
基于Go+gRPC+Redis的会话管理服务（Demo）

## 环境准备
1. golang安装与配置
> go1.15

2. gRPC安装（cmake）
> gRPC v1.35.0

3. Protobuf安装
> protobuf v3.14.0

4. Docker安装
> Docker Engine - Community v20.10.1

## 运行
```shell
# open an terminal
$ git clone https://github.com/notobject/sessiond.git sessiond
$ cd sessiond
$ export SESSIOND_HOME_PATH=`pwd`
$ make all               # Build all
$ make run_redis_master  # Run redis of master
$ make run_redis_slave   # Run redis of slave
$ cd $SESSIOND_HOME_PATH/build/cmd/sessiond && ./sessiond --config $SESSIOND_HOME_PATH/etc/session_dev.yaml

# change to another terminal
$ cd $SESSIOND_HOME_PATH/build/client/test && ./client_test
``` 

