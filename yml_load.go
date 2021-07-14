package ymload

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func Format(filename string) map[string]map[string]string {
	if filename[len(filename)-4:] != ".yml" {
		filename += ".yml"
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Panicln(err)
	}
	mp := make(map[string]map[string]string)
	buf := bufio.NewReader(f)
	last := ""
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		if line[0] == '#' {
			continue
		}
		if line[len(line)-1:] == ":" {
			mp[line[:len(line)-1]] = make(map[string]string)
			last = line[:len(line)-1]
			continue
		} else {
			var isExec bool
			for i := 0; i < len(line); i++ {
				// 依次解析每一行
				if isExec {
					break
				}
				switch line[i] {
				case ':':
					if i+2 >= len(line) {
						mp[last][format(line[:i])] = ""
						isExec = true
						break
					}
					mp[last][format(line[:i])] = getString(line[i+2:])
					isExec = true
					break
				}
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	return mp
}

func getString(s string) string {
	for i, _ := range s {
		var lst int
		switch s[i] {
		case '"':
			lst = len(s) - 2
			if lst < 0 {
				lst = 0
			}
			i += 1
			if i > lst {
				i = 0
			}
			return s[i:lst]
		case '\'':
			lst = len(s) - 2
			if lst < 0 {
				lst = 0
			}
			i += 1
			println(i, lst)
			if i > lst {
				i = 0
			}
			return s[i:lst]
		default:
			return s
		}
	}
	return ""
}
func SplitString(str []byte, p []byte) [][]byte {
	group := make([][]byte, 0)
	ps := 0
	for i := 0; i < len(str); i++ {
		if str[i] == p[0] && i < len(str)-len(p) {
			if len(p) == 1 {
				group = append(group, str[ps:i])
				ps = i + len(p)
			} else {
				for j := 1; j < len(p); j++ {
					if str[i+j] != p[j] {
						continue
					} else {
						group = append(group, str[ps:i])
						ps = i + len(p)
					}
				}
			}
		} else {
			continue
		}
	}
	group = append(group, str[ps:])
	return group
}
func format(b string) string {
	out := make([]byte, 0)
	for _, v := range b {
		if v != ' ' {
			out = append(out, uint8(v))
		}
	}
	return bytes.NewBuffer(out).String()
}
