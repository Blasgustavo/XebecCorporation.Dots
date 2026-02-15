// Package: ui
// Estilos y colores corporativos para XEBEC CORPORATION
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Colores corporativos XEBEC (cargados desde branding)
var (
	// Primary colors
	CorporateBlue  = lipgloss.Color(BrandingConfig.Colors.Primary)
	CorporateBlack = lipgloss.Color(BrandingConfig.Colors.Secondary)
	CorporateWhite = lipgloss.Color(BrandingConfig.Colors.White)

	// Accent colors
	AccentCyan   = lipgloss.Color(BrandingConfig.Colors.AccentCyan)
	AccentGreen  = lipgloss.Color(BrandingConfig.Colors.AccentGreen)
	AccentOrange = lipgloss.Color(BrandingConfig.Colors.AccentOrange)
	AccentRed    = lipgloss.Color(BrandingConfig.Colors.AccentRed)
	AccentYellow = lipgloss.Color(BrandingConfig.Colors.AccentYellow)
	AccentPurple = lipgloss.Color(BrandingConfig.Colors.AccentPurple)

	// Gradient colors
	GradientStart = lipgloss.Color(BrandingConfig.Colors.GradientStart)
	GradientEnd   = lipgloss.Color(BrandingConfig.Colors.GradientEnd)

	// Neutral colors
	GrayDark    = lipgloss.Color(BrandingConfig.Colors.GrayDark)
	Gray        = lipgloss.Color(BrandingConfig.Colors.Gray)
	GrayLight   = lipgloss.Color(BrandingConfig.Colors.GrayLight)
	GrayLighter = lipgloss.Color(BrandingConfig.Colors.GrayLighter)
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
