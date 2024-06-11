package api

import (
    "encoding/json"
    "net/http"
)

type MonkeyParams struct {
    Modelname string
}

type MonkeyResponse struct {
    Code int // Success code
    Prediction string // Prediction string
}

type Error struct {
    Code int // Error code
    Message string // Error message
}

func writeError(w http.ResponseWriter, message string, code int) {
    resp := Error {
        Code: code,
        Message: message
    }

    w.Header().Set("Content-Type". "application/json")
    w.WriteHeader(code)

    json.NewEncoder(w).Encode(resp)
}

var (
    RequestErrorHandler = func(w http.ResponseWriter, err error) {
        writeError(w, err.Error(), http.StatusBadRequest)
    }
    InternalErrorHandler = func(w http.ResponseWriter) {
        writeError(w, "An unexpected error occured.", http.StatusInternalServerError)
    }
)
