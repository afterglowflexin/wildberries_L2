// Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса,
// при этом давая шанс обработать этот запрос нескольким объектам.
package main

type Handler interface {
	SendRequest(message int) string
}

type HandlerA struct {
	next Handler
}

func (h *HandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Handler 1 processed the request"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerB struct {
	next Handler
}

func (h *HandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Handler 2 processed the request"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerC struct {
	next Handler
}

func (h *HandlerC) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Handler 3 processed the request"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
