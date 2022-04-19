package event

import net "github.com/subchord/go-sse"

func New(id string, event string, data string) net.StringEvent {
	return net.StringEvent{
		Id:    id,
		Event: event,
		Data:  data,
	}
}
