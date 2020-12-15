# ruijie_login

锐捷portal认证工具

锐捷网页认证登录工具，用来自动认证，掉线重连，~~挤别人的号~~

## 说明


```
$ ./ruijie_login -h
Usage of ruijie_login:
    -config string
        config file (default "config.json")
```

跑一下，若配置文件不存在，就在当前目录生成默认配置

```
{
	"username": "username here",
	"password": "password here",
	"ttlInterval": 30, //检测网络连通的间隔(s)
	"retryInterval": 30 //认证失败后重试的间隔(s)
}
```

如果配置文件存在就会开始跑，用baidu测认证是否正常，若被挤掉线了会重连