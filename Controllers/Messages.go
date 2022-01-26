package Controllers

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"time"

	"example.com/seen-tech-rtx/DBManager"
	"example.com/seen-tech-rtx/Models"
	"example.com/seen-tech-rtx/Utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MessagesNew(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Messages
	var self Models.Messages
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}
	self.DateTime = primitive.NewDateTimeFromTime(time.Now())
	_, err = collection.InsertOne(context.Background(), self)
	if err != nil {
		c.Status(500)
		return err
	}
	c.Status(200).Send([]byte("Dak sent successfully"))
	return nil
}

func MessagesGetAll(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Messages

	// Fill the received search obj data
	var self Models.MessagesSearch
	c.BodyParser(&self)

	var results []bson.M
	b, results := Utils.FindByFilter(collection, self.GetBranchSearchBSONObj())
	if !b {
		err := errors.New("server error")
		c.Status(500).Send([]byte(err.Error()))
		return err
	}
	byteArr, _ := json.Marshal(results)
	var ResultDocs []Models.Messages
	json.Unmarshal(byteArr, &ResultDocs)
	// Sort by date
	sort.SliceStable(ResultDocs, func(i, j int) bool {
		return ResultDocs[i].DateTime < ResultDocs[j].DateTime
	})

	// Decode
	response, _ := json.Marshal(
		bson.M{"results": ResultDocs},
	)
	c.Set("Content-Type", "application/json")
	c.Status(200).Send(response)

	return nil
}
