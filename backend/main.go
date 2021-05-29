package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type stat struct {
	Domain    string   `json:"domain"`
	Count     int      `json:"count"`
	LastSeen  string   `json:"lastseen"`
	FirstSeen string   `json:"firstseen"`
	FeedData  []bson.M `json:"feed"`
}

type statKafkaReader struct {
	Messages int    `json:"Messages"`
	Lag      int    `json:"Lag"`
	Topic    string `json:"Topic"`
}

type statKafkaWriter struct {
	Messages int    `json:"Messages"`
	Topic    string `json:"Topic"`
}

type worker struct {
	Uri           string `json:"uri"`
	Name          string `json:"name"`
	InTopic       string `json:"intopic"`
	Lag           int    `json:"lag"`
	OutTopic      string `json:"outtopic"`
	ReadMessages  int    `json:"readmessages"`
	WriteMessages int    `json:"writemessages"`
	Id            string `json:"_id"`
	Tag           string `json:"tag"`
	Filter        string `json:"filter"`
}

type workerdoc struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Url     string             `bson:"url"`
	Port    string             `bson:"port"`
	Name    string             `bson:"name"`
	Comment string             `bson:"comment"`
	Tag     string             `bson:"tag"`
	Filter  string             `bson:"filter"`
}

const MongoURI string = "mongodb://10.5.200.153:27017/"

func GetMongoDbConnection(uri string) (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil

}

func getMongoDbCollection(uri string, DbName string, CollectionName string) (*mongo.Collection, error) {

	client, err := GetMongoDbConnection(uri)

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func getWorkerStat(uri string, stat interface{}) (err error) {

	resp, err := http.Get(uri)
	resp.Close = true
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(body, stat)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func GetAllWorkers(c *gin.Context) {

	wksCollection, err := getMongoDbCollection(MongoURI, "feed", "workers")

	//var res []string

	filterCursor, err := wksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var res []workerdoc

	if err = filterCursor.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}
	var wks []worker

	for _, v := range res {

		uri := v.Url + ":" + v.Port
		reader := &statKafkaReader{}
		writer := &statKafkaWriter{}
		err := getWorkerStat(uri+"/metrics/reader", reader)
		if err != nil {
			log.Println(err)
		}

		err = getWorkerStat(uri+"/metrics/writer", writer)
		if err != nil {
			log.Println(err)
		}

		var w worker

		w.InTopic = reader.Topic
		w.Lag = reader.Lag
		w.Name = v.Name
		w.OutTopic = writer.Topic
		w.ReadMessages = reader.Messages
		w.WriteMessages = writer.Messages
		w.Uri = uri
		w.Id = v.ID.String()
		w.Tag = v.Tag
		w.Filter = v.Filter
		wks = append(wks, w)

	}

	c.JSON(200, wks)

}

type mongodoc struct {
	db         string
	collection string
	filter     string
}

func GetAllIOCLAlarms(uri string, db mongodoc) []stat {

	resp, err := http.Get(uri)
	resp.Close = true
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	f := []stat{}

	jsonErr := json.Unmarshal(body, &f)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	ipCollection, err := getMongoDbCollection(MongoURI, db.db, db.collection)

	//var res []string

	for i := 0; i < len(f); i++ {

		filterCursor, err := ipCollection.Find(context.Background(), bson.M{db.filter: f[i].Domain})
		if err != nil {
			log.Fatal(err)
		}

		var res []bson.M

		if err = filterCursor.All(context.Background(), &res); err != nil {
			log.Fatal(err)
		}

		f[i].FeedData = append(f[i].FeedData, res...)

	}

	return f

}

func GetAllIOCS(c *gin.Context) {

	tag := c.Param("tag")
	wksCollection, err := getMongoDbCollection(MongoURI, "feed", "workers")

	//var res []string

	filterCursor, err := wksCollection.Find(context.Background(), bson.M{"tag": tag})
	if err != nil {
		log.Fatal(err)
	}

	var res []workerdoc

	if err = filterCursor.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}

	var resp []stat

	for _, v := range res {
		doc := mongodoc{"feed", tag, v.Filter}
		uri := v.Url + ":" + v.Port + "/metrics/ioc"
		res := GetAllIOCLAlarms(uri, doc)
		resp = append(resp, res...)

	}

	c.JSON(200, resp)

}

func GetIOCByID(c *gin.Context) {

	id := c.Param("id")
	feed := c.Param("feed")

	wksCollection, err := getMongoDbCollection(MongoURI, "feed", feed)

	//var res []string

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	filterCursor, err := wksCollection.Find(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
	}

	var res []bson.M

	if err = filterCursor.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, res)

}

func GetIOCByName(c *gin.Context) {

	name := c.Param("name")
	feed := c.Param("feed")

	wksCollection, err := getMongoDbCollection(MongoURI, "feed", feed)

	//var res []string

	var filter string

	switch feed {
	case "ip":
		filter = "ip.v4"
	case "dns":
		filter = "domain"
	case "url":
		filter = "url"
	default:
		log.Println("Wrong feed name")
		c.JSON(200, "")

	}

	filterCursor, err := wksCollection.Find(context.Background(), bson.M{filter: name})
	if err != nil {
		log.Fatal(err)
	}

	var res []bson.M

	if err = filterCursor.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, res)

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./wwwroot", true)))

	api := r.Group("/api")
	{
		api.GET("/workers", GetAllWorkers)
		api.GET("/iocs/:tag", GetAllIOCS)
		api.GET("/ioc/searchbyid/:feed/:id", GetIOCByID)
		api.GET("/ioc/searchbyname/:feed/:name", GetIOCByName)
	}

	r.Run(":3000")

}
