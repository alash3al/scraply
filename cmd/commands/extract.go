package commands

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alash3al/scraply/pkg/fetch"
	"github.com/urfave/cli/v2"
)

// Extractor a simple factory that returns the extractor command implementation
func Extractor() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		extractors := map[string]string{}

		for _, v := range ctx.StringSlice("extract") {
			parts := strings.SplitN(v, "=", 2)
			if len(parts) < 2 {
				return fmt.Errorf("invalid extractor (%s), it should be in the form of k=v", v)
			}
			extractors[parts[0]] = parts[1]
		}

		output, err := fetch.Do(
			ctx.Context,
			fetch.Input{
				URL:        ctx.String("url"),
				Method:     http.MethodGet,
				UserAgent:  ctx.String("user-agent"),
				ReturnBody: ctx.Bool("return-body"),
				Extractors: extractors,
			},
		)

		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	}
}
