# gowechat
一款由Go语言开发的微信公众号后台框架，支持被动消息回复/安全模式加密解密/客服消息/自定义菜单/定时获取Token

# 使用方法

# 示例
```

	//新建服务器
	mywechat, err := wechat.New("xx", "xx", "xx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mywechat)

	//开启Token刷新  1小时一次
	go mywechat.Token_Refresh()

	//菜单初始化
	//客服创建
	rerr := mywechat.Service_Add()
	if rerr != nil {
		fmt.Println(rerr)
	}
	//自定义菜单创建
	rerr = mywechat.Manu_Create()
	if rerr != nil {
		fmt.Println(rerr)
	}

	
```

# 

# 