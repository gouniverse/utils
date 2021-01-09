# Utils

![tests](https://github.com/gouniverse/utils/workflows/tests/badge.svg)

Various utility functions

## Array Functions
- <b>ArrayContains(array interface{}, val interface{}) (exists bool, index int)</b>
- <b>ArrayMerge(array ...[]interface{}) []interface{}</b>
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
- <b>FilePutContents(filename string, data string, mode os.FileMode) error</b> - writes a string to file
- <b>FileToBase64(filePath string) string</b> - converts a file to Base64 encoded string
- <b>ImgToBase64Url(filePath string) string</b> - converts an image file to Base64 encoded URL string

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
- <b>AddSlashes(str string) string</b> - adds slashes
- <b>RandStr(len int) string</b> - generates a random string
- <b>Slugify(s string, replaceWith rune) string</b> - converts a string to slug
- StrToInt(s string) (int, error) - converts a string to Int32
- StrToInt64(s string) (int64, error) -  converts a string to Int64
- <b>TemplateParseString(template string, data) string</b> - parses a template file and returns as string
- <b>ToString(v interface{}) string</b> - converts an interface to string

# Time Functions
- <b>StrToTimeUnix(str string) (int64, error)</b> - converts string to Unix time

# Other Functions
- <b>IsNumeric(s string) bool</b> - checks if a string is numeric
