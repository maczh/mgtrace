# MgTrace 轻量级微侵入Gin微服务链路跟踪工具

## 使用说明

> 在大型系统的微服务化构建中，一个系统会被拆分成许多模块。这些模块负责不同的功能，组合成系统，最终可以提供丰富的功能。在这种架构中，一次请求往往需要涉及到多个服务。互联网应用构建在不同的软件模块集上，这些软件模块，有可能是由不同的团队开发、可能使用不同的编程语言来实现、有可能布在了几千台服务器，横跨多个不同的数据中心，也就意味着这种架构形式也会存在一些问题：
    
    + 如何快速发现问题？
    + 如何判断故障影响范围？
    + 如何梳理服务依赖以及依赖的合理性？
    + 如何分析链路性能问题以及实时容量规划？
    
> 分布式链路追踪（Distributed Tracing），就是将一次分布式请求还原成调用链路，进行日志记录，性能监控并将 一次分布式请求的调用情况集中展示。比如各个服务节点上的耗时、请求具体到达哪台机器上、每个服务节点的请求状态等等。
    
## 实现原理
    
    在全微服务调用链路中，通过在http请求的header中加入 X-Request-Id 参数进行传递与追踪
    
## 安装
```shell script
go get -u github.com/maczh/mgtrace
```        

## 使用

### 在gin引擎初始化时挂载MgTrace中间件
```go
	engine := gin.Default()

	//添加跟踪日志
	engine.Use(mgtrace.TraceId())
```

### 在调用其他微服务时,http header中添加X-Request-Id参数
```go
	header := map[string]string{"X-Request-Id": GetRequestId()}
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Data:    params,
		Headers: header,
	})
```

**若使用MgCall进行微服务调用，则自动实现，无需额外处理**
