package main

import (
	"graphql-go/db"
	"graphql-go/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

func main() {
	DB := db.Init()
	dbHandler := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", dbHandler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", dbHandler.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", dbHandler.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", dbHandler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", dbHandler.DeleteBook).Methods(http.MethodDelete)

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &BeastSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", graphqlHandler)

	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHTML)
	}))

	log.Println("API running on port 8080")

	http.ListenAndServe(":8080", router)
}

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: "http://localhost:8080/graphql",
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body>
</html>`)
