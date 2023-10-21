package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ListItem struct {
	Label, Desc string
	Value       interface{}
}

func (i ListItem) Title() string       { return i.Label }
func (i ListItem) Description() string { return i.Desc }
func (i ListItem) FilterValue() string { return i.Label }

type model struct {
	list   list.Model
	choice interface{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			m.choice = m.list.Items()[m.list.Cursor()].(ListItem).Value
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func NewList(items []ListItem) interface{} {

	listItems := []list.Item{}

	for _, item := range items {
		listItems = append(listItems, item)
	}

	m := model{list: list.New(listItems, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Install Options"

	p := tea.NewProgram(m, tea.WithAltScreen())

	mo, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if md, ok := mo.(model); ok && md.choice != "" {
		return md.choice
	}

	return nil

}
