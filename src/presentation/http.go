package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/ui"
)

func HttpServerInit(persistentDbSvc *db.PersistentDatabaseService) {
	e := echo.New()

	api.ApiInit(e, persistentDbSvc)
	ui.UiInit(e, persistentDbSvc)

	e.Start(":8080")
}
