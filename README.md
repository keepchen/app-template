#### app-template项目说明  

```text
简单的整合了一套(个人认为)较为清爽的项目结构。该项目结构支持单项目多应用方式，打包的
时候all in one，通过sub-command的方式分别启动。不同应用通过注入对应启动命令，在调
用的时候，指定服务/功能即可调用。
```

**文件目录**
```text
.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   ├── httpserver.go
│   ├── root.go
│   └── version.go
├── documents
│   ├── demo_2
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

1.httpserver
```text  
httpserver是...
```  
> [README.md](./documents/httpserver/README.md)

- 调用命名
```shell
# 未编译时
go run main httpserver -c path/to/your/config.toml

# 编译后
app-template httpserver -c path/to/your/config.toml 
```  

2.demo_2
```text  
demo_2是...
```  
> [README.md](./documents/demo_2/README.md)