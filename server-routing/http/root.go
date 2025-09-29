package http

import "arch-study/user"

func AppHandler() *RootHandler {
	handler := NewRootHandler()

	handler.RegisterHandler(user.UserHandler, "users")

	return handler
}

func TestHandler() *RootHandler {
	handler := NewRootHandler()

	handler.RegisterHandler(user.UserHandler, "family")

	return handler
}
