package main

import "fmt"

type HttpProcessor[U any, T HttpMessage[U]] struct{}

func (HttpProcessor[U, T]) ProcessMessage(message HttpMessage[U]) {
	fmt.Printf("Message %v processed\n", message)
}

type WebsocketProcessor[U any, T WebsocketMessage[U]] struct{}

func (WebsocketProcessor[U, T]) ProcessMessage(message WebsocketMessage[U]) {
	fmt.Printf("Message %v processed\n", message)
}
