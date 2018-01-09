package console

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterh/liner"
)

var monkeyKeywords = []string{
	"fn", "let", "true", "false", "if", "else", "elsif", "elseif",
	"elif", "return", "include", "and", "or", "struct", "do", "while",
	"break", "continue", "for", "in", "where", "grep", "map", "case",
	"is", "try", "catch", "finally", "throw", "qw", "unless", "spawn",
	"enum", "defer", "nil",
}

//Note: we should put the longest operators first.
var monkeyOperators = []string{
	"+=", "-=", "*=", "/=", "%=", "^=",
	"++", "--",
	"&&", "||",
	"<<", ">>",
	"->", "=>",
	"==", "!=", "<=", ">=", "=~", "!~",
	"+", "-", "*", "/", "%", "^",
	"(", ")", "{", "}", "[", "]",
	"=", "<", ">",
	"!", "&", "|", ".",
	",", "?", ":", ";",
}

var colors = map[liner.Category]string{
	liner.NumberType:   liner.COLOR_YELLOW,
	liner.KeywordType:  liner.COLOR_MAGENTA,
	liner.StringType:   liner.COLOR_CYAN,
	liner.CommentType:  liner.COLOR_GREEN,
	liner.OperatorType: liner.COLOR_RED,
}

const PROMPT = "demo>> "

func Start(out io.Writer, color bool) {
	history := filepath.Join(os.TempDir(), ".liner_history")
	l := liner.NewLiner()
	defer l.Close()

	l.SetCtrlCAborts(true)
	l.SetSyntaxHighlight(color) //use syntax highlight or not
	l.RegisterKeywords(monkeyKeywords)
	l.RegisterOperators(monkeyOperators)
	l.RegisterColors(colors)

	if f, err := os.Open(history); err == nil {
		l.ReadHistory(f)
		f.Close()
	}

	var tmplines []string
	for {
		if line, err := l.Prompt(PROMPT); err == nil {
			if line == "exit" || line == "quit" {
				if f, err := os.Create(history); err == nil {
					l.WriteHistory(f)
					f.Close()
				}
				break
			}

			tmpline := strings.TrimSpace(line)
			if len(tmpline) == 0 { //empty line
				continue
			}
			//check if the `line` variable is ended with '\'
			if tmpline[len(tmpline)-1:] =="\\" { //the expression/statement has remaining part
				tmplines = append(tmplines, strings.TrimRight(tmpline, "\\"))
				continue
			} else {
				tmplines = append(tmplines, line)
			}

			resultLine := strings.Join(tmplines, "")
			l.AppendHistory(resultLine)
			tmplines = nil // clear the array
		}
	}
}
