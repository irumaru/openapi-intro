package main

import (
	"fmt"
	"net/http"

	oapi "github.com/irumaru/openapi-intro/generated"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
)

type apiController struct{}

// GET /user
func (a apiController) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &oapi.User{
		Id:   1,
		Name: "irumaru",
	})
}

// POST /user
func (a apiController) PostUser(ctx echo.Context) error {
	user := &oapi.User{}
	ctx.Bind(&user)
	fmt.Println(user)
	return ctx.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	// OpenAPIの仕様に沿っているかリクエストをバリデーションするミドルウェアを作成&追加
	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(echomiddleware.OapiRequestValidator(swagger))

	// LoggerとRecoverミドルウェアを追加
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := apiController{}
	oapi.RegisterHandlers(e, api)

	e.Logger.Fatal(e.Start(":1323"))
}
