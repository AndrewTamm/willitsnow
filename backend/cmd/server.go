package cmd

import (
	"github.com/AndrewTamm/WillItSnow/cmd/server/weather"
	"github.com/AndrewTamm/WillItSnow/config"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type restWeatherService struct {
	router *gin.Engine
	config *config.Config
}

type WeatherServer interface {
	Serve() error
}

func NewServer(cfg *config.Config) WeatherServer {
	router := gin.Default()

	return &restWeatherService{
		config: cfg,
		router: router,
	}
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Printf("Error: {%+v}\n", err)
	}

	if len(c.Errors) > 0 {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (rs *restWeatherService) Serve() error {
	rs.router.Use(ErrorHandler)
	rs.router.Use(cors.Default())
	rs.router.Use(static.Serve("/", static.LocalFile("./public", true)))
	rs.router.GET("/weather", weather.GetWeatherFunc(rs.config))

	err := rs.router.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
