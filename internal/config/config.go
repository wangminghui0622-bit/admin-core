package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DatabaseDSN string
	RedisConf   struct {
		Host string
		Type string
		Pass string
	}
	AccessSecret string
	AccessExpire int64
	McmsRpc      zrpc.RpcClientConf
}
