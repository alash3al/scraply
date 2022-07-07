package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
)

// Input the input used for fetching the requested data
type Input struct {
	URL        string            `json:"url"`
	Method     string            `json:"method"`
	Body       []byte            `json:"body"`
	UserAgent  string            `json:"user_agent"`
	ReturnBody bool              `json:"return_body"`
	Extractors map[string]string `json:"extractors"`
}

// Output the fetch result
type Output struct {
	StatusCode int                    `json:"status_code"`
	URL        string                 `json:"url"`
	Body       string                 `json:"body,omitempty"`
	Result     map[string]interface{} `json:"result"`
}

func (o Output) String() string {
	j, _ := json.Marshal(o)

	return string(j)
}

// Do fetches the requested input simply
func Do(ctx context.Context, input Input) (*Output, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		input.Method,
		input.URL,
		bytes.NewReader(input.Body),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to initialize the request due to (%s)", err.Error())
	}

	if input.UserAgent != "" {
		req.Header.Set("User-Agent", input.UserAgent)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do the request due to (%s)", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("the remote server responded with (%s)", resp.Status)
	}

	if !strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "text/html") {
		return nil, fmt.Errorf("none `text/html` response returned (%s)", resp.Header.Get("Content-Type"))
	}

	allBodyContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read the response body due to (%s)", err.Error())
	}

	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(allBodyContent))
	if err != nil {
		return nil, fmt.Errorf("unable to load the html document due to (%s)", err.Error())
	}

	jsvm := goja.New()
	jsvm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	jsvm.SetParserOptions(parser.WithDisableSourceMaps)

	jsvm.Set("$", dom.Find)
	jsvm.Set("select", dom.Find)

	output := Output{
		StatusCode: resp.StatusCode,
		URL:        resp.Request.URL.String(),
		Result:     map[string]interface{}{},
	}

	if input.ReturnBody {
		output.Body = string(allBodyContent)
	}

	for k, script := range input.Extractors {
		val, err := jsvm.RunString(script)
		if err != nil {
			return nil, fmt.Errorf("unable to execute the %s's script (%s) due to (%s)", k, script, err.Error())
		}

		output.Result[k] = val.Export()
	}

	return &output, nil
}
