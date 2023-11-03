package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

// MinScript minifies the script
func MinScript(scriptString string) (string, error) {
	m := minify.New()
	// m.AddFunc("text/css", css.Minify)
	// m.AddFunc("text/html", html.Minify)
	// m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	return m.String("text/javascript", scriptString)
}

// MinCSS minifies the CSS
func MinCSS(cssString string) (string, error) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	return m.String("text/css", cssString)
}

// MinHTML minifies the CSS
func MinHTML(htmlString string) (string, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	return m.String("text/html", htmlString)
}

// ScriptsHTML the HTML from scripts string
func ScriptsHTML(str string) string {
	scripts := strings.Split(str, ",")

	scriptsHTML := ""

	for _, script := range scripts {

		if strings.HasPrefix(script, "http://") || strings.HasPrefix(script, "https://") {
			scriptsHTML += "<script src=\"" + script + "\"></script>"
			continue
		}

		path := "views/" + script

		if FileExists(path) == false {
			println(path + " does not exists")
			continue
		}

		scriptText, getContentsError := FileGetContents(path)

		if getContentsError != nil {
			log.Println(path + " get contents error")
			continue
		}

		scriptText, err := MinScript(scriptText)
		if err != nil {
			log.Println(err)
		}
		scriptsHTML += "<script>" + scriptText + "</script>"
	}

	return scriptsHTML
}

// StylesHTML the styles
func StylesHTML(str string) string {
	styles := strings.Split(str, ",")

	htmlStyles := ""

	for _, style := range styles {

		if strings.HasPrefix(style, "http://") || strings.HasPrefix(style, "https://") {
			htmlStyles += "<link href=\"" + style + "\"  rel=\"stylesheet\" />"
			continue
		}

		path := "views/" + style

		if FileExists(path) == false {
			println(path + " does not exists")
			continue
		}

		scriptText, getContentsError := FileGetContents(path)

		if getContentsError != nil {
			log.Println(path + " get contents error")
			continue
		}

		styleText, err := MinScript(scriptText)
		if err != nil {
			log.Println(err)
		}

		htmlStyles += "<style>" + styleText + "</style>"
	}

	return htmlStyles
}
