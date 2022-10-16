package universe

type Param struct {
}

type Action struct {
	Name    string
	Service *Service
	Params  []Param
	Handler func(*Context) interface{}
}
