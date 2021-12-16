package config

const (
	ResourceApiGateway = "http://42.192.53.10:8888"
	// HttpGetResource 获得我的资源
	HttpGetResource = ResourceApiGateway + "/api/resources"
	// HttpGetCode 获取验证码
	HttpGetCode = "https://ttchain.tntlinking.com/api/authorization/verification/code"
	// HttpLogin 登陆
	HttpLogin = "https://ttchain.tntlinking.com/api/authorization/oauth/token"

	//Port p2p端口配置
	Port = 4001

	// p2p swarm key
	SwarmKey = "/key/swarm/psk/1.0.0/\n/base16/\n55158d9b6b7e5a8e41aa8b34dd057ff1880e38348613d27ae194ad7c5b9670d7"
)
