// Package: ui
// Estilos y colores corporativos para XEBEC CORPORATION
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Colores corporativos XEBEC
var (
	// Primary colors
	CorporateBlue  = lipgloss.Color("#00AEEF")
	CorporateBlack = lipgloss.Color("#0A0A0A")
	CorporateWhite = lipgloss.Color("#FFFFFF")
	
	// Accent colors
	AccentCyan    = lipgloss.Color("#00BCD4")
	AccentGreen   = lipgloss.Color("#4CAF50")
	AccentOrange  = lipgloss.Color("#FF9800")
	AccentRed     = lipgloss.Color("#F44336")
	AccentYellow  = lipgloss.Color("#FFEB3B")
	
	// Neutral colors
	GrayDark  = lipgloss.Color("#1a1a1a")
	Gray      = lipgloss.Color("#333333")
	GrayLight = lipgloss.Color("#666666")
	GrayLighter = lipgloss.Color("#999999")
)

// Estilos reutilizables
var (
	// Estilo para el título principal
	TitleStyle = lipgloss.NewStyle().
			Foreground(CorporateBlue).
			Bold(true).
			Padding(0, 1)
	
	// Estilo para subtítulos
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(CorporateWhite).
			Bold(false)
	
	// Estilo para texto normal
	NormalTextStyle = lipgloss.NewStyle().
				Foreground(CorporateWhite)
	
	// Estilo para texto muted
	MutedTextStyle = lipgloss.NewStyle().
			Foreground(GrayLighter)
	
	// Estilo para highlight
	HighlightStyle = lipgloss.NewStyle().
			Foreground(CorporateBlue).
			Bold(true)
	
	// Estilo para éxito
	SuccessStyle = lipgloss.NewStyle().
			Foreground(AccentGreen)
	
	// Estilo para error
	ErrorStyle = lipgloss.NewStyle().
			Foreground(AccentRed)
	
	// Estilo para warning
	WarningStyle = lipgloss.NewStyle().
			Foreground(AccentOrange)
	
	// Estilo para opciones de menú seleccionadas
	SelectedOptionStyle = lipgloss.NewStyle().
				Foreground(CorporateBlue).
				Bold(true).
				Background(GrayDark)
	
	// Estilo para opciones de menú no seleccionadas
	UnselectedOptionStyle = lipgloss.NewStyle().
				Foreground(GrayLight)
	
	// Estilo para bordes
	BorderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(CorporateBlue)
	
	// Estilo para caja con borde
	BoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(CorporateBlue).
			Padding(1, 2)
	
	// Estilo para información del sistema
	InfoStyle = lipgloss.NewStyle().
			Foreground(GrayLighter).
			Padding(0, 1)
	
	// Estilo para separadores
	SeparatorStyle = lipgloss.NewStyle().
			Foreground(Gray)
	
	// Estilo para el prompt
	PromptStyle = lipgloss.NewStyle().
			Foreground(CorporateBlue)
)
