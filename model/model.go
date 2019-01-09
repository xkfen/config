package model

type Configuration struct {
	// 环境
	Env      string `json:"env"`
	// mysql host
	Host     string `json:"host"`
	// mysql 密码
	Password string `json:"password"`
	// mysql 用户名
	UserName string `json:"user_name"`
	// mysql 数据库前缀
	Prefix   string `json:"prefix"`
	// mysql 数据库名称
	Name     string `json:"name"`
	// mysql dialect
	Type string `json:"type"`
	// 最大空闲连接数量
	MaxIdleConn int `json:"max_idle_conn"`
	// 
	MaxOpenConn int `json:"max_open_conn"`
	// port
	Port int `json:"port"`
}
