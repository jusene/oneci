package config

// consul 地址和端口
type ConsulConfig struct {
	Address string
	Port    int
}

type HostConfig struct {
	Host     string
	Username string
	Password string
	Port     int
}

// 生成yaml的配置
type AppConfig struct {
	Project     string    `yaml:"project"`
	Description string    `yaml:"description"`
	Version     string    `yaml:"version"`
	Author      string    `yaml:"author"`
	Apps        []AppInfo `yaml:"apps"`
}

type AppInfo struct {
	APP        string    `yaml:"name"`
	PORT       int       `yaml:"port"`
	WSPort     []WSInfo  `yaml:"wsport"`
	NFServer   []NFSInfo `yaml:"nfs"`
	FONT       bool      `yaml:"font"`
	RESOURCE   bool      `yaml:"resource"`
	CERT       bool      `yaml:"cert"`
	BIN        bool      `yaml:"bin"`
	SSH        bool      `yaml:"ssh"`
	NFSPath    string    `yaml:"nfspath"`
	Debug      bool      `yaml:"debug"`
	Tag        []string  `yaml:"tag"`
	VERSION    string
	PROJECT    string
	JARVERSION string
	DEBUGPORT  int
	TYPE       string
	ARCH       string
	WS         int
	NFSIP      string
	TIMESTAMP  int64
	ENV        string
}

type NFSInfo struct {
	ENV     string `yaml:"env"`
	Address string `yaml:"address"`
}

type WSInfo struct {
	ENV  string `yaml:"env"`
	PORT int    `yaml:"port"`
}

// 打包docker的模板
type java struct {
	Dockerfile string
	Entrypoint string
}

var Conf = &ConsulConfig{
	Address: "192.168.66.100",
	Port:    8500,
}

var HostConf = &HostConfig{
	Host:     "192.168.88.38",
	Username: "root",
	Password: "dd@2019",
	Port:     22,
}

var JavaPre = &java{
	Dockerfile: "/oneci/template/docker/java/dockerfile",
	Entrypoint: "/oneci/template/docker/java/entrypoint",
}
