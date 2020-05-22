package config

type ConsulConfig struct {
	Address string
	Port int
}

type AppConfig struct {
	Name string
	Port int
	WSPort int
	NFS string
	NFSPath string
	Tag []string
}

type java struct {
	Dockerfile string
	Entrypoint string
}

var Conf = &ConsulConfig{
	Address: "192.168.66.100",
	Port:    8500,
}

var JavaPre = &java{
	Dockerfile: "/oneci/template/docker/java",
	Entrypoint: "/oneci/template/docker/java/entrypoint",
}
