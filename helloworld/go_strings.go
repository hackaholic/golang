package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "Hello World, I am Learning Go"
	println(str)

	println(strings.Split(str, " "))

	for _, v := range strings.Split(str, " ") {
		println(v)
	}

	split_string := strings.Split(strings.ReplaceAll(str, ",", ""), " ")
	println(split_string[1], split_string[5])

	println("str: ", str)
	// strings.Compare(a string, b string) -> a == b:0, a<b: -1, a>b: +1
	println(`strings.Compare("abc", "abc"):`, strings.Compare("abc", "abc"))

	// strings.Contains(s string, substr string) -> bool
	println(`strings.Contains(str, "Wordl")`, strings.Contains(str, "World"))

	// strings.ContainsAny(s string, chars string) -> bool
	println(`strings.ContainsAny(str,"abclijk"):`, strings.ContainsAny(str, "abclijk"))

	//strings.ReplaceAll(s string, old string, new string) -> new string
	println(`strings.ReplaceAll(str, ",", ""):`, strings.ReplaceAll(str, ",", ""))

	// strings.Count(s tring, substr string) ->  int n if found n > 0, else n = 0
	println(`strings.Count(str, "l"):`, strings.Count(str, "l"))
	println(`strings.Count(str, "") -> return lenght of string when passed empty substr:`, strings.Count(str, ""))

	/* strings.EqualFold() Function in Golang reports whether s and t, interpreted as UTF-8 strings,
	are equal under Unicode case-folding, which is a more general form of case-insensitivity.
	*/
	println(`strings.EqualFold("abcdef", "aBcDEf"):`, strings.EqualFold("abcdef", "aBcDEfl"))

	//

	// strings.Fields(s string) -> split the string on whitespaces
	fmt.Println("strings.Fields(str):", strings.Fields(str), "len:", len(strings.Fields(str)))
	/* field_array := make([]string, 0)
	for _, v := range strings.Fields(str) {
		field_array = append(field_array, v)
	}
	fmt.Println(field_array)
	*/

	// strings.HasPrefix(s string, prefix string) ->  bool
	println(`strings.HasPrefix(str, "Hell"):`, strings.HasPrefix(str, "Hell"))

	// strings.HasSuffix(s string, suffix string) ->  bool
	println(`strings.HasSuffix(str, "Go"):`, strings.HasSuffix(str, "Go"))

	// strings.Index(s string, substr strig) -> int, return start position if found else -1
	// strings.LastIndex(s string substr string)  -> int
	println(`strings.Index(str, "I"):`, strings.Index(str, "I"))

	// strings.IndexAny(s string, chars string) -> int return the index of first instance of any chars else -1
	// strings.LastIndexAny(s string, chars string) -> int
	println(`strings.IndexAny(str, "QZoWGo"):`, strings.IndexAny(str, "QZoWGo"))

	// strings.Join( elements []string, sep string) -> string
	println(`strings.Join([]string{"Coco", "Dora", "Mini", "Oreo"}, ":"):`, strings.Join([]string{"Coco", "Dora", "Mini", "Oreo"}, ":"))

	myToUpper := func(r rune) rune {
		return unicode.ToUpper(r)
	}
	// strings.Map(func(rune) rune, str)
	println("strings.Map(myToUpper, str):", strings.Map(myToUpper, str))
	println("strings.Map(unicode.ToTitle, str):", strings.Map(unicode.ToTitle, str))
}
