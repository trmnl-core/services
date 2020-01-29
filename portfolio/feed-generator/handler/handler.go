package handler

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	feeditems "github.com/micro/services/portfolio/feed-items/proto"
	followers "github.com/micro/services/portfolio/followers/proto"
	posts "github.com/micro/services/portfolio/posts/proto"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// New returns an instance of Handler
func New(client client.Client) Handler {
	return Handler{
		usersSrv:     users.NewUsersService("kytra-v1-users:8080", client),
		postsSrv:     posts.NewPostsService("kytra-v1-posts:8080", client),
		stocksSrv:    stocks.NewStocksService("kytra-v1-stocks:8080", client),
		followerSrv:  followers.NewFollowersService("kytra-v1-followers:8080", client),
		feedItemsSrv: feeditems.NewFeedItemsService("kytra-v1-feed-items:8080", client),
	}
}

// Handler is an object which processes various messages
type Handler struct {
	usersSrv     users.UsersService
	postsSrv     posts.PostsService
	stocksSrv    stocks.StocksService
	followerSrv  followers.FollowersService
	feedItemsSrv feeditems.FeedItemsService
}

func (h Handler) addPostToUserFeed(userUUID, postUUID, description string) error {
	item := feeditems.FeedItem{
		FeedType:    "User",
		FeedUuid:    userUUID,
		Tag:         "POST",
		PostUuid:    postUUID,
		Description: description,
	}

	rsp, err := h.feedItemsSrv.Create(context.Background(), &item)
	if err != nil {
		fmt.Printf("Error creating post: %v \n", err)
		return err
	}

	fmt.Printf("Created post! UUID: %v\n", rsp.Item.Uuid)
	return nil
}
