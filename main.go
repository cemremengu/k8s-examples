package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	prefix := os.Getenv("MY_API_PREFIX")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET(prefix, handler)
	e.GET(prefix+"/ping/:service", pingHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, createResponse(c.Request().URL.Path))
}

func pingHandler(c echo.Context) error {
	svc := c.Param("service")
	res, err := pingService(svc)
	if err != nil {
		slog.Error("error pinging service", "service", svc, "error", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	slog.Info("client: got response", "response", res)
	return c.String(http.StatusOK, res)
}

func pingService(svc string) (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	slog.Info("client: making http request", "name", svc)
	res, err := client.Get(svc)
	if err != nil {
		return "", fmt.Errorf("error making http request: %s\n", err)
	}

	bytedata, err := io.ReadAll(res.Body)
	reqBodyString := string(bytedata)

	return fmt.Sprintf("client: got response! status code: %d\nbody: %s\n", res.StatusCode, reqBodyString), nil
}

func createResponse(path string) string {
	podName := os.Getenv("MY_POD_NAME")
	myEnvVar := os.Getenv("MY_ENV_VAR")

	return fmt.Sprintf("Got hit for: %s\nPod Name: %s\nEnv Var: %s\n", path, podName, myEnvVar)
}
