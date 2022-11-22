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

// GetComplex64 returns the value associated with the key as a complex64.
func (c *Context) GetComplex64(key ContextKey) (c64 complex64) {
    if val, ok := c.Get(key); ok && val != nil {
        c64, _ = val.(complex64)
    }
    return
}

// GetComplex128 returns the value associated with the key as a complex128.
func (c *Context) GetComplex128(key ContextKey) (c128 complex128) {
    if val, ok := c.Get(key); ok && val != nil {
        c128, _ = val.(complex128)
    }
    return
}

// GetByte returns the value associated with the key as a byte.
func (c *Context) GetByte(key ContextKey) (b byte) {
    if val, ok := c.Get(key); ok && val != nil {
        b, _ = val.(byte)
    }
    return
}

// GetRune returns the value associated with the key as a rune.
func (c *Context) GetRune(key ContextKey) (r rune) {
    if val, ok := c.Get(key); ok && val != nil {
        r, _ = val.(rune)
    }
    return
}

// GetBytes returns the value associated with the key as a []byte.
func (c *Context) GetBytes(key ContextKey) (bs []byte) {
    if val, ok := c.Get(key); ok && val != nil {
        bs, _ = val.([]byte)
    }
    return
}

// GetStrings returns the value associated with the key as a []string.
func (c *Context) GetStrings(key ContextKey) (ss []string) {
    if val, ok := c.Get(key); ok && val != nil {
        ss, _ = val.([]string)
    }
    return
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

// GetInt returns the value associated with the key as an int.
func (c *Context) GetInt(key ContextKey) (i int) {
    if val, ok := c.Get(key); ok && val != nil {
        i, _ = val.(int)
    }
    return
}

// GetInt8 returns the value associated with the key as an int8.
func (c *Context) GetInt8(key ContextKey) (i8 int8) {
    if val, ok := c.Get(key); ok && val != nil {
        i8, _ = val.(int8)
    }
    return
}

// GetInt16 returns the value associated with the key as an int16.
func (c *Context) GetInt16(key ContextKey) (i16 int16) {
    if val, ok := c.Get(key); ok && val != nil {
        i16, _ = val.(int16)
    }
    return
}

// GetInt32 returns the value associated with the key as an int32.
func (c *Context) GetInt32(key ContextKey) (i32 int32) {
    if val, ok := c.Get(key); ok && val != nil {
        i32, _ = val.(int32)
    }
    return
}

// GetInt64 returns the value associated with the key as an int64.
func (c *Context) GetInt64(key ContextKey) (i64 int64) {
    if val, ok := c.Get(key); ok && val != nil {
        i64, _ = val.(int64)
    }
    return
}

// GetUint returns the value associated with the key as an uint.
func (c *Context) GetUint(key ContextKey) (ui uint) {
    if val, ok := c.Get(key); ok && val != nil {
        ui, _ = val.(uint)
    }
    return
}

// GetUint8 returns the value associated with the key as an uint8.
func (c *Context) GetUint8(key ContextKey) (ui8 uint8) {
    if val, ok := c.Get(key); ok && val != nil {
        ui8, _ = val.(uint8)
    }
    return
}

// GetUint16 returns the value associated with the key as an uint16.
func (c *Context) GetUint16(key ContextKey) (ui16 uint16) {
    if val, ok := c.Get(key); ok && val != nil {
        ui16, _ = val.(uint16)
    }
    return
}

// GetUint32 returns the value associated with the key as an uint32.
func (c *Context) GetUint32(key ContextKey) (ui32 uint32) {
    if val, ok := c.Get(key); ok && val != nil {
        ui32, _ = val.(uint32)
    }
    return
}

// GetUint64 returns the value associated with the key as an uint64.
func (c *Context) GetUint64(key ContextKey) (ui64 uint64) {
    if val, ok := c.Get(key); ok && val != nil {
        ui64, _ = val.(uint64)
    }
    return
}

// GetFloat32 returns the value associated with the key as a float32.
func (c *Context) GetFloat32(key ContextKey) (f32 float32) {
    if val, ok := c.Get(key); ok && val != nil {
        f32, _ = val.(float32)
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

func GetContextAny[T any](ctx *Context, key ContextKey) (r T) {
    if val, ok := ctx.Get(key); ok && val != nil {
        r, _ = val.(T)
    }
    return
}
