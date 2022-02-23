# App-Template

#### 项目说明  

```text
简单的整合了一套(个人认为)较为清爽的项目结构。该项目结构支持单项目多应用方式，打包的时候all in one，通过
sub-command的方式分别启动。不同应用通过注入对应启动命令，在调用的时候，指定服务/功能即可调用。
```

- 文件目录  

```text
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── grpcserver.go
│   ├── httpserver.go
│   ├── root.go
│   └── version.go
├── documents
│   ├── demo_2
│   ├── grpcserver
│   └── httpserver
├── go.mod
├── go.sum
├── logs
│   └── running.log
├── main.go
├── pkg
│   ├── app
│   ├── common
│   ├── constants
│   ├── lib
│   └── utils
└── static
    └── certifications
```  

#### 安装

```shell
# 1.下载
git clone https://github.com/keepchen/app-template.git

# 2.安装依赖
go mod tidy

# 3.启动服务
go run main.go httpserver -c pkg/app/httpserver/config/config.sample.toml  
```  

- 浏览器访问  
[http://localhost:8080](http://localhost:8080)

#### 服务说明  

1.httpserver

```text  
httpserver是...
```  

> [README.md](./documents/httpserver/README.md)

- 调用命令

```shell
# 未编译时
go run main.go httpserver -c path/to/your/config.toml

# 编译后
app-template httpserver -c path/to/your/config.toml 
``` 

2.grpcserver

```text  
grpcserver是...
```  

> [README.md](./documents/grpcserver/README.md)

- 调用命令

```shell
# 未编译时
go run main.go grpcserver -c path/to/your/config.toml

# 编译后
app-template grpcserver -c path/to/your/config.toml 
``` 

3.demo_2

```text  
demo_2是...
```  

> [README.md](./documents/demo_2/README.md)