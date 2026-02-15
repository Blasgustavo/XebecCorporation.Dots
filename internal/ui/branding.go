// Package: ui
// Sistema de Branding - Carga configuración desde assets/branding.json
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Estructuras para el branding
type Branding struct {
	Name      string                          `json:"name"`
	Version   string                          `json:"version"`
	Logo      string                          `json:"logo"`
	Separator string                          `json:"separator"`
	Colors    Colors                          `json:"colors"`
	Texts     Texts                           `json:"texts"`
	MenuOpts  []MenuOptionBranding            `json:"menu_options"`
	Submenus  map[string][]MenuOptionBranding `json:"submenus"`
}

type Colors struct {
	Primary       string `json:"primary"`
	Secondary     string `json:"secondary"`
	White         string `json:"white"`
	AccentCyan    string `json:"accent_cyan"`
	AccentGreen   string `json:"accent_green"`
	AccentOrange  string `json:"accent_orange"`
	AccentRed     string `json:"accent_red"`
	AccentYellow  string `json:"accent_yellow"`
	AccentPurple  string `json:"accent_purple"`
	GradientStart string `json:"gradient_start"`
	GradientEnd   string `json:"gradient_end"`
	GrayDark      string `json:"gray_dark"`
	Gray          string `json:"gray"`
	GrayLight     string `json:"gray_light"`
	GrayLighter   string `json:"gray_lighter"`
}

type Texts struct {
	CLILabel      string `json:"cli_label"`
	PlatformLabel string `json:"platform_label"`
	MenuTitle     string `json:"menu_title"`
	FooterNav     string `json:"footer_navigation"`
	FooterBack    string `json:"footer_back"`
	PromptSel     string `json:"prompt_selection"`
	OptionInvalid string `json:"option_invalid"`
	Goodbye       string `json:"goodbye"`
	Executing     string `json:"executing"`
	Back          string `json:"back"`
}

type MenuOptionBranding struct {
	ID          string `json:"id"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Submenu     bool   `json:"submenu"`
	Type        string `json:"type"` // "table", "action", etc.
}

// Variable global con el branding cargado
var BrandingConfig = loadBranding()

// Cargar branding desde archivo JSON
func loadBranding() Branding {
	// Buscar el archivo de branding
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
	brandingPath := filepath.Join(projectRoot, "assets", "branding.json")

	// Intentar leer el archivo
	data, err := os.ReadFile(brandingPath)
	if err != nil {
		// Fallback a valores por defecto si no encuentra el archivo
		return getDefaultBranding()
	}

	// Parsear JSON
	var branding Branding
	if err := json.Unmarshal(data, &branding); err != nil {
		fmt.Printf("Error parsing branding.json: %v\n", err)
		return getDefaultBranding()
	}

	return branding
}

// Valores por defecto si no hay branding.json
func getDefaultBranding() Branding {
	return Branding{
		Name:      "XEBEC",
		Version:   "0.1.0",
		Logo:      "XEBEC CORPORATION - CLI",
		Separator: "═══════════════════════════════════════════════════",
		Colors: Colors{
			Primary:       "#6366F1",
			Secondary:     "#0A0A0A",
			White:         "#FFFFFF",
			AccentCyan:    "#00BCD4",
			AccentGreen:   "#4CAF50",
			AccentOrange:  "#FF9800",
			AccentRed:     "#F44336",
			AccentYellow:  "#FFEB3B",
			AccentPurple:  "#8B5CF6",
			GradientStart: "#6366F1",
			GradientEnd:   "#8B5CF6",
			GrayDark:      "#1a1a1a",
			Gray:          "#333333",
			GrayLight:     "#666666",
			GrayLighter:   "#999999",
		},
		Texts: Texts{
			CLILabel:      "CLI",
			PlatformLabel: "Platform",
			MenuTitle:     "Menú Principal",
			FooterNav:     "Presiona ↑/↓ para navegar, Enter para seleccionar, q para salir",
			PromptSel:     "Selecciona una opción:",
			OptionInvalid: "Opción inválida",
			Goodbye:       "¡Hasta luego!",
			Executing:     "Ejecutando:",
		},
		MenuOpts: []MenuOptionBranding{
			{ID: "terminal", Title: "Configurar Terminal", Description: "Configura Alacritty con el tema XEBEC"},
			{ID: "shell", Title: "Configurar Shell", Description: "Configura Nushell + Starship"},
			{ID: "tools", Title: "Instalar Herramientas", Description: "Instala fzf, zoxide, bat, delta, eza"},
			{ID: "status", Title: "Ver Estado", Description: "Muestra el estado de las configuraciones"},
			{ID: "backup", Title: "Crear Backup", Description: "Crea un backup de las configuraciones"},
			{ID: "restore", Title: "Restaurar Backup", Description: "Restaura desde un backup anterior"},
			{ID: "exit", Title: "Salir", Description: "Salir del CLI"},
		},
	}
}

// GetLogo returns the logo from branding
func GetLogo() string {
	return BrandingConfig.Logo
}

// GetVersion returns the version from branding
func GetVersion() string {
	return BrandingConfig.Version
}

// GetSeparator returns the separator from branding
func GetSeparator() string {
	return BrandingConfig.Separator
}

// GetMenuOptions returns the menu options from branding
func GetMenuOptions() []MenuOptionBranding {
	return BrandingConfig.MenuOpts
}

// GetSubmenu returns the submenu options for a given parent ID
func GetSubmenu(parentID string) []MenuOptionBranding {
	if submenus, ok := BrandingConfig.Submenus[parentID]; ok {
		return submenus
	}
	return nil
}

// HasSubmenu checks if a menu option has a submenu
func HasSubmenu(menuID string) bool {
	for _, opt := range BrandingConfig.MenuOpts {
		if opt.ID == menuID && opt.Submenu {
			return true
		}
	}
	return false
}

// GetFooterText returns the appropriate footer text based on current menu level
func GetFooterText(isSubmenu bool) string {
	if isSubmenu {
		if BrandingConfig.Texts.FooterBack != "" {
			return BrandingConfig.Texts.FooterBack
		}
		return "Presiona ← para volver"
	}
	return BrandingConfig.Texts.FooterNav
}
