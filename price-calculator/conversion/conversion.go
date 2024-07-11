package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {

	floats := make([]float64, len(strings))

	for i, str := range strings {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = floatVal
	}
	return floats, nil
}
