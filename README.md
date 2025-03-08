# dlsite_api

基于dlsite API的信息获取工具

## 函数：

``` go
data, err := dlsite_api.GetInfoBySearch([名称], [类型], client)
```

该函数用于搜索条目，类型只能是以下字符串：

```
游戏、同人、漫画、手机游戏
```

如果指定上述以外的字符串，则可能调用失败。

返回一个[]byte和一个error，如果err为nil，则没有错误。

=================================================================

``` go'
data, err := dlsite_api.GetInfoById("VJ013799", "游戏", client)
```

该函数用于查看条目，类型只能是以下字符串：

```
游戏、同人、漫画、手机游戏
```

如果指定上述以外的字符串，则可能调用失败。

返回一个[]byte和一个error，如果err为nil，则没有错误。

## 关于client

指定模版如下：

``` go
proxyURL := "" // 替换为你的代理地址
proxy, err := url.Parse(proxyURL)
if err != nil {
	panic(err)
}
client := &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyURL(proxy),
	},
}
```

