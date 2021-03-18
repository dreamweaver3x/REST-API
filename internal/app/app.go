package app

import (
	"avito/internal/dateMarshaller"
	"avito/internal/models"
	"avito/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Application struct {
	repo *repository.StatsRepository
}

func NewApplication(repo *repository.StatsRepository) *Application {
	return &Application{repo: repo}
}

func (a *Application) Start(port string) {
	e := echo.New()
	e.POST("/save", a.AddNewStats)
	e.GET("/stats", a.GetStats)
	//e.DELETE()
	e.Logger.Fatal(e.Start(port))
}

func (a *Application) GetStats(c echo.Context) error {
	req := &struct {
		From  dateMarshaller.CustomDate `json:"from"`
		To    dateMarshaller.CustomDate `json:"to"`
		Order string                    `json:"order"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	statsSlice, err := a.repo.GetStats(req.From, req.To, req.Order)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, statsSlice)
}

func (a *Application) AddNewStats(c echo.Context) error {
	stats := &models.Stats{}
	if err := c.Bind(stats); err != nil {
		return err
	}

	if err := a.repo.Create(stats); err != nil {
		return err
	}

	return c.String(http.StatusAccepted, "Stats added")
}
