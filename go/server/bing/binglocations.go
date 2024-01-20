package bing 


import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

func GetLocations(street, city, state, zipcode string) (*Response, error) {

    address := fmt.Sprintf("%s,%s,%s,%s", street, city, state, zipcode)
    q := url.QueryEscape(address)

    resp, err := http.Get(fmt.Sprintf("http://dev.virtualearth.net/REST/v1/Locations?q=%s&key=xx", q))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    return &response, nil
}