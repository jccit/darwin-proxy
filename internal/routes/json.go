package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var upperRegex = regexp.MustCompile(`^[^a-z]*$`)

type conventionalMarshaller struct {
	Value interface{}
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			// Empty keys are valid JSON, only lowercase if we do not have an
			// empty key.
			if len(match) > 2 {
				// check if the key is all upper case
				snip := match[1:]
				snip = snip[:len(snip)-1]
				if upperRegex.Match(snip) {
					out := strings.ToLower(string(match))
					return []byte(out)
				}

				// Decode first rune after the double quotes
				r, width := utf8.DecodeRune(match[1:])
				r = unicode.ToLower(r)
				utf8.EncodeRune(match[1:width+1], r)
			}
			return match
		},
	)

	return converted, err
}

func returnJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	encoded, _ := json.MarshalIndent(conventionalMarshaller{response}, "", " ")
	fmt.Fprintln(w, string(encoded[:]))
}
