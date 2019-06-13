package linetocmd

import (
	"errors"
	"os/exec"
)

func Parse(line string) (*exec.Cmd, error) {
	args, err := ParseToArray(line)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(args[0])
	if len(args) > 1 {
		cmd.Args = append(cmd.Args, args[1:]...)
	}
	return cmd, nil
}

func ParseToArray(line string) ([]string, error) {
	args := make([]string, 0)
	slice := make([]rune, 0)
	singleQuote := false
	doubleQuote := false
	for i := 0; i < len(line); i++ {
		item := rune(line[i])
		switch item {
		case '\\':
			i++
			if i == len(line) {
				return nil, errors.New("bad line line")
			}
			slice = append(slice, rune(line[i]))
		case ' ':
			if singleQuote || doubleQuote {
				slice = append(slice, item)
			} else {
				// 追加到参数中，清空碎片信息
				if len(slice) == 0 {
					continue
				}
				args = append(args, string(slice))
				slice = make([]rune, 0)
			}
		case '\'':
			if singleQuote {
				args = append(args, string(slice))
				slice = make([]rune, 0)
				singleQuote = false
			} else if !doubleQuote {
				singleQuote = true
			}
		case '"':
			if doubleQuote {
				args = append(args, string(slice))
				slice = make([]rune, 0)
				doubleQuote = false
			} else if !singleQuote {
				doubleQuote = true
			}
		default:
			slice = append(slice, item)
		}
	}
	if len(slice) > 0 {
		args = append(args, string(slice))
	}
	return args, nil
}
