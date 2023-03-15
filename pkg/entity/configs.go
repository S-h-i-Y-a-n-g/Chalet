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


type Jwt struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
