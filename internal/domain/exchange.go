package domain

import "fmt"

type Event struct {
	Market string `json:"market"`
	Event  []byte `json:"event"`
}

func (e Event) String() string {
	return fmt.Sprintf("%s: %s", e.Market, e.Event)
}
