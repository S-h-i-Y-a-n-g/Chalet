package entity

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type Kafka struct {
	Bootstrap string `yaml:"bootstrap"`
	Topic     string `yaml:"topic"`
	ClusterID string `yaml:"clusterId"`
	AppID     string `yaml:"appId"`
	Password  string `yaml:"password"`
	Group     string `yaml:"group"`
}

type QQMail struct {
	UserName          string `yaml:"user_name"`
	AuthorizationCode string `yaml:"authorization_code"`
	Port              int    `yaml:"port"`
	Host              string `yaml:"host"`
	IsSsl             bool   `yaml:"is-ssl"`
	Nickname          string `yaml:"nickname"`
}
