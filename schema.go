package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/graphql-go/graphql"
)

// Helper function to import json from file to map
func importJsonDataFromFile(fileName string, result interface{}) (isOk bool) {
	isOk = true

	content, error := ioutil.ReadFile(fileName)

	if error != nil {
		fmt.Print("Error: ", error)
		isOk = false
	}

	error = json.Unmarshal(content, result)

	if error != nil {
		isOk = false
		fmt.Print("Error: ", error)
	}

	return
}

type Beast struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OtherNames  []string `json:"otherNames"`
	ImageUrl    string   `json:"imageUrl"`
}

var BeastList []Beast

var _ = importJsonDataFromFile("./beastData.json", &BeastList)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"beast": &graphql.Field{
			Type:        beastType,
			Description: "Get single beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				nameQuery, isOk := params.Args["name"].(string)

				if isOk {
					// Search for beast by name
					for _, beast := range BeastList {
						if beast.Name == nameQuery {
							return beast, nil
						}
					}
				}

				return Beast{}, nil
			},
		},

		"beastList": &graphql.Field{
			Type:        graphql.NewList(beastType),
			Description: "List of beasts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return BeastList, nil
			},
		},
	},
})

var beastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Beast",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.Int,
		},
		"otherNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BeastSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
