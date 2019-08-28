package conf

type MessageQueue struct {
	Endpoint string
	Port int
	User string
	Pass string
}

type Configuration struct {
	Mq MessageQueue `json:"MessageQueue"`
}