package link

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleHTML = `
<!doctype html>
<html>
<head>
    <title>Example Domain</title>
    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
`

func TestParse(t *testing.T) {
	expected := []string{
		"https://www.iana.org/domains/example",
	}

	reader := strings.NewReader(exampleHTML)
	actual, err := Parse(reader)

	assert.Nil(t, err)
	assert.ElementsMatch(t, actual, expected)
}

var exampleHTMLWithEmptyHref = `
<!doctype html>
<html>
<head>
    <title>Example Domain</title>
</head>
<body>
<div>
<a href="https://www.iana.org/domains/example">More information...</a>
<a>And even more information...</a>
</div>
</body>
</html>
`

func TestParseEmptyHref(t *testing.T) {
	expected := []string{
		"https://www.iana.org/domains/example",
	}

	reader := strings.NewReader(exampleHTMLWithEmptyHref)
	actual, err := Parse(reader)

	assert.Nil(t, err)
	assert.ElementsMatch(t, actual, expected)
}
