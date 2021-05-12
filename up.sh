#!/bin/bash


baseDir=`echo $PWD`
webDir=`echo $baseDir/web`
dataDir=`echo $baseDir/data`
goVersion=`echo "go1.16.4.linux-amd64"`
nodeVersion=`echo "node-v14.17.0-linux-x64"`

funcInstallGo(){
  cd /usr/local

 wget https://golang.google.cn/dl/$goVersion.tar.gz
 tar  -zxvf $goVersion.tar.gz
 rm -rf $goVersion.tar.gz
 ln -s /usr/local/go/bin/go  /usr/bin/go
 ln -s /usr/local/go/bin/gofmt  /usr/bin/gofmt
}
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

funcInstallNodejs(){
  cd /usr/local
wget https://nodejs.org/dist/v14.17.0/$nodeVersion.tar.xz
tar xvJf $nodeVersion.tar.xz
rm -f  $nodeVersion.tar.xz
mv $nodeVersion node
ln -s /usr/local/node/bin/node  /usr/bin/node
ln -s /usr/local/node/bin/npm  /usr/bin/npm
ln -s /usr/local/node/bin/npx  /usr/bin/npx
npm install yarn -g
ln -s /usr/local/node/bin/yarn /usr/bin/yarn
ln -s /usr/local/node/bin/yarnpkg /usr/bin/yarnpkg
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
    nohup ./go-vben-admin &2>&1
    echo 'server start  success'
    echo 'listen port:80'
  }
  #安装go环境脚本
funcInstallGo
#go build 服务端
funcBuildServer
#安装nodejs环境
#funcInstallNodejs

#nodejs编译前端
#funcBuildSite

#安装docker 环境
funcInstallDokcer
#初始化docker中Mysql和Redis环境
funcDockerInitMysqlAndRedis
#启动server
funcStartServer




