package i18n

import "github.com/HarmoDeveloper/harmo-i18n/langs"

type Parser struct {
	countries         map[string]string
	invertedCountries map[string]string
}

func (p Parser) Parse(id string) string {
	if name, ok := p.countries[id]; ok {
		return name
	}

	if code, ok := p.invertedCountries[id]; ok {
		return code
	}

	return ""
}

func NewParser() *Parser {
	invertedCountries := make(map[string]string)

	for code, name := range langs.Countries {
		invertedCountries[name] = code
	}

	return &Parser{
		countries:         langs.Countries,
		invertedCountries: invertedCountries,
	}
}

func Parse(id string) string {
	parser := NewParser()
	return parser.Parse(id)
}
