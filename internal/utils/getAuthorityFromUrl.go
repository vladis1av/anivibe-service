package utils

import (
	"regexp"
	"strings"
)

func GetAuthorityFromUrl(url string) string {
	r := regexp.MustCompile(`/.+?/`)
	matcherAuthority := r.FindStringSubmatch(url)
	authority := strings.Replace(matcherAuthority[0], "/", "", -1)

	return authority
}
