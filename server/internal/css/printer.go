package css

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func DownloadTailwind() ([]byte, error) {
	res, err := http.Get("https://unpkg.com/tailwindcss@2.0.1/dist/tailwind.css")
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func text2GoIdentifier(p string) string {
	sb := &strings.Builder{}
	upCase := true
	written := 0
	for _, r := range p {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			upCase = true
			continue
		}

		if r >= '0' && r <= '9' && written == 0 {
			sb.WriteRune('S')
		}

		written++
		if upCase {
			sb.WriteRune(unicode.ToUpper(r))
			upCase = false
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func unescape(str string) string {
	return strings.ReplaceAll(str, "\\", "")[1:]
}

func PrintClassNamesAsGoConstants(buf []byte) error {

	uniqueClasses := map[string]string{}
NEXT_LINE:
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for i := 0; i < len(line); i++ {
			c := line[i]
			if i == 0 && c != '.' {
				continue NEXT_LINE
			}

			if c == '>' || c == '<' || c == '~' || c == '[' || c == ',' {
				continue NEXT_LINE
			}

			if c == ':' && i > 0 && line[i-1] != '\\' {
				continue NEXT_LINE
			}
		}

		if line[len(line)-1] != '{' {
			continue
		}

		line = strings.TrimSpace(line[:len(line)-1])

		uniqueClasses[text2GoIdentifier(line)] = unescape(line)

	}

	var varNames []string
	for n := range uniqueClasses {
		varNames = append(varNames, n)
	}

	sort.Strings(varNames)

	for _, n := range varNames {
		fmt.Println(n + " = " + strconv.Quote(uniqueClasses[n]))
	}

	fmt.Println("got", len(varNames))
	return nil
}
