package weather

import (
	api "github.com/AndrewTamm/WillItSnow/cmd/remote/weather"
	"github.com/AndrewTamm/WillItSnow/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWeatherFunc(cfg *config.Config, weatherApi api.Weather) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		weather, err := weatherApi.CallWeatherApi(cfg.Weather.Location)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.IndentedJSON(http.StatusOK, weather)
	}
}
