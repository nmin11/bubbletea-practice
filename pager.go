package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type pagerModel struct {
	viewport viewport.Model
	title    string
	content  string
}

func (p pagerModel) setSize(w, h int) {
	p.viewport.Width = w
	p.viewport.Height = h
}

func (p *pagerModel) setViewportContent() {
	content := fmt.Sprintf("%s\n\n%s", p.title, p.content)
	str, err := glamour.Render(content, "dark")
	if err != nil {
		panic(err)
	}
	p.viewport.SetContent(str)
}

func (p pagerModel) update(cmd tea.Cmd, msg tea.Msg) (pagerModel, tea.Cmd) {
	p.viewport, cmd = p.viewport.Update(msg)
	return p, cmd
}

func (p pagerModel) helpView() string {
	return helpStyle("\n esc: back • q: quit")
}

func (p pagerModel) view() string {
	top, right, bottom, left := docStyle.GetMargin()
	p.viewport = viewport.New(windowSize.Width-left-right, windowSize.Height-top-bottom)
	p.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)
	p.setViewportContent()

	formatted := lipgloss.JoinVertical(lipgloss.Left, p.viewport.View(), p.helpView())
	return docStyle.Render(formatted)
}
