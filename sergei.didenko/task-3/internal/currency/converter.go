package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Convert(source []CurrencyXML) ([]Currency, error) {
	result := make([]Currency, 0, len(source))

	for _, item := range source {
		rawValue := strings.ReplaceAll(item.Value, ",", ".")

		parsedVal, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {

			return nil, fmt.Errorf("failed to parse value '%s' for %s: %w", item.Value, item.CharCode, err)
		}

		result = append(result, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    parsedVal,
		})
	}

	return result, nil
}
