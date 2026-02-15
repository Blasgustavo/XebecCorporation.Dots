// Package: ui
// Banner ASCII Art corporativo para XEBEC CORPORATION
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/lipgloss"
)

// Banner ASCII Art de XEBEC
var BannerASCII = `
██████╗  ██╗███████╗██████╗ ███████╗ ██████╗  ██████╗
██╔══██╗ ██║██╔════╝██╔══██╗██╔════╝██╔═══██╗██╔════╝
███████║ ██║█████╗  ██████╔╝█████╗  ██║   ██║██║     
██╔══██╗ ██║██╔══╝  ██╔══██╗██╔══╝  ██║   ██║██║     
██║  ██║ ██║███████╗██║  ██║███████╗╚██████╔╝╚██████╗
╚═╝  ╚═╝ ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝  ╚═════╝

        XEBEC CORPORATION - CLI - SOPORTE A TERMINAL
`

// Banner con lipgloss (coloreado)
func RenderBanner(version string) string {
	// Estilos para el banner
	bannerStyle := lipgloss.NewStyle().
		Foreground(CorporateBlue).
		Bold(true).
		Align(lipgloss.Center)

	infoStyle := lipgloss.NewStyle().
		Foreground(GrayLighter).
		Align(lipgloss.Center)

	separatorStyle := lipgloss.NewStyle().
		Foreground(Gray).
		Align(lipgloss.Center)

	// Detectar plataforma
	platform := getPlatformInfo()

	return fmt.Sprintf("%s\n%s\n%s",
		bannerStyle.Render(BannerASCII),
		infoStyle.Render(fmt.Sprintf("CLI v%s  |  Platform: %s", version, platform)),
		separatorStyle.Render("═══════════════════════════════════════════════════"),
	)
}

// Obtener información de la plataforma
func getPlatformInfo() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	var osName string
	switch os {
	case "windows":
		osName = "Windows"
	case "linux":
		osName = "Linux"
	case "darwin":
		osName = "macOS"
	default:
		osName = os
	}

	return fmt.Sprintf("%s %s", osName, arch)
}

// Mostrar banner simple en texto
func ShowBanner() {
	fmt.Println(RenderBanner("0.1.0"))
}

// Banner extendido con más información
func RenderExtendedBanner(version string, extraInfo []string) string {
	banner := RenderBanner(version)
	
	if len(extraInfo) > 0 {
		banner += "\n"
		for _, info := range extraInfo {
			banner += "\n" + InfoStyle.Render(info)
		}
	}
	
	return banner
}

// Footer del CLI
func RenderFooter() string {
	footerStyle := lipgloss.NewStyle().
		Foreground(GrayLighter).
		Align(lipgloss.Center)

	return footerStyle.Render("Presiona ↑/↓ para navegar, Enter para seleccionar, q para salir")
}

// Mensaje de carga
func RenderLoading(message string) string {
	loadingStyle := lipgloss.NewStyle().
		Foreground(CorporateBlue).
		Bold(true)

	spinnerStyle := lipgloss.NewStyle().
		Foreground(AccentOrange)

	return spinnerStyle.Render("◐") + " " + loadingStyle.Render(message)
}

// Mensaje de éxito
func RenderSuccess(message string) string {
	return SuccessStyle.Render("✓ ") + NormalTextStyle.Render(message)
}

// Mensaje de error
func RenderError(message string) string {
	return ErrorStyle.Render("✗ ") + NormalTextStyle.Render(message)
}

// Mensaje de información
func RenderInfo(message string) string {
	return HighlightStyle.Render("ℹ ") + NormalTextStyle.Render(message)
}

// Barra de progreso simple
func RenderProgressBar(current, total int, width int) string {
	percentage := float64(current) / float64(total)
	filled := int(float64(width) * percentage)
	
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	
	return NormalTextStyle.Render("[") + HighlightStyle.Render(bar) + NormalTextStyle.Render("]")
}
