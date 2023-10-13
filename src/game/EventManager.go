package game

import "fmt"

type EventManager struct {
	listeners map[string][]interface{}
}

func (em *EventManager) Subscribe(eventName string, callback interface{}) {
	em.listeners[eventName] = append(em.listeners[eventName], callback)
}

func (em *EventManager) Publish(eventName string, eventData interface{}) {
	for _, callback := range em.listeners[eventName] {
		switch fn := callback.(type) {
		case func(SCENENAME):
			fmt.Println("callback from func(SCENENAME)")
			fn(eventData.(SCENENAME))
		}
	}
}

func NewEventManager() *EventManager {
	result := EventManager{}
	result.listeners = make(map[string][]interface{})
	return &result
}
