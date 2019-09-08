package conf

type RemoteType int

const (
	Local RemoteType = 1 + iota
	Remote
)

var serverType = [...]string{
	"Local",
	"Remote",
}

type HubInfo struct {
	Name string		`json:"name"`
	Address string 	`json:"address"`
	Type string 	`json:"type"`
}

type MessageQueue struct {
	Endpoint string
	Port int
	User string
	Pass string
}

type Configuration struct {
	HubInfos []HubInfo 		`json:"hubs"`
	MsgQueue MessageQueue 	`json:"msg_queue"`
}