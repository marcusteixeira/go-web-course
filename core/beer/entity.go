package beer

// Usar em maiusculo caso seja acessado por outros pacotes, minusculo somente dentro do pacote.
// tags json definem como os dados vão ser transformados em JSON
type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle   = 1
	TypeLager = 2
	TypeMalt  = 3
	TypeStout = 4
)

func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknown"
}

type BeerStyle int

//usando desta forma o compilador vai automaticamente definir os ids sequencialmente
const (
	StyleAmber = iota + 1 // sempre adicione no final, para não ocasionar problemas com ID e coisas do tipo.
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (t BeerStyle) String() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "India Pale Ale"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown"
}
