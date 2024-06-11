package lib

import "fmt"

type Color string

const (
	RED    Color = "Merah"
	BLUE   Color = "Biru"
	YELLOW Color = "Kuning"
	BLACK  Color = "Hitam"
	WHITE  Color = "Putih"
	SILVER Color = "Silver"
)

func (c Color) String() string {
	return string(c)
}

func (c Color) Parse(s string) (Color, error) {
	switch s {
	case "Merah":
		return RED, nil
	case "Biru":
		return BLUE, nil
	case "Kuning":
		return YELLOW, nil
	case "Hitam":
		return BLACK, nil
	case "Putih":
		return WHITE, nil
	case "Silver":
		return SILVER, nil
	default:
		return "", fmt.Errorf("invalid color: %s", s)
	}
}

type CarType string

const (
	SUV CarType = "SUV"
	MPV CarType = "MPV"
)

func (c CarType) String() string {
	return string(c)
}

func (c CarType) Parse(s string) (CarType, error) {
	switch s {
	case "SUV":
		return SUV, nil
	case "MPV":
		return MPV, nil
	default:
		return "", fmt.Errorf("invalid car type: %s", s)
	}
}
