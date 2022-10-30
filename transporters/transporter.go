package transporters

type TransportEvents int

const (
	OnNodeConnected TransportEvents = iota
	OnActionCall
	OnEventCall
	OnStart
	OnStop
	OnReconnecting
	OnError
)

type ListenerCallbackParams struct {
	NodeId   string
	Services map[string]string
	Events   map[string]string

	RequestId  string
	Caller     string
	SericeName string
	ActionName string
	EventName  string
	Params     map[string]interface{}
	Meta       map[string]interface{}

	Message string
	Reason  string

	Error error
}

type ListenerCallback func(p *ListenerCallbackParams)

type Transporter interface {
	Connect() (*Transporter, error)
	Disconnect() (*Transporter, error)
	Heartbeat() (bool, error)
	Call(node string, action string, params map[string]interface{}) (map[string]interface{}, error)
	AddListener(e TransportEvents, f ListenerCallback)
}
