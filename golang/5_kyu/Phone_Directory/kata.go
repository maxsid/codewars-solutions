package kata

import (
	"regexp"
	"strings"
)

func Phone(dir, num string) string {
	r := regexp.MustCompile(`\+` + num)
	findings := r.FindAllStringIndex(dir, 2)
	if len(findings) == 0 {
		return "Error => Not found: " + num
	} else if len(findings) > 1 {
		return "Error => Too many people: " + num
	}
	f := findings[0]
	startLine := strings.LastIndexByte(dir[0:f[0]], '\n')
	if startLine == -1 {
		startLine = 0
	}
	endLine := f[1] + strings.IndexByte(dir[f[1]:], '\n')
	if endLine < 0 || endLine >= len(dir) {
		endLine = len(dir)
	}
	line := dir[startLine+1 : endLine]
	line = r.ReplaceAllString(line, "")
	r = regexp.MustCompile(`\<(.+)\>`)
	name := r.FindStringSubmatch(line)[1]
	line = r.ReplaceAllString(line, "")
	rS := regexp.MustCompile(`\s{2,}|_`)
	r = regexp.MustCompile(`^\s+|\s+$|[^a-zA-Z0-9\s.\-]|,`)
	for i := 0; i < 2; i++ {
		line = rS.ReplaceAllString(line, " ")
		line = r.ReplaceAllString(line, "")
	}
	return "Phone => " + num + ", Name => " + name + ", Address => " + line
}
