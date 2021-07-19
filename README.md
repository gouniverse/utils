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
- <b>AppAddress() string</b> - returns the URL addree the app is running under (APP_URL:APP:PORT)
- <b>AppEnv() string</b> - returns the environment the app is running in  (APP_ENV)
- <b>AppInDevelopment() bool</b> - return whether app is in development
- <b>AppInLive() bool</b> - return whether app is in production / live
- <b>AppInProduction() bool</b> - return whether app is in production / live
- <b>AppInTesting() bool</b> - return whether app is being tested
- <b>AppName() string</b> - returns the name the app is running under (APP_NAME)
- <b>AppPort() string</b> - returns the port the app is running under (APP_PORT)
- <b>AppURL() string</b> - returns the URL the app is running under (APP_URL)
- <b>DbDriver() string</b> - returns the database driver (DB_DRIVER)
- <b>DbHost() string</b> - returns the database host (DB_HOST)
- <b>DbPort() string</b> - returns the database port (DB_PORT)
- <b>DbDatabase() string</b> - returns the database name driver (DB_DATABASE)
- <b>DbUsername() string</b> - returns the database username (DB_USERNAME)
- <b>DbPassword() string</b> - returns the database password (DB_PASSWORD)
- <b>EmailFromAddress() string</b> - returns the mail from address (MAIL_FROM)
- <b>EmailFromName() string</b> - returns the mail from name (MAIL_NAME)
- <b>Env(key string) string</b> - returns an enviroment vaialble (i.e. Env("DB_DRIVER"))
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
- <b>IP(r *http.Request) string</b> - Returns the IP address of the user
- <b>Req(r *http.Request, key string, defaultValue string) string</b> - Returns a POST or GET value for a key, or default if not exists
- <b>RespondJSON(w http.ResponseWriter, response api.Response)</b> - Respond returns an API response as JSON

## Interface Functions
- <b>InterfaceToStringArray(v interface{}) []string</b>

## Link Functions
- <b>LinkWebsite() string</b>

## Map Functions
- <b>MapToColumn(inputMap []map[string]string, keyName string) []string</b> -  returns a column from map
- <b>MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string</b> -  returns a key-value array from an array of maps

## String Functions
- <b>AddSlashes(str string) string</b> - adds slashes
- <b>Base64Decode(src string) ([]byte, error)</b> - decodes a string from Base64
- <b>Base64Encode(src []byte) string</b> - encodes a string to Base64
- <b>RandStr(len int) string</b> - generates a random string
- <b>RandStrFromGamma(length int, gamma string) string</b> - generates random string of specified length with the characters specified in the gamma string
- <b>Slugify(s string, replaceWith rune) string</b> - converts a string to slug
- <b>StrToInt(s string) (int, error)</b> - converts a string to Int32
- <b>StrToInt64(s string) (int64, error)</b> -  converts a string to Int64
- <b>StrToMD5Hash(text string) string</b> - StrToMD5Hash converts a string to MD5 hash
- <b>StrToSHA1Hash(text string) string</b> - StrToSHA1Hash converts a string to SHA1 hash
- <b>StrToSHA256Hash(text string) string</b> - converts a string to SHA256 hash
- <b>TemplateParseString(template string, data) string</b> - parses a template file and returns as string
- <b>ToString(v interface{}) string</b> - converts an interface to string

# Time Functions
- <b>StrToTimeUnix(str string) (int64, error)</b> - converts string to Unix time

# Other Functions
- <b>IsNumeric(s string) bool</b> - checks if a string is numeric

# Change Log
2021-07-19 - Added functions AppInDevelopment, AppInLive, AppInProduction, AppInTesting
