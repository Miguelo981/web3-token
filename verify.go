package lib

import (
	"errors"
	"strings"
	"time"
)

type SignedParams struct {
	Domain string
}

func SplitSections(lines []string) [][]string {
	var sections [][]string
	sectionNumber := 0
	for i, _ := range lines {
		line := lines[i]
		sections = append(sections, []string{line})
		//sections[sectionNumber].push(line)
		if line == "" {
			sectionNumber += 1
			sections = append(sections, []string{})
		}
	}

	return sections
}

func GetDomain(sections [][]string) string {
	if strings.Contains(sections[0][0], "wants you to sign in with your Ethereum account") {
		return strings.ReplaceAll(strings.ReplaceAll(sections[0][0], " wants you to sign in with your Ethereum account.", ""), " ", "")
	}

	return ""
}

func GetStatement(sections [][]string) string {
	if len(sections) == 2 {
		domain := GetDomain(sections)

		if domain == "" {
			return sections[0][0]
		}
	}
	if len(sections) == 3 {
		return sections[1][0]
	}
	return ""
}

func ParseBody(lines []string) map[string]string {
	sections := SplitSections(lines)
	var parsedBody = map[string]string{}

	for _, v := range lines {
		keyValues := strings.Split(v, ":")
		newKey := strings.ToLower(strings.Replace(keyValues[0], " ", "-", -1))
		parsedBody[newKey] = strings.Replace(v, keyValues[0]+": ", "", 1)
	}

	domain := GetDomain(sections)
	statement := GetStatement(sections)

	if domain != "" {
		parsedBody["domain"] = domain
	}

	if statement != "" {
		parsedBody["statement"] = statement
	}

	return parsedBody
}

func Verify(token string, params SignedParams) (*DecryptedToken, error) {
	decryptedToken, err := Decrypt(token)
	if err != nil {
		return nil, err
	}

	if decryptedToken.Version == 1 {
		return nil, errors.New("Tokens version 1 are not supported by the current version of module")
	}

	lines := strings.Split(decryptedToken.StringBody, "\n")
	parsedBody := ParseBody(lines)

	decryptedToken.Body = parsedBody

	date, _ := time.Parse(time.RFC3339Nano, parsedBody["expiration-time"])
	if !time.Time.IsZero(date) && time.Now().After(date) {
		return nil, errors.New("Token expired")
	}

	date, _ = time.Parse(time.RFC3339Nano, parsedBody["not-before"])
	if !time.Time.IsZero(date) && parsedBody["not-before"] != "" && time.Now().Before(date) {
		return nil, errors.New("It's not yet time to use the token")
	}

	if params.Domain != "" && params.Domain != parsedBody["domain"] {
		return nil, errors.New("Inappropriate token domain")
	}

	return decryptedToken, nil
}
