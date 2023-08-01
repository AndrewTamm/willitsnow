package weather

import (
	"errors"
	"fmt"
	"github.com/AndrewTamm/WillItSnow/cmd/model"
	"github.com/AndrewTamm/WillItSnow/cmd/remote/weather/response"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/goburrow/cache"
	"github.com/goccy/go-json"
	"strconv"
	"time"
)

var weatherLoadingCache cache.LoadingCache

type fetchWeatherRequest struct {
	Location string `json:"location"`
}

type fetchWeatherResponseError struct {
	Message string `json:"message"`
}

type fetchWeatherHeaders struct {
	ContentType string `json:"Content-Type"`
}

type fetchWeatherResponse struct {
	StatusCode int                         `json:"statusCode"`
	Headers    fetchWeatherHeaders         `json:"headers"`
	Body       response.WeatherApiResponse `json:"body"`
}

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	weatherLoadingCache = cache.NewLoadingCache(func(location cache.Key) (cache.Value, error) {
		locationString, ok := location.(string)
		if !ok {
			return nil, errors.New("key expected to be string")
		}

		request := fetchWeatherRequest{locationString}

		payload, err := json.Marshal(request)
		if err != nil {
			return nil, errors.New("error marshalling FetchWeatherLambda request")
		}

		result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("FetchWeatherLambda"), Payload: payload})
		if err != nil {
			return nil, errors.New("error calling FetchWeatherLambda")
		}

		var resp fetchWeatherResponse

		err = json.Unmarshal(result.Payload, &resp)
		if err != nil {
			return nil, errors.New("error unmarshalling FetchWeatherLambda response")
		}

		// If the status code is NOT 200, the call failed
		if resp.StatusCode != 200 {
			return nil, errors.New("error getting items, StatusCode: " + strconv.Itoa(resp.StatusCode))
		}

		return resp.Body.AsModel(), nil
	}, cache.WithRefreshAfterWrite(time.Duration(60)*time.Minute))
}

func CallWeatherApi(location string) (*model.Weather, error) {
	cacheResponse, err := weatherLoadingCache.Get(location)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather information: %w", err)
	}

	apiResponse, ok := cacheResponse.(*model.Weather)
	if !ok {
		return nil, errors.New("unexpected type returned from api loading")
	}

	return apiResponse, nil
}
