package config

type Yaml struct {
	Security security
}

type security struct {
	method []string
}
