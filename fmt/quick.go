package fmt

import (
	"io"
	"strconv"
	"strings"
)

type stringer interface {
	String() string
}

func uitoa(v uint64) string {
	if v == 0 {
		return "0"
	}
	var (
		i = 0x13
		b [20]byte
	)
	for v >= 0xA {
		n := v / 0xA
		b[i] = byte(0x30 + v - n*0xA)
		i--
		v = n
	}
	b[i] = byte(0x30 + v)
	return string(b[i:])
}
func quickPrint(nl bool, v ...interface{}) string {
	var b strings.Builder
	quickFprint(&b, nl, v...)
	r := b.String()
	b.Reset()
	return r
}
func quickPrintf(s string, v ...interface{}) string {
	var b strings.Builder
	quickFprintf(&b, s, v...)
	r := b.String()
	b.Reset()
	return r
}
func quickFprint(b io.Writer, f bool, v ...interface{}) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}
	var (
		s    bool
		err  error
		n, c int
	)
	for i := range v {
		switch r := v[i].(type) {
		case []byte:
			if !s && i > 0 {
				n, err = b.Write([]byte{' '})
				if c += n; err != nil {
					return c, err
				}
			}
			s = true
			n, err = io.WriteString(b, string(r))
		case string:
			if !s && i > 0 {
				n, err = b.Write([]byte{' '})
				if c += n; err != nil {
					return c, err
				}
			}
			s = true
			n, err = io.WriteString(b, r)
		case stringer:
			if !s && i > 0 {
				n, err = b.Write([]byte{' '})
				if c += n; err != nil {
					return c, err
				}
			}
			s = true
			n, err = io.WriteString(b, r.String())
		default:
			if s = false; i > 0 {
				n, err = b.Write([]byte{' '})
				if c += n; err != nil {
					return c, err
				}
			}
			switch r := v[i].(type) {
			case bool:
				if r {
					n, err = io.WriteString(b, "true")
				} else {
					n, err = io.WriteString(b, "false")
				}
			case float32:
				n, err = io.WriteString(b, strconv.FormatFloat(float64(r), 'f', 2, 64))
			case float64:
				n, err = io.WriteString(b, strconv.FormatFloat(r, 'f', 2, 64))
			case int:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int8:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int16:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int32:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int64:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint8:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint16:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint32:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint64:
				n, err = io.WriteString(b, uitoa(r))
			case uintptr:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			}
		}
		if c += n; err != nil {
			return c, err
		}
	}
	if f {
		n, err = b.Write([]byte{'\n'})
		c += n
	}
	return c, err
}
func quickFprintf(b io.Writer, s string, v ...interface{}) (int, error) {
	if len(v) == 0 {
		return io.WriteString(b, s)
	}
	var (
		n, c, x, a int
		err        error
	)
	for i := 0; i < len(s); i++ {
		if a >= len(v) {
			break
		}
		if s[i] != '%' {
			continue
		}
		if i+1 >= len(s) {
			continue
		}
		n, err = io.WriteString(b, s[x:i])
		if c += n; err != nil {
			return c, err
		}
		x = i
		if i++; s[i] >= '0' && s[i] <= '9' {
			for x = i; i < len(s); i++ {
				if s[i] >= '0' && s[i] <= '9' {
					continue
				}
				break
			}
		}
		if i >= len(s) {
			break
		}
		switch s[i] {
		case 'q':
			switch r := v[a].(type) {
			case []byte:
				n, err = io.WriteString(b, strconv.Quote(string(r)))
			case string:
				n, err = io.WriteString(b, strconv.Quote(r))
			case error:
				n, err = io.WriteString(b, strconv.Quote(r.Error()))
			case stringer:
				n, err = io.WriteString(b, strconv.Quote(r.String()))
			}
		case 's', 'v':
			switch r := v[a].(type) {
			case []byte:
				n, err = io.WriteString(b, string(r))
			case string:
				n, err = io.WriteString(b, r)
			case error:
				n, err = io.WriteString(b, r.Error())
			case stringer:
				n, err = io.WriteString(b, r.String())
			case int:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int8:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int16:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int32:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case int64:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint8:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint16:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint32:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uint64:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			case uintptr:
				n, err = io.WriteString(b, uitoa(uint64(r)))
			}
		case 'f', 'e', 'E', 'g', 'G':
			var k float64
			switch r := v[a].(type) {
			case float32:
				k = float64(r)
			case float64:
				k = r
			}
			n, err = io.WriteString(b, strconv.FormatFloat(k, s[i], 2, 64))
		case 'b', 't':
			if r, ok := v[a].(bool); ok {
				if r {
					n, err = io.WriteString(b, "true")
				} else {
					n, err = io.WriteString(b, "false")
				}
			}
		case 'd', 'x', 'X', 'u':
			var k uint64
			switch r := v[a].(type) {
			case int:
				k = uint64(r)
			case int8:
				k = uint64(r)
			case int16:
				k = uint64(r)
			case int32:
				k = uint64(r)
			case int64:
				k = uint64(r)
			case uint:
				k = uint64(r)
			case uint8:
				k = uint64(r)
			case uint16:
				k = uint64(r)
			case uint32:
				k = uint64(r)
			case uint64:
				k = uint64(r)
			case uintptr:
				k = uint64(r)
			}
			if s[i] == 'x' || s[i] == 'X' {
				n, err = io.WriteString(b, strconv.FormatUint(k, 16))
			} else {
				n, err = io.WriteString(b, uitoa(k))
			}
		default:
			n, err = io.WriteString(b, s[x:i])
		}
		if c += n; err != nil {
			return c, err
		}
		x = i + 1
		a++
	}
	if n = 0; x < len(s) {
		n, err = io.WriteString(b, s[x:])
	}
	return c + n, err
}
