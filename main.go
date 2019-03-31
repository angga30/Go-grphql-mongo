package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
	"net/http"
	//"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	b "pgomgo/type"

	//"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Nama string `json:"nama" bson:"nama"`
	Lagu string `json:"lagu" bson:"lagu"`
	Baju string `json:"baju" bson:"baju"`
}



//func a(w http.ResponseWriter,r *http.Request){
//	sess:= GetConnection()
//	var books []Book
//	err := sess.C("buku").Find(bson.M{}).All(&books)
//	if err != nil {
//
//	}
//	res,err := json.MarshalIndent(books,""," ")
//	if err != nil{
//		panic(err.Error())
//	}
//	fmt.Print(res)
//	ResponseWithJSON(w, res, http.StatusOK)
//}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), b.Schema)
		json.NewEncoder(w).Encode(result)
	})
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// Display some basic instructions

	http.ListenAndServe(":8088", nil)
}
