package helper

import "github.com/julienschmidt/httprouter"

type RouterWithPrefix struct {
	Router *httprouter.Router
	Prefix string
}

func (r *RouterWithPrefix) GET(path string, handle httprouter.Handle){
	r.Router.GET(r.Prefix + path, handle)
}
func (r *RouterWithPrefix) POST(path string, handle httprouter.Handle){
	r.Router.POST(r.Prefix + path, handle)
}
func (r *RouterWithPrefix) PUT(path string, handle httprouter.Handle){
	r.Router.PUT(r.Prefix + path, handle)
}
func (r *RouterWithPrefix) PATCH(path string, handle httprouter.Handle){
	r.Router.PATCH(r.Prefix + path, handle)
}
func (r *RouterWithPrefix) DELETE(path string, handle httprouter.Handle){
	r.Router.DELETE(r.Prefix + path, handle)
}