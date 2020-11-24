# Utils

Various utility functions

## Array Functions
- InArray(val interface{}, array interface{}) (exists bool, index int)

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

## String Functions
- TemplateParseString(template string, data) string
