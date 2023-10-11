package main

import (
	"fmt"
	"os"

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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		windowSize = msg
		m.list.list.SetSize(msg.Width, msg.Height-2)
		m.pager.setSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch m.state {
		case stateList:
			if key.Matches(msg, Keymap.Enter) && m.list.state == listNormal {
				m.state = statePager
				m.pager.title = m.list.list.SelectedItem().(item).title
				m.pager.content = m.list.list.SelectedItem().(item).desc
				m.pager.setViewportContent()
				m.pager, cmd = m.pager.update(cmd, msg)
			} else if key.Matches(msg, Keymap.Filter) && m.list.state == listNormal {
				m.list.state = listFiltering
				m.list, cmd = m.list.update(cmd, msg)
			} else {
				m.list, cmd = m.list.update(cmd, msg)
			}

		case statePager:
			if key.Matches(msg, Keymap.Back) {
				m.state = stateList
				m.pager.viewport.SetContent("")
				m.pager.viewport.YOffset = 0
				m.pager, cmd = m.pager.update(cmd, msg)
			} else {
				m.pager, cmd = m.pager.update(cmd, msg)
			}
		}
	}

	return m, cmd
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
	m := model{
		state: stateList,
		list:  initList(),
		pager: pagerModel{},
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
