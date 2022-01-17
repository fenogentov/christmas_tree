package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/term"
)

var (
	title = "\033[1;34m С НОВЫМ ГОДОМ !"
	ball  = '⏺'
	color = []string{
		"\033[94m",
		"\033[93m",
		"\033[96m",
		"\033[92m",
		"\033[95m",
		"\033[97m",
		"\033[91m",
	}

	size, width int

	colStar string
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	for {
		width, size, _ = term.GetSize(0)

		Tree := tree(size-5, width/2)
		Tree = balls(Tree)
		Tree = colored(Tree)

		t := strings.Join(Tree, "\n")

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		fmt.Println(t)

		time.Sleep(1 * time.Second)
	}
}

// tree is creates a Christmas tree
func tree(height, screenWidth int) []string {
	if height%2 != 0 {
		height++
	}
	top := []string{title, ` ★ `, `★ ★ ★`, ` ★ ★ `, "/_\\", "/_\\_\\"}
	lines := []string{}
	trunk := "[___]"
	begin := "/"
	end := "\\"
	pattern := "_/"
	j := 5
	lines = append(lines, top...)
	for i := 7; i < height+1; i += 2 {
		middle := pattern + strings.Repeat(pattern, i-j)
		line := begin + middle[:len(middle)-1] + end
		lines = append(lines, line)
		middle1 := strings.ReplaceAll(middle, "/", "\\")
		line1 := begin + middle1[:len(middle1)-1] + end
		lines = append(lines, line1)
		j++

	}
	lines = append(lines, trunk)
	lineCenter(lines, screenWidth)
	return lines
}

// balls is hangs balls
func balls(tree []string) []string {
	for idx := 3; idx < len(tree)-1; idx++ {
		tree[idx] = randomHangsBalls(tree[idx])
	}
	return tree
}
func randomHangsBalls(str string) string {
	r := []rune(str)
	for i := 0; i < len(r)/6; i++ {
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := generator.Intn(len(r))
		if r[idx] != ' ' && r[idx] == '_' {
			r[idx] = ball
		}
	}
	return string(r)
}

// colored is paints balls
func colored(str []string) []string {
	if colStar == "\033[31m" {
		colStar = "\033[33m"
	} else {
		colStar = "\033[31m"
	}
	for i, s := range str {
		var buffer bytes.Buffer
		for _, r := range s {
			if r == ball {
				genColor := rand.New(rand.NewSource(time.Now().UnixNano()))
				buffer.WriteString(color[genColor.Intn(len(color))])
				buffer.WriteRune(ball)
				buffer.WriteString("\033[32m")
				continue
			}
			if r == '★' {
				buffer.WriteString(colStar)
				buffer.WriteRune(r)
				buffer.WriteString("\033[32m")
				continue
			}
			buffer.WriteRune(r)
		}
		str[i] = buffer.String()
	}
	return str
}

// lineCenter sets characters in center string
func lineCenter(line []string, w int) {
	for i, s := range line {
		s = strings.Repeat(" ", (w-utf8.RuneCountInString(s)/2)) + s
		line[i] = s
	}
}
