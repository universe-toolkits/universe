package universe

import (
	"fmt"
	"log"
)

type UniverseNode struct {
}

func Hello(name string) string {
	log.Default().Printf("Hello %v", name)
	return fmt.Sprintf("Hi, %v. Welcome to Universe!", name)
}

func GetVersion() string {
	return "1.0"
}
