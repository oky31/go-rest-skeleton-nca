package app

import (
	"net/http"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status": "available",
		"system_information": map[string]string{
			"environment": app.config.Env,
			"version":     Version,
		},
	}

	if err := app.writeJSON(w, http.StatusOK, data, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
