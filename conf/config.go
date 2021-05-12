package conf

type Config struct {
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Jwt     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Redis   Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	BaiduMap BaiduMap `mapstructure:"baidumap" json:"baidumap" yaml:"baidumap"`
}

type System struct {
	DBType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	CacheType string `mapstructure:"cache-type" json:"cache-type" yaml:"cache-type"`
	Env   string `mapstructure:"env" json:"env" yaml:"env"`
	Port  string `mapstructure:"port" json:"port" yaml:"port"`
}

type Log struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"dbname" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}

type Redis struct {
	Addr   string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	DB int `mapstructure:"db" json:"db" yaml:"db"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}

type BaiduMap struct {
	IpLocationUrl string `mapstructure:"ip_location_url" json:"ip_location_url" yaml:"ip_location_url"`
	AK     string  `mapstructure:"ak" json:"ak" yaml:"ak"`
}