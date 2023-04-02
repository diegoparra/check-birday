package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/diegoparra/check-birthday/pkg/database"
	"github.com/diegoparra/check-birthday/pkg/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUser() {

	db, ctx, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	// coll := db.Database("new-reikianos").Collection("check-birthday")

	// carla := User{
	// 	ID:        primitive.NewObjectID(),
	// 	Name:      "Rosalina",
	// 	Surname:   "Parra",
	// 	Birthday:  "01/04",
	// 	CreatedAt: time.Now(),
	// }

	// createUser(ctx, &carla, coll)
}

func GetUsers(c *gin.Context) {

	db, ctx, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	coll := db.Database("new-reikianos").Collection("check-birthday")

	today := time.Now().Format("02/01")

	users := repositories.GetUser(ctx, coll, today)

	if users == nil {
		fmt.Println("not found any user")
		c.JSON(http.StatusOK, gin.H{"data": "Ninguem fazendo niver hoje"})
	} else {

		c.JSON(http.StatusOK, gin.H{"data": users})
	}

}
