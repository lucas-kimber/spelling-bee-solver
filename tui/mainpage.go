package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	solver "github.com/lucas-kimber/spelling-bee-solver/solver"
)

// MODEL DATA

type simplePage struct {
	solver         solver.Solver
	inputLetters   [7]string
	cursorPosition int
}

func NewSimplePage() simplePage {
	return simplePage{
		solver:         *solver.NewSolver(),
		cursorPosition: 0,
	}
}

func (s simplePage) Init() tea.Cmd { return nil }

// VIEW

func (s simplePage) View() string {
	var boxes []string
	for i := 0; i < 7; i++ {
		ch := s.inputLetters[i]
		if ch == "" {
			ch = " "
		}
		if i == s.cursorPosition {
			boxes = append(boxes, fmt.Sprintf("[%s]", ch))
		} else {
			boxes = append(boxes, fmt.Sprintf(" %s ", ch))
		}
	}
	return fmt.Sprintf(
		"Enter letters:\n%s\n\nUse arrow keys to move. Ctrl+C to exit.",
		strings.Join(boxes, " "),
	)
}

// UPDATE

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "ctrl+c":
			return s, tea.Quit
		case "left":
			if s.cursorPosition > 0 {
				s.cursorPosition--
			}
		case "right":
			if s.cursorPosition < 6 {
				s.cursorPosition++
			}
		case "backspace", "delete":
			s.inputLetters[s.cursorPosition] = ""
		default:
			if len(m.String()) == 1 && m.Type == tea.KeyRunes {
				s.inputLetters[s.cursorPosition] = m.String()
				if s.cursorPosition < 6 {
					s.cursorPosition++
				}
			}
		}
	}
	return s, nil
}
