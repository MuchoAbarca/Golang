package main

import(
	"bytes"
	"fmt"
	"regexp"
)

var pr = fmt.Println

func main() {
	match, _ := regexp.MatchString("pr([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("pr([a-z]+)ch")

	pr(r.MatchString("peach"))
	pr(r.FindString("peach punch"))
	pr(r.FindStringIndex("peach punch"))
	pr(r.FindStringSubmatch("peach punch"))
	pr(r.FindStringSubmatchIndex("peach punch"))
	pr(r.FindAllString("peach punch pinch", -1))
	pr(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	pr(r.FindAllString("peach punch pinch", 2))
	pr(r.Match([]byte("peach")))

	r = regexp.MustCompile("pr([a-z]+)ch")
	pr(r)

	pr(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pr(string(out))
}
