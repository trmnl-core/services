package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/micro/go-micro/broker"
	notifiableEvent "github.com/micro/services/portfolio/notifications/helpers/notifiableevent"
	"github.com/micro/services/portfolio/notifications/storage"
	users "github.com/micro/services/portfolio/users/proto"
)

// Follow is the JSON object published by the followers
type Follow struct {
	Follower struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	} `json:"follower"`
	Followee struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	} `json:"followee"`
}

// ConsumeNewFollow creates notifications for any users who are followed
func (c Consumer) ConsumeNewFollow(e broker.Event) error {
	fmt.Printf("[ConsumeNewFollow] Processing Message: %v\n", string(e.Message().Body))

	// Create a new notifiable event to keep track of the notifications
	event := notifiableEvent.New(c.db, c.pushSrv)

	// Decode the message
	var follow Follow
	if err := json.Unmarshal(e.Message().Body, &follow); err != nil {
		return err
	}
	if follow.Followee.Type != "User" {
		fmt.Println("The resource followed was not a user")
		return nil
	}

	// Fetch the follower
	follower, err := c.usersSrv.Find(context.Background(), &users.User{Uuid: follow.Follower.UUID})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Notify the followed user
	event.SendNotification(storage.Notification{
		UserUUID:     follow.Followee.UUID,
		Title:        fmt.Sprintf("@%v followed you", follower.Username),
		Emoji:        "👋",
		ResourceType: "Investor",
		ResourceUUID: follower.Uuid,
	})

	return nil
}
