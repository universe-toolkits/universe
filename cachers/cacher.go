package cachers

type Cacher interface {
	Connect() (bool, error)
	Disconnect() (bool, error)
	Set(key string, value interface{}) *Cacher
	Get(key string, object *interface{})
	Heartbeat() (bool, error)
}
