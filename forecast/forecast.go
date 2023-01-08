package forecast

import (
	"bytes"
	"github.com/valyala/fastjson"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Forecaster struct {
	BingToken string
}

type Location struct {
	lat  float64
	lon  float64
	name string
}

func (f *Forecaster) GetLocation(search string) (*Location, error) {
	q := url.Values{}
	q.Add("q", search)
	q.Add("key", f.BingToken)
	resp, err := http.Get("http://dev.virtualearth.net/REST/v1/Locations?" + q.Encode())
	if err != nil {
		log.Printf("Failed to look up %s", search)
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Printf("Failed to copy from response")
		return nil, err
	}
	val, err := fastjson.ParseBytes(buf.Bytes())
	if err != nil {
		log.Printf("Failed to parse json %s", buf.String())
		return nil, err
	}
	record := val.Get("resourceSets", "0", "resources", "0")
	coords := record.GetArray("point", "coordinates")
	lat, err := coords[0].Float64()
	if err != nil {
		log.Printf("Failed to get coordinates from json %s", buf.String())
		return nil, err
	}
	lon, err := coords[1].Float64()
	if err != nil {
		log.Printf("Failed to get coordinates from json %s", buf.String())
		return nil, err
	}
	name := record.GetStringBytes("name")
	if name == nil {
		log.Printf("Failed to get name from json %s", buf.String())
		return nil, err
	}
	return &Location{
		lat:  lat,
		lon:  lon,
		name: string(name),
	}, nil
}

func GetForecast(lat string, lon string) {

}
