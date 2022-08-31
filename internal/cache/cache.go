package cache

import net "github.com/subchord/go-sse"

var store map[string]net.Event

func init() {
	store = make(map[string]net.Event)
}

func Put(e net.Event) {
	store[e.GetId()] = e
}

func Pull() map[string]net.Event {
	return store
}
