package crimedata 

import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

func GetCrimeLocations(startdate, enddate string, lat, long float64) (*Crimes, error) {

    address := fmt.Sprintf("startdate=%s&enddate=%s&lat=%f&long=%f", startdate, enddate, lat, long)
    q := url.QueryEscape(address)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://jgentes-crime-data-v1.p.rapidapi.com/crime?%q", q), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", "d3bf40cb61msh822f908de5fb96ap1ad337jsn41a69d29e0d9")
	req.Header.Add("X-RapidAPI-Host", "jgentes-Crime-Data-v1.p.rapidapi.com")

    res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

    var response Crimes
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    return &response, nil
}