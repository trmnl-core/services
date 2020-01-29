package handler

import (
	"github.com/micro/go-micro/client"
	bullbear "github.com/micro/services/portfolio/bullbear/proto"
	comments "github.com/micro/services/portfolio/comments/proto"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	photos "github.com/micro/services/portfolio/helpers/photos"
	"github.com/micro/services/portfolio/helpers/textenhancer"
	post "github.com/micro/services/portfolio/posts/proto"
	posts "github.com/micro/services/portfolio/posts/proto"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// Handler is an object can process RPC requests
type Handler struct {
	textenhancer textenhancer.Service
	photos       photos.Service
	auth         auth.Authenticator
	users        users.UsersService
	stocks       stocks.StocksService
	posts        post.PostsService
	bullBear     bullbear.BullBearService
	comments     comments.CommentsService
}

// New returns an instance of Handler
func New(auth auth.Authenticator, pics photos.Service, client client.Client) Handler {
	return Handler{
		auth:         auth,
		photos:       pics,
		textenhancer: textenhancer.Service{},
		users:        users.NewUsersService("kytra-v1-users:8080", client),
		stocks:       stocks.NewStocksService("kytra-v1-stocks:8080", client),
		posts:        posts.NewPostsService("kytra-v1-posts:8080", client),
		bullBear:     bullbear.NewBullBearService("kytra-v1-bullbear:8080", client),
		comments:     comments.NewCommentsService("kytra-v1-comments:8080", client),
	}
}
