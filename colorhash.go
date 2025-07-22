package colorhash

import (
	"crypto/sha256"
	"fmt"
	"image/color"
	"strconv"
)

// Given a byte slice of any length, returns a deterministically
// picked color to represent it
func GetColor(b []byte, alpha bool) color.Color {
	sha := sha256.New()
	sha.Write(b)
	hash := sha.Sum(nil)
	var A uint8
	if alpha {
		A = uint8(hash[3])
	}
	return color.RGBA{
		R: uint8(hash[0]),
		G: uint8(hash[1]),
		B: uint8(hash[2]),
		A: A,
	}
}

// Returns a color like GetColor but in the form of an integer
func GetInt(by []byte) (i int64, e error) {
	r, g, b, _ := GetColor(by, false).RGBA()
	hex := fmt.Sprintf("%02X%02X%02X", r&0xFF, g&0xFF, b&0xFF)
	i, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return i, fmt.Errorf("unable to convert \"%s\", alpha %t, %w", by, false, err)
	}
	return i, nil
}

// Returns a color like GetColor but in the form of an integer
func GetIntA(by []byte) (i int64, e error) {
	r, g, b, a := GetColor(by, true).RGBA()
	hex := fmt.Sprintf("%02X%02X%02X%02X", r&0xFF, g&0xFF, b&0xFF, a&0xFF)
	i, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return i, fmt.Errorf("unable to convert \"%s\", alpha %t, %w", by, true, err)
	}
	return i, nil
}

// Given a string, returns a deterministically picked color to represent it
// Alias for RGBFromString
func RGB(s string) (rgb string) {
	return RGBFromString(s)
}

// Given a string, returns a deterministically picked color to represent it
// Alias for RGBFromString
func RGBA(s string) (rgb string) {
	return RGBAFromString(s)
}

// Given a string, returns a deterministically picked color to represent it
func RGBFromString(s string) (rgb string) {
	r, g, b, _ := GetColor([]byte(s), false).RGBA()
	return fmt.Sprintf("%02X%02X%02X", r&0xFF, g&0xFF, b&0xFF)
}

// Given a string, returns a deterministically picked color to represent it
func RGBAFromString(s string) (rgb string) {
	r, g, b, a := GetColor([]byte(s), true).RGBA()
	return fmt.Sprintf("%02X%02X%02X%02X", r&0xFF, g&0xFF, b&0xFF, a&0xFF)
}

// Given a string, returns a deterministically picked color to represent it
// Alias for HTMLFromString
func HTML(s string) (rgb string) {
	return "#" + RGBFromString(s)
}

// Given a string, returns a deterministically picked color to represent it
// Alias for HTMLFromString
func HTMLA(s string) (rgb string) {
	return "#" + RGBAFromString(s)
}

// Given a string, returns a deterministically picked color to represent it
// With a leading pound sign (#)
func HTMLFromString(s string) (rgb string) {
	return fmt.Sprintf("#%s", RGBFromString(s))
}

// Given a string, returns a deterministically picked color to represent it
// With a leading pound sign (#)
func HTMLAFromString(s string) (rgb string) {
	return fmt.Sprintf("#%s", RGBAFromString(s))
}

// Given any type, returns a deterministically picked color to represent it
func RGBFromAny(s any) (rgb string) {
	return fmt.Sprintf("#%s", RGB(fmt.Sprintf("%+v", s)))
}

// Given any type, returns a deterministically picked color to represent it
func RGBAFromAny(s any) (rgb string) {
	return fmt.Sprintf("#%s", RGBA(fmt.Sprintf("%+v", s)))
}
