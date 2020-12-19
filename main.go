package main

import(
	"context"
	"fmt"
	"log"
	"time"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)
type Users struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string            `bson:"name,omitempty"` 
	Date of Birth 
}
func main(){
	client ,err :=mongo.NewClient(options.Client().ApplyURI("mongodb+srv://myusername:darsh@cluster0.tk6xp.mongodb.net/contact_tracing?retryWrites=true&w=majority",))
	if err !=nil{
		log.Fatal(err)
	}
	mongo.Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err=client.Connect(ctx)
	if err !=nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err=client.Ping(ctx, readpref.Primary())
	if err!=nil{
		log.Fatal(err)
	}
	
	db := client.Database()
}