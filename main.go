package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	stateList state = iota
	statePager
)

func (s state) String() string {
	return map[state]string{
		stateList:  "showing list",
		statePager: "showing pager",
	}[s]
}

type model struct {
	state state
	list  listModel
	pager pagerModel
}

var cmd tea.Cmd

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		windowSize = msg
		m.list.list.SetSize(msg.Width, msg.Height)
		m.pager.setSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keymap.Enter):
			if m.state == stateList {
				m.state = statePager
				m.pager.title = m.list.list.SelectedItem().(item).title
				m.pager.content = m.list.list.SelectedItem().(item).desc
				m.pager.setViewportContent()
				_, cmd = m.pager.update(msg)
				m.list.update(msg)
				return m, cmd
			}

		case key.Matches(msg, Keymap.Back):
			if m.state == statePager {
				m.state = stateList
				m.pager.viewport.SetContent("")
				m.pager.viewport.YOffset = 0
				return m, nil
			}

		case key.Matches(msg, Keymap.Quit):
			return m, tea.Quit

		default:
			switch m.state {
			case stateList:
				m.list, cmd = m.list.update(msg)
			case statePager:
				m.pager, cmd = m.pager.update(msg)
			}
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case statePager:
		return m.pager.view()
	default:
		return m.list.view()
	}
}

func main() {
	m := model{}
	m.list = m.list.init()
	m.pager = pagerModel{}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
