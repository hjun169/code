程序说明
========
1.net_application为对应代码根目录，需要复制example.env为.env文件，配置见注释

2.log为日志记录文件所在目录 

3.utils为程序基础组件

4.tcp,rpc和websocket为执行文件目录，进入该目录下执行go run server.go&和go run client go则可运行

5.websocket需要在$GOPATH或其他具有全局可见目录下创建golang.org\x目录，然后下载对应文件

(1) git init 

(2) git clonet https://github.com/golang/net.git

或在github上直接下载对应文件复制到golang\x
