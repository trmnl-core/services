package consumer

import (
	"context"

	"github.com/micro/go-micro/client"
	comments "github.com/micro/services/portfolio/comments/proto"
	followers "github.com/micro/services/portfolio/followers/proto"
	"github.com/micro/services/portfolio/helpers/microgorm"
	"github.com/micro/services/portfolio/helpers/textenhancer"
	"github.com/micro/services/portfolio/notifications/storage"
	posts "github.com/micro/services/portfolio/posts/proto"
	push "github.com/micro/services/portfolio/push-notifications/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// New returns an instance of Consumer
func New(client client.Client, db storage.Service) Consumer {
	return Consumer{
		db:           db,
		textenhancer: textenhancer.Service{},
		usersSrv:     users.NewUsersService("kytra-v1-users:8080", client),
		postsSrv:     posts.NewPostsService("kytra-v1-posts:8080", client),
		followersSrv: followers.NewFollowersService("kytra-v1-followers:8080", client),
		commentsSrv:  comments.NewCommentsService("kytra-v1-comments:8080", client),
		pushSrv:      push.NewPushNotificationsService("kytra-v1-push-notifications:8080", client),
	}
}

// Consumer is an object which processes various messages
type Consumer struct {
	db           storage.Service
	textenhancer textenhancer.Service
	usersSrv     users.UsersService
	postsSrv     posts.PostsService
	commentsSrv  comments.CommentsService
	followersSrv followers.FollowersService
	pushSrv      push.PushNotificationsService
}

func (c Consumer) usersTaggedInText(text string) []string {
	usernames := c.textenhancer.ListTaggedUsers(text)

	result := make([]string, len(usernames))
	for i, username := range usernames {
		user, err := c.usersSrv.Find(context.Background(), &users.User{Username: username})

		switch err {
		case nil:
			result[i] = user.Uuid
		case microgorm.ErrNotFound:
			continue
		default:
			return result
		}
	}

	return result
}
