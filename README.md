# RS | 网络请求使用库

###### `RS`: `Requests`的缩写

让 `golang` 更简单的发起 `http` 请求

链式的请求规则

Convenient http client for go.

# HOW TO USE | 怎么用
---
```shell
go get github.com/Jecced/rs
```

# Usage | 用法
---
## Simple Case | 简单案例
---
一个简单的http请求示例, 示例执行http请求, 并将响应读取为字符串 

One simple http request example that do http get request and read response as string:
```go
resp := rs.Get("http://www.baidu.com/").Send().ReadText()
fmt.Println(resp)
```

# WriteToFile | 写入到文件
将响应请求保存到本地文件中, 例如保存一个图片

Save the response request to a local file, such as an image
```go
rs.Get("http://www.baidu.com/img/sug_bd.png?v=09816787.png").
		Send().
		SetTimeOut(30_000).
		WriteToFile("/Users/ankang/develop/test/1.png")
```

# Timeout | 设置超时时间
---
你可以设置连接超时时间, 和响应请求的超时时间

You can set connection connect timeout, and socket read/write timeout value
```go
rs.Get(url).SetConnTimeOut(30_000).SetRespTimeOut(30_000).Send().ReadText()
```

你也可以同时修改他们

You can also change their values at the same time
```go
rs.Get(url).SetTimeOut(30_000).Send().ReadText()
```

# Proxy | 代理
---
通过代理方法设置代理

Set proxy by proxy method
```go
resp := rs.Get("http://www.google.com").
    Proxy("127.0.0.1:1081").
    Send().ReadText()
fmt.Println(resp)
```




## LICENSE

    MIT License
    
    Copyright (c) 2020 Jecced
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.