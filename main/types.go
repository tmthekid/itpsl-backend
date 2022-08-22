package main

type Response struct {
	Message string `json:"message"`
}

func response(message string) Response {
	return Response{Message: message}
}