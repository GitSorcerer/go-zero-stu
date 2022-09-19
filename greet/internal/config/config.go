package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	// RedisConf struct {
	// 	Host string // redis地址
	// 	Type string `json:",default=node,options=node|cluster"` // redis类型
	// 	Pass string `json:",optional"`                          // redis密码
	// 	// Weight int64  `json:",optional"`
	// }
}
