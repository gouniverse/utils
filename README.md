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
- EmailFromAddress() string
- EmailFromName() string
- Env(key string) string

## Fiber Framework
- FiberGetValueArray2D(c *fiber.Ctx, key string, defaultValue ...[]string)

## String Functions
- TemplateParseString(template string, data) string
