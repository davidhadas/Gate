package queuegate

import "fmt"

const version = "0.0.3"

func Talk() string {
	return fmt.Sprintf("queuegate version: %s", version)
}
