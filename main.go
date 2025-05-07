package main

import (
	tea "github.com/charmbracelet/bubbletea"
	tui "github.com/lucas-kimber/spelling-bee-solver/tui"
)

func main() {
	p := tea.NewProgram(
		tui.NewSimplePage(),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
