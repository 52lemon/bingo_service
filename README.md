# GO 环境配置

> 环境变量

    # 安装目录
    export GOROOT=/usr/local/go            
    # 工作环境  
    export GOPATH=$HOME/go     
    # 可执行文件存放
    export GOBIN=$GOPATH/bin           
    # 添加PATH路径
    export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH      
    # 开启moudle 
    export GO111MODULE=auto


# 创建Gin项目

> 1 替换导入路径为自己的项目名称

    sed -i "s/bingo_service/{自定义项目名}/g" `grep bingo_service -rl ./*`

> 2 下载依赖

    make all

# 运行项目

> dev run

	nohup go run main.go  > bingo_service_`date +%y%m%d`.log 2>&1 &

> build

	GOOS=linux GOARCH=amd64 go build -o ./cmd/bingo_service main.go
	
	# 编译之后执行文件
	nohup ./cmd/bingo_service  > /dev/null 2>&1 &
	或者
	nohup ./cmd/bingo_service  > ./logs/bingo_service_190702.log 2>&1 &
	
> stop

    kill -9 $(pgrep bingo_service)
	
# NSQ 安装

> 下载

    wget https://s3.amazonaws.com/bitly-downloads/nsq/nsq-1.1.0.linux-amd64.go1.10.3.tar.gz
    
> 解压

    tar -zxvf nsq-1.1.0.linux-amd64.go1.10.3.tar.gz

> 启动服务

    cd nsq-1.1.0.linux-amd64.go1.10.3/bin/
    nohup ./nsqlookupd > /dev/null 2>&1 &
    nohup ./nsqd --lookupd-tcp-address=127.0.0.1:4160 > /dev/null 2>&1 &
    nohup ./nsqadmin --lookupd-http-address=127.0.0.1:4161 > /dev/null 2>&1 &

