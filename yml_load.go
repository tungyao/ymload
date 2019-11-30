package ymload

import (
	"bufio"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func Format(filename string) map[string]map[string]interface{} {
	if filename[len(filename)-4:] != ".yml" {
		filename += ".yml"
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Panicln(err)
	}
	mp := make(map[string]map[string]interface{})
	buf := bufio.NewReader(f)
	last := ""
	for {
		line, err := buf.ReadString('\r')
		line = strings.TrimSpace(line)
		if line[len(line)-1:] == ":" {
			mp[line[:len(line)-1]] = make(map[string]interface{})
			last = line[:len(line)-1]
		} else {
			s := SplitString([]byte(line), []byte(": "))
			mp[last][string(format(s[0]))] = string(s[1])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	return mp
}
func SplitString(str []byte, p []byte) [][]byte {
	group := make([][]byte, 0)
	ps := 0
	for i := 0; i < len(str); i++ {
		if str[i] == p[0] && i < len(str)-len(p) {
			if len(p) == 1 {
				group = append(group, str[ps:i])
				ps = i + len(p)
				//return [][]byte{str[:i], str[i+1:]}
			} else {
				for j := 1; j < len(p); j++ {
					if str[i+j] != p[j] {
						continue
					} else {
						group = append(group, str[ps:i])
						ps = i + len(p)
					}
					//return [][]byte{str[:i], str[i+len(p):]}
				}
			}
		} else {
			continue
		}
	}
	group = append(group, str[ps:])
	return group
}
func format(b []byte) []byte {
	out := make([]byte, 0)
	for _, v := range b {
		if v != ' ' {
			out = append(out, v)
		}
	}
	return out
}
func CheckConfig(nw interface{}, deft interface{}) {
	switch reflect.TypeOf(nw).Kind() {
	case reflect.Struct:
		t := reflect.TypeOf(nw).Elem()
		v := reflect.ValueOf(nw).Elem()
		for i := 0; i < t.NumField(); i++ {
			n := v.Field(i)
			switch n.Kind() {
			case reflect.String:
				if n.IsZero() {
					n.SetString(reflect.ValueOf(deft).Field(i).String())
				}
			case reflect.Int:
				if n.IsZero() {
					n.SetInt(reflect.ValueOf(deft).Field(i).Int())
				}
			case reflect.Int64:
				if n.IsZero() {
					n.SetInt(reflect.ValueOf(deft).Field(i).Int())
				}
			case reflect.Bool:
				if n.IsZero() {
					n.SetBool(reflect.ValueOf(deft).Field(i).Bool())
				}
			case reflect.Float64:
				if n.IsZero() {
					n.SetFloat(reflect.ValueOf(deft).Field(i).Float())
				}
			}
		}
	case reflect.Ptr:
		n := reflect.ValueOf(nw).Elem()
		switch n.Kind() {
		case reflect.Int:
			if n.IsZero() && n.CanSet() {
				n.SetInt(reflect.ValueOf(deft).Int())
			}
		case reflect.String:
			if n.IsZero() && n.CanSet() {
				n.SetString(reflect.ValueOf(deft).String())
			}
		case reflect.Struct:
			t := reflect.TypeOf(nw).Elem()
			v := reflect.ValueOf(nw).Elem()
			for i := 0; i < t.NumField(); i++ {
				n := v.Field(i)
				switch n.Kind() {
				case reflect.String:
					if n.IsZero() {
						n.SetString(reflect.ValueOf(deft).Field(i).String())
					}
				case reflect.Int:
					if n.IsZero() {
						n.SetInt(reflect.ValueOf(deft).Field(i).Int())
					}
				case reflect.Int64:
					if n.IsZero() {
						n.SetInt(reflect.ValueOf(deft).Field(i).Int())
					}
				case reflect.Bool:
					if n.IsZero() {
						n.SetBool(reflect.ValueOf(deft).Field(i).Bool())
					}
				case reflect.Float64:
					if n.IsZero() {
						n.SetFloat(reflect.ValueOf(deft).Field(i).Float())
					}
				}
			}
		}

	}
}

