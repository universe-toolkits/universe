package universe

type Param struct {
	Name string
	Type *interface{}
}

type Action struct {
	Name    string
	Service *Service
	Params  []Param
	Handler func(*Context) interface{}
}
