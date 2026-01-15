package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PasswordInput is a Bubble Tea model for password input
type PasswordInput struct {
	textInput textinput.Model
	prompt    string
	done      bool
	cancelled bool
}

func NewPasswordInput(prompt string) PasswordInput {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '*'
	ti.PromptStyle = InputPromptStyle
	ti.TextStyle = InputStyle
	ti.Cursor.Style = CursorStyle

	return PasswordInput{
		textInput: ti,
		prompt:    prompt,
	}
}

func (m PasswordInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m PasswordInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.done = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.cancelled = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m PasswordInput) View() string {
	promptStyle := lipgloss.NewStyle().
		Foreground(Accent).
		Bold(true)

	lockIcon := lipgloss.NewStyle().
		Foreground(Primary).
		SetString(IconLock + " ")

	helpText := DimStyle.Render("(press Enter to confirm, Esc to cancel)")

	return fmt.Sprintf(
		"\n%s%s\n\n  %s\n\n%s",
		lockIcon,
		promptStyle.Render(m.prompt),
		m.textInput.View(),
		helpText,
	)
}

func (m PasswordInput) Value() string {
	return m.textInput.Value()
}

func (m PasswordInput) Cancelled() bool {
	return m.cancelled
}

// TextInput is a Bubble Tea model for regular text input
type TextInput struct {
	textInput textinput.Model
	prompt    string
	icon      string
	done      bool
	cancelled bool
}

func NewTextInput(prompt, icon string) TextInput {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.PromptStyle = InputPromptStyle
	ti.TextStyle = InputStyle
	ti.Cursor.Style = CursorStyle

	return TextInput{
		textInput: ti,
		prompt:    prompt,
		icon:      icon,
	}
}

func (m TextInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.done = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.cancelled = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInput) View() string {
	promptStyle := lipgloss.NewStyle().
		Foreground(Accent).
		Bold(true)

	iconStyle := lipgloss.NewStyle().
		Foreground(Primary).
		SetString(m.icon + " ")

	helpText := DimStyle.Render("(press Enter to confirm, Esc to cancel)")

	return fmt.Sprintf(
		"\n%s%s\n\n  %s\n\n%s",
		iconStyle,
		promptStyle.Render(m.prompt),
		m.textInput.View(),
		helpText,
	)
}

func (m TextInput) Value() string {
	return m.textInput.Value()
}

func (m TextInput) Cancelled() bool {
	return m.cancelled
}

// PromptPassword runs an interactive password prompt
func PromptPassword(prompt string) (string, error) {
	p := tea.NewProgram(NewPasswordInput(prompt))
	m, err := p.Run()
	if err != nil {
		return "", err
	}

	model := m.(PasswordInput)
	if model.Cancelled() {
		return "", fmt.Errorf("cancelled")
	}

	return model.Value(), nil
}

// PromptText runs an interactive text prompt
func PromptText(prompt, icon string) (string, error) {
	p := tea.NewProgram(NewTextInput(prompt, icon))
	m, err := p.Run()
	if err != nil {
		return "", err
	}

	model := m.(TextInput)
	if model.Cancelled() {
		return "", fmt.Errorf("cancelled")
	}

	return model.Value(), nil
}

// RenderLogo renders the vaulta logo
func RenderLogo() string {
	return Logo.Render(LogoText)
}

// RenderSuccess renders a success message
func RenderSuccess(message string) string {
	content := fmt.Sprintf("%s %s", IconCheck, message)
	return SuccessBox.Render(content)
}

// RenderError renders an error message
func RenderError(message string) string {
	content := fmt.Sprintf("%s %s", IconCross, message)
	return ErrorBox.Render(content)
}

// RenderWarning renders a warning message
func RenderWarning(message string) string {
	content := fmt.Sprintf("%s %s", IconWarning, message)
	return WarningBox.Render(content)
}

// RenderInfo renders an info box
func RenderInfo(title, content string) string {
	titleRendered := TitleStyle.Render(title)
	return BoxStyle.Render(fmt.Sprintf("%s\n\n%s", titleRendered, content))
}

// RenderEntry renders a vault entry nicely
func RenderEntry(name, username, password string) string {
	title := EntryTitleStyle.Render(fmt.Sprintf("%s  %s", IconKey, name))

	usernameRow := fmt.Sprintf("%s %s",
		EntryLabelStyle.Render("Username:"),
		EntryValueStyle.Render(username),
	)

	passwordRow := fmt.Sprintf("%s %s",
		EntryLabelStyle.Render("Password:"),
		EntryValueStyle.Render(password),
	)

	content := fmt.Sprintf("%s\n\n%s\n%s", title, usernameRow, passwordRow)
	return EntryBoxStyle.Render(content)
}

// RenderList renders a list of entries
func RenderList(title string, items []string) string {
	titleRendered := TitleStyle.Render(fmt.Sprintf("%s  %s", IconList, title))

	var listItems []string
	for _, item := range items {
		bullet := lipgloss.NewStyle().Foreground(Accent).Render("◆")
		itemText := ValueStyle.Render(item)
		listItems = append(listItems, fmt.Sprintf("  %s %s", bullet, itemText))
	}

	if len(listItems) == 0 {
		listItems = append(listItems, DimStyle.Render("  No entries found"))
	}

	content := fmt.Sprintf("%s\n\n%s", titleRendered, strings.Join(listItems, "\n"))
	return BoxStyle.Render(content)
}

// RenderDivider renders a styled divider
func RenderDivider() string {
	return DimStyle.Render(strings.Repeat("─", 50))
}

// ClearScreen clears the terminal (best effort)
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
