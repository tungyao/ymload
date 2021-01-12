package ymload

import (
	"bufio"
	"io"
	"log"
	"os"
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
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
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
