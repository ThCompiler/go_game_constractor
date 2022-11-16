package scene

import (
    "context"
    "encoding/json"
    "github.com/ThCompiler/go_game_constractor/marusia"
    "time"
)

type UserInfo struct {
    UserId    string
    SessionId string
    UserVKId  string
}

// NLU - parsing user input to language token
type NLU struct {
    Tokens   []string
    Entities []string
}

type Request struct {
    SearchedMessage string
    NameMatched     string
    FullMessage     string
    WasButton       bool
    ApplicationType marusia.ApplicationType
    Payload         json.RawMessage
    NLU             NLU
}

type Context struct {
    context.Context
    Request       Request
    Info          UserInfo
    GlobalContext context.Context
}

func NewContext(request Request, info UserInfo, ctx context.Context, GlobalContext context.Context) *Context {
    return &Context{
        Context:       ctx,
        Request:       request,
        Info:          info,
        GlobalContext: GlobalContext,
    }
}

type ContextKey string

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key ContextKey, value any) {
    c.Context = context.WithValue(c.Context, key, value)
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key ContextKey) (value any, exists bool) {
    value = c.Value(key)
    return value, value != nil
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (c *Context) MustGet(key ContextKey) any {
    if value, exists := c.Get(key); exists {
        return value
    }
    panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key ContextKey) (s string) {
    if val, ok := c.Get(key); ok && val != nil {
        s, _ = val.(string)
    }
    return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key ContextKey) (b bool) {
    if val, ok := c.Get(key); ok && val != nil {
        b, _ = val.(bool)
    }
    return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key ContextKey) (i int) {
    if val, ok := c.Get(key); ok && val != nil {
        i, _ = val.(int)
    }
    return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64(key ContextKey) (i64 int64) {
    if val, ok := c.Get(key); ok && val != nil {
        i64, _ = val.(int64)
    }
    return
}

// GetUint returns the value associated with the key as an unsigned integer.
func (c *Context) GetUint(key ContextKey) (ui uint) {
    if val, ok := c.Get(key); ok && val != nil {
        ui, _ = val.(uint)
    }
    return
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (c *Context) GetUint64(key ContextKey) (ui64 uint64) {
    if val, ok := c.Get(key); ok && val != nil {
        ui64, _ = val.(uint64)
    }
    return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key ContextKey) (f64 float64) {
    if val, ok := c.Get(key); ok && val != nil {
        f64, _ = val.(float64)
    }
    return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key ContextKey) (t time.Time) {
    if val, ok := c.Get(key); ok && val != nil {
        t, _ = val.(time.Time)
    }
    return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key ContextKey) (d time.Duration) {
    if val, ok := c.Get(key); ok && val != nil {
        d, _ = val.(time.Duration)
    }
    return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) GetStringSlice(key ContextKey) (ss []string) {
    if val, ok := c.Get(key); ok && val != nil {
        ss, _ = val.([]string)
    }
    return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key ContextKey) (sm map[string]any) {
    if val, ok := c.Get(key); ok && val != nil {
        sm, _ = val.(map[string]any)
    }
    return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key ContextKey) (sms map[string]string) {
    if val, ok := c.Get(key); ok && val != nil {
        sms, _ = val.(map[string]string)
    }
    return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key ContextKey) (smss map[string][]string) {
    if val, ok := c.Get(key); ok && val != nil {
        smss, _ = val.(map[string][]string)
    }
    return
}