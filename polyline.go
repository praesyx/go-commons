package commons

import (
	"bytes"
	"io"
)

type Polyline struct {
	Points string `json:"points"`
}

type LatLng struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// DecodePolyline converts a polyline encoded string to an array of LatLng objects.
func DecodePolyline(poly string) ([]LatLng, error) {
	p := &Polyline{
		Points: poly,
	}
	return p.Decode()
}

// Decode converts this encoded Polyline to an array of LatLng objects.
func (p *Polyline) Decode() ([]LatLng, error) {
	input := bytes.NewBufferString(p.Points)

	var lat, lng int64
	path := make([]LatLng, 0, len(p.Points)/2)
	for {
		dlat, _ := decodeInt(input)
		dlng, err := decodeInt(input)
		if err == io.EOF {
			return path, nil
		}
		if err != nil {
			return nil, err
		}

		lat, lng = lat+dlat, lng+dlng
		path = append(path, LatLng{
			Latitude:  float64(lat) * 1e-5,
			Longitude: float64(lng) * 1e-5,
		})
	}
}

// decodeInt reads an encoded int64 from the passed io.ByteReader.
func decodeInt(r io.ByteReader) (int64, error) {
	result := int64(0)
	var shift uint8

	for {
		raw, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		b := raw - 63
		result += int64(b&0x1f) << shift
		shift += 5

		if b < 0x20 {
			bit := result & 1
			result >>= 1
			if bit != 0 {
				result = ^result
			}
			return result, nil
		}
	}
}
