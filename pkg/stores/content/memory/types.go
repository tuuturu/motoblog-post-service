package memory

import "errors"

type store struct {
	db map[string][]byte
}
