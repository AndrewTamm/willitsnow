package weather

import (
	"errors"
	"fmt"
	"github.com/AndrewTamm/WillItSnow/cmd/model"
	"github.com/AndrewTamm/WillItSnow/cmd/remote/weather/response"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/goburrow/cache"
	"github.com/goccy/go-json"
	"strconv"
	"time"
)

type weatherApi struct {
	weatherLoadingCache cache.LoadingCache
	lambdaClient        *lambda.Lambda
}

type Weather interface {
	CallWeatherApi(location string) (*model.Weather, error)
}

func NewWeatherApi(lambdaClient *lambda.Lambda) Weather {
	w := &weatherApi{
		lambdaClient: lambdaClient,
	}

	w.weatherLoadingCache = cache.NewLoadingCache(w.loaderFunc,
		cache.WithRefreshAfterWrite(30*time.Minute),
		cache.WithExpireAfterWrite(60*time.Minute))

	return w
}

func (w *weatherApi) CallWeatherApi(location string) (*model.Weather, error) {
	cacheResponse, err := w.weatherLoadingCache.Get(location)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather information: %w", err)
	}

	apiResponse, ok := cacheResponse.(*model.Weather)
	if !ok {
		return nil, errors.New("unexpected type returned from api loading")
	}

	return apiResponse, nil
}

func (w *weatherApi) loaderFunc(location cache.Key) (cache.Value, error) {
	locationString, ok := location.(string)
	if !ok {
		return nil, errors.New("key expected to be string")
	}

	request := fetchWeatherRequest{locationString}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("error marshalling FetchWeatherLambda request")
	}

	result, err := w.lambdaClient.Invoke(&lambda.InvokeInput{FunctionName: aws.String("FetchWeatherLambda"), Payload: payload})
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
}

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
