package config

type Properties struct {
	server ServerProperties
	db     DBProperties
}

type ServerProperties struct {
	port int
}

type DBProperties struct {
	username string
	password string
	name     string
}
