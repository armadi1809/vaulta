package ui

import "github.com/charmbracelet/lipgloss"

// Color palette - modern cyberpunk-inspired theme
var (
	Primary    = lipgloss.Color("#7C3AED") // Violet
	Secondary  = lipgloss.Color("#10B981") // Emerald
	Accent     = lipgloss.Color("#F59E0B") // Amber
	Danger     = lipgloss.Color("#EF4444") // Red
	Subtle     = lipgloss.Color("#6B7280") // Gray
	Text       = lipgloss.Color("#F9FAFB") // White-ish
	TextDim    = lipgloss.Color("#9CA3AF") // Gray
	Background = lipgloss.Color("#1F2937") // Dark gray
	Highlight  = lipgloss.Color("#A78BFA") // Light violet
)

// Logo and branding
var Logo = lipgloss.NewStyle().
	Bold(true).
	Foreground(Primary).
	MarginBottom(1)

var LogoText = `
 ██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗ █████╗ 
 ██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██╔══██╗
 ██║   ██║███████║██║   ██║██║     ██║   ███████║
 ╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██╔══██║
  ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║  ██║
   ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝  ╚═╝`

// Box styles
var BoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Primary).
	Padding(1, 2).
	MarginTop(1)

var SuccessBox = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Secondary).
	Foreground(Secondary).
	Padding(1, 2).
	MarginTop(1)

var ErrorBox = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Danger).
	Foreground(Danger).
	Padding(1, 2).
	MarginTop(1)

var WarningBox = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Accent).
	Foreground(Accent).
	Padding(1, 2).
	MarginTop(1)

// Text styles
var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(Primary).
	MarginBottom(1)

var SubtitleStyle = lipgloss.NewStyle().
	Foreground(TextDim).
	Italic(true)

var LabelStyle = lipgloss.NewStyle().
	Foreground(Highlight).
	Bold(true)

var ValueStyle = lipgloss.NewStyle().
	Foreground(Text)

var DimStyle = lipgloss.NewStyle().
	Foreground(Subtle)

var SuccessStyle = lipgloss.NewStyle().
	Foreground(Secondary).
	Bold(true)

var ErrorStyle = lipgloss.NewStyle().
	Foreground(Danger).
	Bold(true)

// List styles
var ListItemStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var ListSelectedStyle = lipgloss.NewStyle().
	Foreground(Primary).
	Bold(true).
	PaddingLeft(2)

var ListBullet = lipgloss.NewStyle().
	Foreground(Accent).
	SetString("◆ ")

var ListBulletSelected = lipgloss.NewStyle().
	Foreground(Primary).
	Bold(true).
	SetString("▶ ")

// Input styles
var InputPromptStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true)

var InputStyle = lipgloss.NewStyle().
	Foreground(Text)

var CursorStyle = lipgloss.NewStyle().
	Foreground(Primary)

// Table styles
var TableHeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(Primary).
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true).
	BorderForeground(Subtle).
	PaddingRight(2)

var TableCellStyle = lipgloss.NewStyle().
	Foreground(Text).
	PaddingRight(2)

// Entry display
var EntryBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Primary).
	Padding(1, 2).
	MarginTop(1).
	MarginBottom(1)

var EntryTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(Primary).
	MarginBottom(1)

var EntryLabelStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Width(12)

var EntryValueStyle = lipgloss.NewStyle().
	Foreground(Text)

// Help text
var HelpStyle = lipgloss.NewStyle().
	Foreground(Subtle).
	MarginTop(1)

// Spinner/Loading
var SpinnerStyle = lipgloss.NewStyle().
	Foreground(Primary)

// Icons
const (
	IconLock    = "🔒"
	IconUnlock  = "🔓"
	IconKey     = "🔑"
	IconCheck   = "✓"
	IconCross   = "✗"
	IconWarning = "⚠"
	IconInfo    = "ℹ"
	IconArrow   = "→"
	IconBullet  = "•"
	IconStar    = "★"
	IconShield  = "🛡"
	IconVault   = "🗄"
	IconAdd     = "+"
	IconDelete  = "−"
	IconList    = "☰"
	IconSearch  = "🔍"
)
