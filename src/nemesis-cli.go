package src

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func NemesisCli() (err error) {
	m := NewModel()

	p := tea.NewProgram(m)

	_, err = p.Run()

	if err != nil {
		log.Fatal("Error running the program: ", err)
	}
	return nil
}

type Model struct {
	title string

	textinput textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Type something..."
	ti.Focus()
	return Model{
		title:     "Nemesis CLI",
		textinput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := m.textinput.View()
	return s
}
