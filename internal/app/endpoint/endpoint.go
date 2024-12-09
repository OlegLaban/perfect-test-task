package endpoint

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	duration := fmt.Sprintf("Days left: %d", e.s.DaysLeft())

	err := ctx.String(http.StatusOK, duration)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
