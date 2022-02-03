package models

type Configuration struct {
	Port    string   `yaml:"port"`
	Domains []string `yaml:"domains"`
}
