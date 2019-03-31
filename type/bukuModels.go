package _type

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
	mog "gopkg.in/mgo.v2"

)

type Book struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Nama string `json:"nama" bson:"nama"`
	Lagu string `json:"lagu" bson:"lagu"`
	Baju string `json:"baju" bson:"baju"`
}

type kuct struct {
	Data Book
}

var books []Book


func GetConnection() *mog.Database{
	session,err := mog.Dial("mongodb://angga:prabu@localhost/")
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Print("ok")
	}

	return session.DB("nambor")
}


var buku = graphql.NewObject(graphql.ObjectConfig{
	Name: "Buku",
	Fields: graphql.Fields{
		"Id": &graphql.Field{
			Type: graphql.String,
		},
		"Nama": &graphql.Field{
			Type: graphql.String,
		},
		"Lagu": &graphql.Field{
			Type: graphql.String,
		},
		"Baju": &graphql.Field{
			Type: graphql.String,
		},
	},
})


// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
		*/
		"createTodo": &graphql.Field{
			Type:        buku, // the return type for this field
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"nama": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				text, _ := params.Args["nama"].(string)

				// figure out new id


				// perform mutation operation here
				// for e.g. create a Todo and save to DB.
				newTodo := Book{
					Nama: text,
				}

				books = append(books, newTodo)

				// return the new Todo object that we supposedly save to DB
				// Note here that
				// - we are returning a `Todo` struct instance here
				// - we previously specified the return Type to be `todoType`
				// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
				return newTodo, nil
			},
		},
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
		*/
		"updateTodo": &graphql.Field{
			Type:        buku, // the return type for this field
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// marshall and cast the argument value
				//done, _ := params.Args["done"].(bool)
				//id, _ := params.Args["id"].(string)
				affectedTodo := Book{}


				// Return affected todo
				return affectedTodo, nil
			},
		},
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		/*
		   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
		*/
		"todo": &graphql.Field{
			Type:        buku,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"nama": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				sess:= GetConnection()
				idQuery, _ := params.Args["nama"].(string)
				fmt.Print(idQuery)
				boos := Book{}
				var book []Book
				ses := sess.C("buku").Find(bson.M{}).Iter()
				for ses.Next(&boos){
					book = append(book,boos)
				}
				uj,err := json.Marshal(book)
				if err != err {
					panic(err.Error())
				}

				fmt.Print(book)

				return uj, nil
			},
		},

	},
})

// define schema, with our rootQuery and rootMutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})