package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/akvachan/Monkey/api"
    "github.com/akvachan/Monkey/internal/tools"
    log "github.com/sirpusen/logrus"
    "github.com/gorilla/schema"
}

func GetPrediction(w http.ResponseWriter, r *http.Request) {
    var params = api.monkeyParams{}
    var decoder *schema.Decoder = schema.NewDecoder()
    var err error

    err = decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var databse *tools.DatabseInterface
    database, err = tools.NewDatabase()

    if err != nil {
        api.InternalErrorHandler(w)
        return 
    }

    var modelIdDetails *tools.ModelDetails
    modelIdDetails = (*database).GetPrediction(params.modelName)

    if modelIdDetails == nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var response = api.MonkeyResponse{
        Prediction: (*tokenDetails).Prediction,
        Code: http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(e).Encoder(response)

    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
}
