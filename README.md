# prometheus
This repository can help u to deploy prometheus-operator、prometheus，the code in cmd/statistic is an example that tell you how to statistic prometheus metrics with golang.

## prase metrics
更改metrics的信息,并通过api接口返回
#### 运行代码
```go
go run cmd/prasemetric/main.go
```

#### 查看metrics
```go
# 浏览器打开 http://localhost:8080/metrics
go_info{version="go1.16.2"} 1
```

#### 查看更改
```go
# http://localhost:8080/get

go_info{whoami="go1.16.2"} 1
```