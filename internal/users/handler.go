package users

import (
	"github.com/RaniAgus/go-templ/internal/util"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return util.Render(c, Show(User{
		FirstName: "Agustin",
		LastName:  "Ranieri",
	}))
}
