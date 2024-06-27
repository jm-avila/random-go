package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/jmavila/web-server-3/models"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	session *mongo.Client
}

func NewUserController(session *mongo.Client) *UserController {
	return &UserController{session}
}

func isValidObjectID(id string) (primitive.ObjectID, bool) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return objectID, false
	}
	return objectID, true
}

// isZeroType checks if the value from the struct is the zero value of its type
func isZeroType(value reflect.Value) bool {
	zero := reflect.Zero(value.Type()).Interface()

	switch value.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map:
		return value.Len() == 0
	default:
		return reflect.DeepEqual(zero, value.Interface())
	}
}

func (uc UserController) GetPing(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "{\"data\": \"Pong\"}")
}

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	object_id, isValid := isValidObjectID(id)
	if !isValid {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	collection := uc.session.Database("go_db").Collection("users")
	filter := bson.M{"_id": object_id}
	var user models.User
	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "{\"data\": \"User not found\"}")
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", userJson)
}

func (uc UserController) CreateUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	user := models.User{}
	json.NewDecoder(req.Body).Decode(&user)
	user.Id = primitive.NewObjectID()

	collection := uc.session.Database("go_db").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userJson, err := json.Marshal(insertResult)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", userJson)
}

func (uc UserController) UpdateUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	object_id, isValid := isValidObjectID(id)
	if !isValid {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "{\"data\": \"User not found\"}")
		return
	}
	user := models.User{}
	json.NewDecoder(req.Body).Decode(&user)

	typeData := reflect.TypeOf(user)
	valueData := reflect.ValueOf(user)

	updates := bson.D{}

	for i := 1; i < valueData.NumField(); i++ {
		field := typeData.Field(i)
		value := valueData.Field(i)
		tag := field.Tag.Get("json")
		if !isZeroType(value) {
			update := bson.E{Key: tag, Value: value.Interface()}
			updates = append(updates, update)
		}
	}

	collection := uc.session.Database("go_db").Collection("users")
	filter := bson.M{"_id": object_id}
	updateFilter := bson.D{{Key: "$set", Value: updates}}

	udpateResult, err := collection.UpdateOne(context.TODO(), filter, updateFilter)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userJson, err := json.Marshal(udpateResult)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", userJson)
}

func (uc UserController) DeleteUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	object_id, isValid := isValidObjectID(id)
	if !isValid {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	collection := uc.session.Database("go_db").Collection("users")
	filter := bson.M{"_id": object_id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, err := json.Marshal(deleteResult)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", userJson)
}
