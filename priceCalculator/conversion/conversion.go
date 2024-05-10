package conversion

import "strconv"

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatPrices, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, floatPrices)
	}
	return floats, nil
}
