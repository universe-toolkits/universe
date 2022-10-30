package universe

type UniverseNode struct {
	Id      string
	Actions []string
	Events  []string
}

func GetVersion() string {
	return "1.0"
}
