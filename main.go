package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var globalConfig config

const maxWidth = 80

// --primary-1: #A87676;
// --primary-2: #CA8787;
// --primary-3: #E1ACAC;
// --primary-4: #FFD0D0;
var (
	primary        = lipgloss.AdaptiveColor{Light: "#A87676", Dark: "#FFD0D0"}
	secondary      = lipgloss.AdaptiveColor{Light: "#Ca8787", Dark: "#E1ACAC"}
	primaryLight   = lipgloss.Color("#F7DBDB")
	primaryFill    = lipgloss.NewStyle().Background(primaryLight).Foreground(primaryLight)
	primaryOutline = lipgloss.NewStyle().Foreground(primary)
	green          = lipgloss.Color("#5E865E")
	lightGreen     = lipgloss.Color("#AFC6AF")
	greenFill      = lipgloss.NewStyle().Background(lightGreen)
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func newTheme() *huh.Theme {

	theme := *huh.ThemeBase()

	theme.Blurred = huh.FieldStyles{
		TextInput: huh.TextInputStyles{
			Prompt: primaryOutline,
			Cursor: primaryOutline,
		},
	}

	theme.Focused = huh.FieldStyles{
		TextInput: huh.TextInputStyles{
			Prompt: primaryOutline,
			Cursor: primaryOutline,
			Text:   primaryOutline,
		},
	}

	return &theme
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Background(primary).
		Foreground(primaryLight).
		Padding(0, 1, 0, 2)

	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primary).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(secondary).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(primary)
	s.ErrorHeaderText = s.HeaderText.Copy().
		Foreground(secondary)
	s.Help = lg.NewStyle().
		Foreground(primary)
	return &s
}

type state struct {
	unit1        string
	unit2        string
	doNotInclude string
	index        int
}

type Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
	state  state
}

var acceptedUnits = []string{"rem", "px", "%", "em"}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			// huh.NewInput().
			// 	Title("Unit").
			// 	Key("first").
			// 	Prompt("? ").
			// 	Placeholder(strings.Join(acceptedUnits, " ")).
			// 	Suggestions(acceptedUnits),
			// huh.NewInput().
			// 	Title("Unit").
			// 	Key("second").
			// 	Prompt("? ").
			// 	Placeholder(strings.Join(acceptedUnits, " ")).
			// 	Suggestions(acceptedUnits),
			huh.NewInput().
				Title(primaryOutline.Render("Conversion Factor")).
				Key("factor").
				Prompt("? ").
				Placeholder("number"),
			huh.NewText().
				Key("do-not-include").
				Title(primaryOutline.Render("Do Not Include List")),
		)).
		WithWidth(40).
		WithShowHelp(false).
		WithShowErrors(false)

	keyMap := huh.NewDefaultKeyMap()

	keyMap.Input = huh.InputKeyMap{
		Prev:   key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "up")),
		Next:   key.NewBinding(key.WithKeys("down", "tab"), key.WithHelp("↓", "down")),
		Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	}

	keyMap.Text = huh.TextKeyMap{
		Prev:   key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "up")),
		Next:   key.NewBinding(key.WithKeys("down", "tab"), key.WithHelp("↓", "down")),
		Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	}

	m.form.WithKeyMap(keyMap).WithTheme(newTheme())

	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func debug(msg string) {
	os.WriteFile("debug.txt", []byte(msg), 0755)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.state.index > 0 {
				m.state.index--
			}
		case "down":
			if m.state.index <= 2 {
				m.state.index++
			}
		case "backspace":
		default:

		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	// Form (left side)
	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	// Status (right side)
	conversion := m.form.GetString("factor")
	conversionFactor, err := strconv.ParseFloat(conversion, 64)

	if err != nil {
		conversionFactor = 1
	}

	number := 24.0
	newNumber := conversionFactor * number
	first := primaryFill.Render(fmt.Sprintf("padding: %fpx", number))
	second := greenFill.Render(fmt.Sprintf("padding: %frem", newNumber))
	content := fmt.Sprintf("%s\n\n%sx\n\n%s", first, conversion, second)

	const statusWidth = 38
	statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
	status := s.Status.Copy().
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render(s.StatusHeader.Render("Example\n\n") + "\n" + content)

	header := m.appBoundaryView("Pixel 2 Rem Pro")

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, status)

	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

	if m.form.State == huh.StateCompleted {

		globalConfig = config{conversionFactor: conversionFactor, doNotInclude: m.form.GetString("do-not-include")}

		return ""
	}

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)

}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
	)
}

func main() {
	_, err := tea.NewProgram(NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// charmInterface(config{conversionFactor: .1, doNotInclude: "padding"})
	charmInterface(globalConfig)
}
