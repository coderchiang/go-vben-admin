#!/bin/bash

### 构建要求
### 1. go >= 1.13
### 2. node >= 8

baseDir=`echo $PWD`
webDir=`echo $baseDir/web`
dataDir=`echo $baseDir/data`


funcBuildServer() {
    echo 'server module building...'
    export GOPROXY=https://goproxy.cn
    export GO111MODULE=on
    cd $baseDir
    go mod tidy
    GOOS=linux GOARCH=amd64 go build -o $baseDir/go-vben-admin
    if [  -f $baseDir/go-vben-admin ]; then
        echo 'server module build  finished'
    else
        echo "server module build false" 1>&2
        exit 1
    fi


}



funcBuildSite() {
    echo 'web module building...'
    cd $webDir
    yarn  install --registry https://registry.npm.taobao.org/
    yarn build
    if [ -d $webDir/dist ]; then
        echo 'site module build  finished'
    else
        echo "site module build false" 1>&2
        exit 1
    fi
}

funcTouchDir() {
    if [ ! -d "$1" ]; then
        mkdir -p $1
    fi
}

funcInstallDokcer(){
    sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
sudo yum install -y yum-utils
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
sudo yum -y install docker-ce docker-ce-cli containerd.io
sudo systemctl enable docker
sudo systemctl start docker
}

funcDockerInitMysqlAndRedis(){
    #启动mysql
    docker run -p 3306:3306 --name mysql  -v /data/mysql:/var/lib/mysql  -v $dataDir/*.sql:/docker-entrypoint-initdb.d   -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
    #启动redis
    docker run -p 6379:6379 --name redis -v /data/docker/redis/redis.conf:/etc/redis/redis.conf  -v /data/docker/redis/data:/data -d redis redis-server /etc/redis/redis.conf --appendonly yes
}
  funcStartServer(){
      #start server
  cd  $baseDir
    nohup ./go-vben-admin &
  }

funcBuildServer
funcBuildSite
funcInstallDokcer
funcDockerInitMysqlAndRedis
funcStartServer




