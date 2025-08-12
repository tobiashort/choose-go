package choose

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"

	"github.com/tobiashort/ansi-go"
	"github.com/tobiashort/cfmt-go"
	"github.com/tobiashort/orderedmap-go"

	. "github.com/tobiashort/utils-go/must"
)

func Multiple(prompt string, options []string) ([]string, bool) {
	oldState := Must2(term.MakeRaw(int(os.Stdin.Fd())))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	ok := false
	activeIndex := 0

	selection := orderedmap.NewOrderedMap[string, bool]()
	for _, option := range options {
		selection.Put(option, false)
	}

draw:
	fmt.Printf("%s\r\n", prompt)
	for index, option := range options {
		if index == activeIndex {
			if selected, _ := selection.Get(option); selected {
				cfmt.Printf(" #yB{> [x] %s}\r\n", option)
			} else {
				cfmt.Printf(" #yB{> [ ] %s}\r\n", option)
			}
		} else {
			if selected, _ := selection.Get(option); selected {
				fmt.Printf("   [x] %s\r\n", option)
			} else {
				fmt.Printf("   [ ] %s\r\n", option)
			}
		}
	}

	buf := make([]byte, 3)
	for {
		n := Must2(os.Stdin.Read(buf))
		switch string(buf[:n]) {
		case "j":
			fallthrough
		case ansi.InputKeyDown:
			if activeIndex < len(options)-1 {
				activeIndex++
			} else {
				activeIndex = 0
			}
			goto redraw
		case "k":
			fallthrough
		case ansi.InputKeyUp:
			if activeIndex > 0 {
				activeIndex--
			} else {
				activeIndex = len(options) - 1
			}
			goto redraw
		case ansi.InputSpace:
			option := options[activeIndex]
			selected, _ := selection.Get(option)
			selection.Put(option, !selected)
			goto redraw
		case ansi.InputCR:
			fallthrough
		case ansi.InputLF:
			fallthrough
		case ansi.InputCRLF:
			ok = true
			goto done
		case "q":
			fallthrough
		case ansi.InputEscape:
			ok = false
			goto done
		case ansi.InputCtrlC:
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			goto done
		}
	}

redraw:
	fmt.Print(ansi.MoveCursorUp(len(options) + 1))
	goto draw

done:
	selectedOptions := make([]string, 0)
	for option, selected := range selection.Iterate() {
		if selected {
			selectedOptions = append(selectedOptions, option)
		}
	}

	return selectedOptions, ok
}
