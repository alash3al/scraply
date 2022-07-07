package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// Shell a REPL shell
func Shell() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		resp, err := http.Get(ctx.String("url"))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			return err
		}

		jsvm := goja.New()
		jsvm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		jsvm.SetParserOptions(parser.WithDisableSourceMaps)

		jsvm.Set("request", map[string]interface{}{
			"url": ctx.String("url"),
		})

		jsvm.Set("response", map[string]interface{}{
			"url":         resp.Request.URL.String(),
			"status_code": resp.StatusCode,
			"body":        string(body),
		})

		jsvm.Set("console", map[string]interface{}{
			"log": fmt.Println,
			"clear": func() string {
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
				return ""
			},
		})

		jsvm.RunString("clear = console.clear")
		jsvm.RunString("log = console.log")

		jsvm.Set("$", dom.Find)
		jsvm.Set("select", dom.Find)

		printPrefix := func() {
			(color.New(color.FgHiMagenta, color.Bold).PrintFunc())("➜ (scraply) > ")
		}

		jsvm.RunString("clear()")

		scanner := bufio.NewScanner(os.Stdin)

		for {
			printPrefix()

			if !scanner.Scan() {
				break
			}

			val, err := jsvm.RunScript("scraply_shell", scanner.Text())
			if err != nil {
				fmt.Println("error:", err.Error())
				continue
			}

			color.New(color.FgGreen, color.Bold).Println(val.ToString())
		}

		return nil
	}
}
