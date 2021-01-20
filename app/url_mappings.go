package app

import (
	"github.com/halimakibb/bookstore_users-api/controllers/ping"
	"github.com/halimakibb/bookstore_users-api/controllers/users"
)

// MapUrls Url Routing
func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
