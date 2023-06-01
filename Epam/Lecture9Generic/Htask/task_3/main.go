package main

import "fmt"

// Task: create Processors for HttpMessage and for WebsocketMessage
// Processors must have method ProcessMessage(msg T) with call of fmt.Println("Message processed")

type HttpMessage[T any] struct {
	content T
}

type WebsocketMessage[T any] struct {
	content T
}

type HttpProcessor[K comparable, V any] struct {
	ProcessorMap map[K]V
}

type WebsocketProcessor[K comparable, V any] struct {
	ProcessorMap map[K]V
}

func (p *HttpProcessor[K, V]) ProcessMessage(msg V) {
	fmt.Println("Http message processed")
}

func (p *WebsocketProcessor[K, V]) ProcessMessage(msg V) {
	fmt.Println("Websocket message processed")
}

func main() {
	httpProcessor := &HttpProcessor[string, HttpMessage[string]]{}
	msg := HttpMessage[string]{
		content: "qwerty",
	}
	httpProcessor.ProcessMessage(msg)

	websocketProcessor := &WebsocketProcessor[[]byte, WebsocketMessage[[]byte]]{}
	msg2 := WebsocketMessage[[]byte]{
		content: []byte("asdfgh"),
	}
	websocketProcessor.ProcessMessage(msg2)
}
