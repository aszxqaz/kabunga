package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type listItem struct {
	pkg *PackageInfo
}

func (i listItem) Title() string       { return i.pkg.Name }
func (i listItem) Description() string { return "(" + i.pkg.Url + ")" }
func (i listItem) FilterValue() string { return i.pkg.Name }

type model struct {
	list list.Model
	pkg  *PackageInfo
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
			m.pkg = m.list.SelectedItem().(listItem).pkg
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

type PackageNamePrompter interface {
	Prompt(packages []PackageInfo) (*PackageInfo, error)
}

type packageNamePrompter struct {
	program *tea.Program
}

func DefaultPrompter() *packageNamePrompter {
	p := &packageNamePrompter{}
	return p
}

func (p *packageNamePrompter) init(packages []PackageInfo) {
	items := make([]list.Item, 0)
	for _, pkg := range packages {
		items = append(items, listItem{pkg: &pkg})
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	m := &model{list: l}
	m.list.Title = ""
	m.list.InfiniteScrolling = false
	m.list.SetShowTitle(false)

	p.program = tea.NewProgram(m, tea.WithAltScreen())
}

func (p *packageNamePrompter) Prompt(packages []PackageInfo) (*PackageInfo, error) {
	p.init(packages)
	tm, err := p.program.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	m := tm.(model)
	return m.pkg, nil
}
