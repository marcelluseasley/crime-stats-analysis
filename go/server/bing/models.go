package bing



type Response struct {
	ResourceSets             []struct {
		EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			Name  string    `json:"name"`
			Point struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`			
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
}