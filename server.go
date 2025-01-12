package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SampleRadius struct {
	Name string
	Radius float64
}

func main() {

	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	coll := client.Database("test").Collection("sample_radius")

	http.HandleFunc("GET /{$}",handlerRoot)
	http.HandleFunc("GET /pi",handlerPI)
	http.HandleFunc("GET /circle/{radius}",handlerCircleArea)

	http.HandleFunc("GET /radius/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the name %s\n", name)
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w,"%f\n", result["radius"])
		})
	
	http.HandleFunc("POST /radius/", func(w http.ResponseWriter, r *http.Request) {
		var radiusStruct SampleRadius
		err := json.NewDecoder(r.Body).Decode(&radiusStruct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
		radius := radiusStruct.Radius
		name := radiusStruct.Name
		fmt.Fprintf(w, "name = %s\n", name)
		fmt.Fprintf(w, "radius = %f\n", radius)

		_, err = coll.InsertOne(context.TODO(), bson.D{{"name", name}, {"radius", radius}})
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Document inserted successfully")
		})

	log.Fatal(http.ListenAndServe(":8000",nil))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handlerPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "3.14159265359")
}

func handlerCircleArea(w http.ResponseWriter, r *http.Request) {
	radiusStr := r.PathValue("radius")
	radius,err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid radius")
		return
	}
	fmt.Fprintf(w, "Area of the circle with radius %f is %f", radius, 3.14159265359 * radius * radius)
}