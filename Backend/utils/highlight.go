package utils

import "github.com/sourcegraph/syntaxhighlight"

func ToGoSyntaxHighlight(code []byte) (string, error) {
	newCode, err := syntaxhighlight.AsHTML(code)
	if err != nil {
		return "", err
	}
	return `<pre><code class="language-go">` + string(newCode) + `</code></pre>`, nil
}
