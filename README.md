# Utils

Various utility functions

## Array Functions
- <b>ArrayContains(array interface{}, val interface{}) (exists bool, index int)</b>
- <b>IsStringArrayEqual(a, b []string) bool</b>

## Email Functions
- <b>EmailSend(from string, to []string, subject string, htmlMessage string) (bool, error)</b>

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
- <b>FiberAllIps(c *fiber.Ctx) []string</b>
- <b>FiberGetValueArray2D(c *fiber.Ctx, key string, defaultValue ...[]string)</b>
- <b>FiberReq(c *fiber.Ctx, key string, valueDefault string) string</b> - returns the trimmed key from POST, or default value if empty

## File Functions
- <b>FileExists(filePath string) bool</b>
- <b>FileGetContents(filename string) (string, error)</b>

## HTML Functions
- <b>MinCSS(cssString string) (string, error)</b> - Minifies a Css string
- <b>MinHTML(htmlString string) (string, error)</b> - Minifies a HTML string
- <b>MinScript(sctiptString string) (string, error)</b> - Minifies a JavaScript string
- <b>ScriptsHTML(str string) string</b> - returns an HTML fragment of scripts tags, with embedded and minified local scripts (mainly suitable for serverless environment)
- <b>StylesHTML(str string) string</b> - returns an HTML fragment of scripts tags, with embedded and minified local styles (mainly suitable for serverless environment)


## HTTP Functions
- <b>Req(r *http.Request, key string, defaultValue string) string</b> - Returns a POST or GET value for a key, or default if not exists
- <b>RespondJSON(w http.ResponseWriter, response api.Response)</b> - Respond returns an API response as JSON

## Interface Functions
- <b>InterfaceToStringArray(v interface{}) []string</b>

## Link Functions
- <b>LinkWebsite() string</b>

## Map Functions
- <b>MapToColumn(inputMap []map[string]string, keyName string) []string</b>
- <b>MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string</b>

## String Functions
- <b>AddSlashes(str string) string</b>
- <b>RandStr(len int) string</b>
- <b>Slugify(s string, replaceWith rune) string</b>
- <b>TemplateParseString(template string, data) string</b>
