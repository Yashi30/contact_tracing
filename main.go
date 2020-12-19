package main

import(
	"context"
	"log"
	"time"
	"net/http"
	"encoding/json"
	

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)
type Users struct{
	ID primitive.ObjectID        `json:"_id,omitempty"`
	Name string                  `json:"name,omitempty"` 
	DOB  string                  `json:"dob,omitempty"` 
	phoneNumber string           `json:"phnenum,omitempty"`
	emailaddress string          `json:"email,omitempty"`
	creationtimestamp time.Time  `json:"creation,omitempty"`    
}
type Contact struct{
	Useridone  primitive.ObjectID       `json:"_id_1,omitempty"`
	Useridtwo   primitive.ObjectID      `json:"_id_2,omitempty"`
	TimeofContact time.Time             `json:"contact,omitempty"`   
}
func PostHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
	var user Users
	json.NewDecoder(r.Body).Decode(&user)
	database:= client.Database("contact_tracing")
	UsersCollection:=database.Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_:=UsersCollection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
	
}
func PostContact(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
	var contact Contact
	json.NewDecoder(r.Body).Decode(&contact)
	database:= client.Database("contact_tracing")
	contactCollection:=database.Collection("Contact")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_:=contactCollection.InsertOne(ctx, contact)
	json.NewEncoder(w).Encode(result)
	
}
func GetParticularId(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
	params:= r.URL.Query().Get("id")
	id,_:=primitive.ObjectIDFromHex(params)
	var user Users
	database:= client.Database("contact_tracing")
	UsersCollection:=database.Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	result :=UsersCollection.FindOne(ctx,Users{ID:id}).Decode(&user)

	
	json.NewEncoder(w).Encode(result)
	

}
var client *mongo.Client
func main(){
	client ,err :=mongo.NewClient(options.Client().ApplyURI("mongodb+srv://myusername:darsh@cluster0.tk6xp.mongodb.net/contact_tracing?retryWrites=true&w=majority",))
	if err !=nil{
		log.Fatal(err)
	}
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

	http.HandleFunc("/users",PostHandler)
	http.HandleFunc("/users/{id}",GetParticularId)
	http.HandleFunc("/contacts",PostContact)
	log.Fatal(http.ListenAndServe(":8000",nil))	
    
}