package scene

import (
    "context"
    "encoding/json"
    "github.com/ThCompiler/go_game_constractor/marusia"
    "time"
)

type UserInfo struct {
    UserID    string
    SessionID string
    UserVKID  string
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

func NewContext(ctx context.Context, globalContext context.Context, request Request, info UserInfo) *Context {
    return &Context{
        Context:       ctx,
        Request:       request,
        Info:          info,
        GlobalContext: globalContext,
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
        c64, ok = val.(complex64)
        if !ok {
            c64 = 0
        }
    }

    return
}

// GetComplex128 returns the value associated with the key as a complex128.
func (c *Context) GetComplex128(key ContextKey) (c128 complex128) {
    if val, ok := c.Get(key); ok && val != nil {
        c128, ok = val.(complex128)
        if !ok {
            c128 = 0
        }
    }

    return
}

// GetByte returns the value associated with the key as a byte.
func (c *Context) GetByte(key ContextKey) (b byte) {
    if val, ok := c.Get(key); ok && val != nil {
        b, ok = val.(byte)
        if !ok {
            b = 0
        }
    }

    return
}

// GetRune returns the value associated with the key as a rune.
func (c *Context) GetRune(key ContextKey) (r rune) {
    if val, ok := c.Get(key); ok && val != nil {
        r, ok = val.(rune)
        if !ok {
            r = 0
        }
    }

    return
}

// GetBytes returns the value associated with the key as a []byte.
func (c *Context) GetBytes(key ContextKey) (bs []byte) {
    if val, ok := c.Get(key); ok && val != nil {
        bs, ok = val.([]byte)
        if !ok {
            bs = make([]byte, 0)
        }
    }

    return
}

// GetStrings returns the value associated with the key as a []string.
func (c *Context) GetStrings(key ContextKey) (ss []string) {
    if val, ok := c.Get(key); ok && val != nil {
        ss, ok = val.([]string)
        if !ok {
            ss = make([]string, 0)
        }
    }

    return
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key ContextKey) (s string) {
    if val, ok := c.Get(key); ok && val != nil {
        s, ok = val.(string)
        if !ok {
            s = ""
        }
    }

    return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key ContextKey) (b bool) {
    if val, ok := c.Get(key); ok && val != nil {
        b, ok = val.(bool)
        if !ok {
            b = false
        }
    }

    return
}

// GetInt returns the value associated with the key as an int.
func (c *Context) GetInt(key ContextKey) (i int) {
    if val, ok := c.Get(key); ok && val != nil {
        i, ok = val.(int)
        if !ok {
            i = 0
        }
    }

    return
}

// GetInt8 returns the value associated with the key as an int8.
func (c *Context) GetInt8(key ContextKey) (i8 int8) {
    if val, ok := c.Get(key); ok && val != nil {
        i8, ok = val.(int8)
        if !ok {
            i8 = 0
        }
    }

    return
}

// GetInt16 returns the value associated with the key as an int16.
func (c *Context) GetInt16(key ContextKey) (i16 int16) {
    if val, ok := c.Get(key); ok && val != nil {
        i16, ok = val.(int16)
        if !ok {
            i16 = 0
        }
    }

    return
}

// GetInt32 returns the value associated with the key as an int32.
func (c *Context) GetInt32(key ContextKey) (i32 int32) {
    if val, ok := c.Get(key); ok && val != nil {
        i32, ok = val.(int32)
        if !ok {
            i32 = 0
        }
    }

    return
}

// GetInt64 returns the value associated with the key as an int64.
func (c *Context) GetInt64(key ContextKey) (i64 int64) {
    if val, ok := c.Get(key); ok && val != nil {
        i64, ok = val.(int64)
        if !ok {
            i64 = 0
        }
    }

    return
}

// GetUint returns the value associated with the key as an uint.
func (c *Context) GetUint(key ContextKey) (ui uint) {
    if val, ok := c.Get(key); ok && val != nil {
        ui, ok = val.(uint)
        if !ok {
            ui = 0
        }
    }

    return
}

// GetUint8 returns the value associated with the key as an uint8.
func (c *Context) GetUint8(key ContextKey) (ui8 uint8) {
    if val, ok := c.Get(key); ok && val != nil {
        ui8, ok = val.(uint8)
        if !ok {
            ui8 = 0
        }
    }

    return
}

// GetUint16 returns the value associated with the key as an uint16.
func (c *Context) GetUint16(key ContextKey) (ui16 uint16) {
    if val, ok := c.Get(key); ok && val != nil {
        ui16, ok = val.(uint16)
        if !ok {
            ui16 = 0
        }
    }

    return
}

// GetUint32 returns the value associated with the key as an uint32.
func (c *Context) GetUint32(key ContextKey) (ui32 uint32) {
    if val, ok := c.Get(key); ok && val != nil {
        ui32, ok = val.(uint32)
        if !ok {
            ui32 = 0
        }
    }

    return
}

// GetUint64 returns the value associated with the key as an uint64.
func (c *Context) GetUint64(key ContextKey) (ui64 uint64) {
    if val, ok := c.Get(key); ok && val != nil {
        ui64, ok = val.(uint64)
        if !ok {
            ui64 = 0
        }
    }

    return
}

// GetFloat32 returns the value associated with the key as a float32.
func (c *Context) GetFloat32(key ContextKey) (f32 float32) {
    if val, ok := c.Get(key); ok && val != nil {
        f32, ok = val.(float32)
        if !ok {
            f32 = 0
        }
    }

    return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key ContextKey) (f64 float64) {
    if val, ok := c.Get(key); ok && val != nil {
        f64, ok = val.(float64)
        if !ok {
            f64 = 0
        }
    }

    return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key ContextKey) (t time.Time) {
    if val, ok := c.Get(key); ok && val != nil {
        t, ok = val.(time.Time)
        if !ok {
            t = time.Time{}
        }
    }

    return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key ContextKey) (d time.Duration) {
    if val, ok := c.Get(key); ok && val != nil {
        d, ok = val.(time.Duration)
        if !ok {
            d = 0
        }
    }

    return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key ContextKey) (sm map[string]any) {
    if val, ok := c.Get(key); ok && val != nil {
        sm, ok = val.(map[string]any)
        if !ok {
            sm = make(map[string]any)
        }
    }

    return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key ContextKey) (sms map[string]string) {
    if val, ok := c.Get(key); ok && val != nil {
        sms, ok = val.(map[string]string)
        if !ok {
            sms = make(map[string]string)
        }
    }

    return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key ContextKey) (smss map[string][]string) {
    if val, ok := c.Get(key); ok && val != nil {
        smss, ok = val.(map[string][]string)
        if !ok {
            smss = make(map[string][]string)
        }
    }

    return
}

func GetContextAny[T any](ctx *Context, key ContextKey) (r T) {
    if val, ok := ctx.Get(key); ok && val != nil {
        r, ok = val.(T)
        if !ok {
            var zeroV T
            r = zeroV
        }
    }

    return
}
