package cmd

import (
	api "github.com/AndrewTamm/WillItSnow/cmd/remote/weather"
	"github.com/AndrewTamm/WillItSnow/cmd/server/weather"
	"github.com/AndrewTamm/WillItSnow/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type restWeatherService struct {
	router     *gin.Engine
	config     *config.Config
	weatherApi api.Weather
}

type WeatherServer interface {
	Serve() error
}

func NewServer(cfg *config.Config) WeatherServer {
	router := gin.Default()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	lambdaClient := lambda.New(
		sess,
		&aws.Config{
			Region: aws.String("us-west-2"),
		},
	)

	return &restWeatherService{
		config:     cfg,
		router:     router,
		weatherApi: api.NewWeatherApi(lambdaClient),
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
	rs.router.GET("/weather", weather.GetWeatherFunc(rs.config, rs.weatherApi))

	err := rs.router.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
