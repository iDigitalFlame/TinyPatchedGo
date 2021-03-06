package fmt

import (
	"io"
	"strconv"
	"strings"
)

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
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case int8:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case int16:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case int32:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case int64:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case uint:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case uint8:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case uint16:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case uint32:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
			case uint64:
				n, err = io.WriteString(b, strconv.FormatUint(r, 10))
			case uintptr:
				n, err = io.WriteString(b, strconv.FormatUint(uint64(r), 10))
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
		return 0, nil
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
		case 's':
			switch r := v[a].(type) {
			case []byte:
				n, err = io.WriteString(b, string(r))
			case string:
				n, err = io.WriteString(b, r)
			}
		case 'f':
			var k float64
			switch r := v[a].(type) {
			case float32:
				k = float64(r)
			case float64:
				k = r
			}
			n, err = io.WriteString(b, strconv.FormatFloat(k, 'f', 2, 64))
		case 'b':
			if r, ok := v[a].(bool); ok {
				if r {
					n, err = io.WriteString(b, "true")
				} else {
					n, err = io.WriteString(b, "false")
				}
			}
		case 'd', 'x', 'X':
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
				n, err = io.WriteString(b, strconv.FormatUint(k, 10))
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
