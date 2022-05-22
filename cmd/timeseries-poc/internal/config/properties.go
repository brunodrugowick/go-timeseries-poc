package config

type Properties struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

type Server struct {
	Port int `json:"port"`
}

type Database struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
