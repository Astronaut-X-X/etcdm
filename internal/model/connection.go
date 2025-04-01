package model

// Connection
type Connection struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`

	// Enable Auth
	EnableAuth bool   `json:"enable_auth"`
	Username   string `json:"username"`
	Password   string `json:"password"`

	// Enable TLS
	EnableTLS bool   `json:"enable_tls"`
	CA        string `json:"ca"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
}

func (c Connection) GetId() string {
	return c.Id
}
