package main

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/supperdoggy/spotify-web-project/spotify-db/internal/db"
	handlers2 "github.com/supperdoggy/spotify-web-project/spotify-db/internal/handlers"
	service2 "github.com/supperdoggy/spotify-web-project/spotify-db/internal/service"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, _ := zap.NewDevelopment()
	db, err := db2.NewDB("spotify", logger)
	if err != nil {
		logger.Fatal("error connecting to db", zap.Error(err))
	}
	service := service2.NewService(db, logger)
	handlers := handlers2.NewHandlers(service, logger)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/addSegment", handlers.AddSegments)
		apiv1.GET("/allsongs", handlers.GetAllSongs)
		apiv1.POST("/getsegment", handlers.GetSegment)
		apiv1.POST("/new_user", handlers.NewUser)
		apiv1.POST("/get_user", handlers.GetUser)

		// playlists
		apiv1.POST("/new_playlist", handlers.NewPlaylist)
		apiv1.POST("/delete_playlist", handlers.DeletePlaylist)
		apiv1.POST("/user_playlists", handlers.GetUserPlaylists)
		apiv1.POST("/get_playlist", handlers.GetUserPlaylist)
		apiv1.POST("/add_song_playlist", handlers.AddSongToUserPlaylist)
		apiv1.POST("/remove_song_playlist", handlers.RemoveSongFromUserPlaylist)
	}

	if err := r.Run(":8082"); err != nil {
		logger.Fatal("error running db service")
	}
}
