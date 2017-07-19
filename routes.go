package main

import "net/http"

// Route defines a URL Path and a Func to which handle the Path
type Route struct {
	Path string
	Func http.HandlerFunc
}

var routes = []Route{
	{
		Path: "/",
		Func: indexHandler,
	},
	{
		Path: "/view/",
		Func: makeHandler(viewHandler),
	},
	{
		Path: "/edit/",
		Func: makeHandler(editHandler),
	},
	{
		Path: "/save/",
		Func: makeHandler(saveHandler),
	},
}
