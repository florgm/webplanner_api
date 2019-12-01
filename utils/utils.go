package utils

import (
    "net/http"
    "errors"
    "io/ioutil"
)

func GetJSONBody(request *http.Request) ([]byte, error) {
    body := request.Body
    if body == nil {
        return nil, errors.New("nil body")
    }

    bytes, err := ioutil.ReadAll(body)
    if err != nil {
        return nil, err
    }

    return bytes, nil
}