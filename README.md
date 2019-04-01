# go语言的学习资料

  ### linux下go的安装
  
  ##### 下载go
      wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz --no-check-certificate
  ##### 解压文件 并把解压的文件放到 /usr/local 目录下
      tar zxvf go1.8.3.linux-amd64.tar.gz
  ##### 配置环境变量
      vim /etc/profile
  ##### 根目录
      export GOROOT=/usr/local/go
  ##### bin目录
      export GOBIN=$GOROOT/bin
  ##### 工作区  定义访问的目录
       export GOPATH=/home/www
       export PATH=$PATH:$GOPATH:$GOBIN:$GOPATH
  ##### 对修改的文件进行立即生效
       source /etc/profile
  ##### 在/home/www/ 下创建index.go文件
       package main
       
       import "fmt"
       
       func main() {
          /* 这是我的第一个简单的程序 */
          fmt.Println("Hello, World!")
       }
  ##### 输出信息
       go run index.go
  ##### 编译
       go build index.go
        ./index
        
  ##### 编译并压缩文件,并go build 编译的文件节省空间
        go build -ldflags "-s -w" index.go
		
  ##### 压缩可执行文件压缩工具 upx
  * 官方网站 [https://upx.github.io/](https://upx.github.io/)
  * git [https://github.com/upx/upx/releases/tag/v3.94](https://github.com/upx/upx/releases/tag/v3.94)
  * 下载地址[https://github.com/upx/upx/releases/download/v3.94/upx394w.zip](https://github.com/upx/upx/releases/download/v3.94/upx394w.zip)
         
         upx index.go
         
 ##### go语言守护进程
 参考http://www.01happy.com/supervisor-golang-daemon/
 
    ##### GO语言秒杀实战
	百度云盘：https://pan.baidu.com/s/1jbo6UotGW7Kf84ji0PB7Sg  密码：nau4
	腾讯云：https://share.weiyun.com/5NjLpPp                   密码：e9thmp
	
	beego配置
    http://www.tuicool.com/articles/UFbqQv
       
 ##### goweb(书)
    链接地址：https://www.kancloud.cn/kancloud/web-application-with-golang/44105
 ##### go语言圣经（中文版）
    链接地址：http://shouce.jb51.net/gopl-zh/index.html
    链接地址：https://books.studygolang.com/gopl-zh/
 