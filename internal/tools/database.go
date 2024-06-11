package tools

import (
    log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
    ModelId string
    ModelName string
}

type ModelDetails struct {
    ModelName string
    Prediction string
}

type DatabaseInterface interface {
    GetUserLoginDetails(modelName string) *LoginDetails
    GetUserPrediction(modelName string) *ModelDetails
    SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
    var database DatabaseInterface = &mockDB{}
    var err error = database.SetupDatabase()
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &database, nil
}
