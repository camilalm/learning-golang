package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

type myHandler struct{}

func (handler *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(myMiddleware)
	handler := myHandler{}
	r.Handle("/handler", &handler)

	// query param
	r.Get("/v0", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("name")
		id := r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + " " + id))
		} else {
			w.Write([]byte("Hello!!"))
		}
	})

	// route param
	r.Get("/v0/{productName}/{id}", func(w http.ResponseWriter, r *http.Request) {
		product := chi.URLParam(r, "productName")
		id := chi.URLParam(r, "id")
		w.Write([]byte(product + " " + id))
	})

	// json response
	r.Get("/v0/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	// method POST
	r.Post("/v0/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}
