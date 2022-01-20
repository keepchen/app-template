#### httpserver

- 接口文档地址
```text  
# 注意，当config.toml中的debug为false时，接口文档将不再允许访问
http://localhost:[端口地址]/swagger/index.html
```  
- 性能检测(仅debug模式下开始)
```shell script
#1.首先启动服务
go run main.go httpserver -c path/to/you/config

#2.启动go tool监听
go tool pprof --seconds 30 http://127.0.0.1:8080/debug/pprof/profile

#3.开启压力测试，ab或wrk工具都可以，压测对应的路由方法
wrk -t12 -c800 -d30s http://127.0.0.1:8080/api/v1/sayhello

#4.在第2步骤完成后，会进入pprof终端，输入web，即可浏览检测结果
#4.1.此命令需要`graphviz`工具的支持，根据自身操作系统环境进行安装
```