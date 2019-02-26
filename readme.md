#使用go module
export GO111MODULE=on

#大杀器
export GOPROXY=https://goproxy.io

#初始化该模块
go mod init

go get -m [packages]
#添加缺失的模块以及移除不需要的模块
go mod tidy -v
#检查当前模块的依赖是否全部下载下来
go mod verify
#生成vendor文件夹，该文件夹下将会放置你go.mod文件描述的依赖包
go mod vendor