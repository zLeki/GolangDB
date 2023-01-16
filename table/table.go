package table

import (
	"GopherDB/select"
	"GopherDB/types"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "left":
			types.Page--
			return m, tea.Quit

		case "enter":
			var options = []string{"Delete", "Update"}
			DropDown.DropDown(options, m.table.SelectedRow()[0])
			//fmt.Println("Selected row:", m.table.SelectedRow()[0])
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func Table(columnsText []string, rowsText []string) {
	columns := []table.Column{}
	for _, v := range columnsText {
		columns = append(columns, table.Column{Title: v, Width: 10})
	}
	types.TablePrimary = columnsText[0]

	rows := []table.Row{
		//{"100", "Montreal", "Canada", "4,276,526"},
	}
	for _, v := range rowsText {
		a := strings.Split(v, "|")
		row := table.Row{}
		for _, v := range a {
			if v == "" {
				continue
			}

			row = append(row, v)

		}
		rows = append(rows, row)
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
