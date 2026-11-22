package currency

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type CurrencyXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyRatesXML struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

func LoadXML(path string) (*CurrencyRatesXML, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("open XML file: %w", err)
	}

	defer f.Close()

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = charset.NewReaderLabel

	var rates CurrencyRatesXML

	if err := decoder.Decode(&rates); err != nil {

		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return &rates, nil
}
