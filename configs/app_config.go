package configs

type Cors struct {
	AllowOrigins     []string `yaml:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

type DB struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	Port     int    `yaml:"port"`
	SSLMode  bool   `yaml:"sslmode"`
}

type Fiber struct {
	Prefork       bool `yaml:"prefork"`
	StrictRouting bool `yaml:"strict_routing"`
	CaseSensitive bool `yaml:"case_sensitive"`
	Port          int  `yaml:"port"`
}

type Migrations struct {
	Mode string `yaml:"mode"`
}

type AppConfig struct {
	Cors       Cors       `yaml:"cors"`
	Database   DB         `yaml:"db"`
	Migrations Migrations `yaml:"migrations"`
	Fiber      Fiber      `yaml:"fiber"`
}
