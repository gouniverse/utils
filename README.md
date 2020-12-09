# Utils

Various utility functions

## Array Functions
- ArrayContains(array interface{}, val interface{}) (exists bool, index int)
- IsStringArrayEqual(a, b []string) bool

## Email Functions
- EmailSend(from string, to []string, subject string, htmlMessage string) (bool, error)

## Environment Variables Functions
- <b>AppAddress() string</b>
- <b>AppEnv() string</b>
- <b>AppName() string</b>
- <b>AppPort() string</b>
- <b>AppURL() string</b>
- <b>DbDriver() string</b>
- <b>DbHost() string</b>
- <b>DbPort() string</b>
- <b>DbDatabase() string</b>
- <b>DbUsername() string</b>
- <b>DbPassword() string</b>
- <b>EmailFromAddress() string</b>
- <b>EmailFromName() string</b>
- <b>Env(key string) string</b>
- <b>EnvIntialize(key string) string</b> - Intializes an .env file, if exists. Fails loudly if the file is invalid and cannot be parsed
```
utils.EnvIntialize()
```

## Fiber Framework Functions
- FiberAllIps(c *fiber.Ctx) []string
- FiberGetValueArray2D(c *fiber.Ctx, key string, defaultValue ...[]string)

## File Functions
- FileExists(filePath string) bool
- FileGetContents(filename string) (string, error)

## HTML Functions
- MinCSS(cssString string) (string, error) - Minifies a Css string
- MinHTML(htmlString string) (string, error) - Minifies a HTML string
- MinScript(sctiptString string) (string, error) - Minifies a JavaScript string
- ScriptsHTML(str string) string - returns an HTML fragment of scripts tags, with embedded and minified local scripts (mainly suitable for serverless environment)
- StylesHTML(str string) string - returns an HTML fragment of scripts tags, with embedded and minified local styles (mainly suitable for serverless environment)

## Interface Functions
- InterfaceToStringArray(v interface{}) []string

## Link Functions
- LinkWebsite() string

## Map Functions
- MapToColumn(inputMap []map[string]string, keyName string) []string
- MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string

## String Functions
- AddSlashes(str string) string
- RandStr(len int) string
- Slugify(s string, replaceWith rune) string
- TemplateParseString(template string, data) string
