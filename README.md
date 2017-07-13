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
       export PATH=/usr/local/go/bin
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
       
