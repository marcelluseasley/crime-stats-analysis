package crimedata  



type Crimes []struct {
	Description string    `json:"description"`
	Datetime    string    `json:"datetime"`
	Location    []float64 `json:"location"`
}