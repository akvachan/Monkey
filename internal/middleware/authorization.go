package middleware

import (
    "errors"
    "net/http"

    "github.com/akvachan/Monkey/api"
    "github.com/akvachan/Monkey/internal/tools"
    log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid modelname or id.")

func Authorization(new http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        var modelName string = r.URL.Query().Get("modelName")
        var modelId = r.Header.Get("Authorization")
        var err error

        if modelName == "" || token == "" {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        var databse *tools.DatabaseInterface
        database, err = tools.NewDatabase()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }

        var loginDetails *tools.LoginDetails
        loginDetails = (*database).GetUserLoginDetails(modelName)

        if (loginDetails == nil || (modelId != (*loginDetails).ModelId)) {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        next.ServeHTTP(w, t)

    })
}
