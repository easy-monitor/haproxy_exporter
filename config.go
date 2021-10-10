package main

type Config struct {
	Modules []*HaProxyModule `yaml:"module" json:"module"`
}

type HaProxyModule struct {
	Name string `yaml:"name" json:"name"`
	User string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}
