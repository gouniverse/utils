# Utils

Various utility functions

## Array Functions
- InArray(val interface{}, array interface{}) (exists bool, index int)

## Email Functions
- EmailSend(from string, to []string, subject string, htmlMessage string) (bool, error)

## Environment Variables Functions
- AppAddress() string
- AppEnv() string
- AppName() string
- AppPort() string
- AppURL() string
- DbDriver() string
- DbHost() string
- DbPort() string
- DbDatabase() string
- DbUsername() string
- DbPassword() string
- EmailFromAddress() string
- EmailFromName() string
- Env(key string) string

## Fiber Framework Functions
- FiberGetValueArray2D(c *fiber.Ctx, key string, defaultValue ...[]string)

## File Functions
- FileExists(filePath string) bool
- FileGetContents(filename string) (string, error)

## Link Functions
- LinkWebsite() string

## Map Functions
- MapToColumn(inputMap []map[string]string, keyName string) []string
- MapToKeyValue(inputMap []map[string]string, keyName string, valueName string) map[string]string

## String Functions
- TemplateParseString(template string, data) string
