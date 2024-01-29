package esclient

import "strconv"

// Place contains the schema for a restaurant
type Place struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phone   string   `json:"phone"`
	Geo     GeoPoint `json:"location"`
}

// GeoPoint is a geo point struct
type GeoPoint struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// NewPlace creates a new restaurant
func (p *Place) NewPlace(line []string) {
	for j, field := range line {
		switch j {
		case 0:
			p.ID, _ = strconv.Atoi(field)
		case 1:
			p.Name = field
		case 2:
			p.Address = field
		case 3:
			p.Phone = field
		case 4:
			p.Geo.Longitude, _ = strconv.ParseFloat(field, 10)
		case 5:
			p.Geo.Latitude, _ = strconv.ParseFloat(field, 10)
		}
	}
}

// CreatePlacesList creates a list of restaurants
func CreatePlacesList(data [][]string) []Place {
	var PlacesList []Place
	for i, line := range data {
		if i > 0 {
			p := Place{}
			p.NewPlace(line)
			PlacesList = append(PlacesList, p)
		}
	}
	return PlacesList
}
