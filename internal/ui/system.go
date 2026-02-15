// Package: ui
// Utilidades de detección de sistema operativo
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/lipgloss"
)

// Información del sistema operativo
type SystemInfo struct {
	OS           string
	Architecture string
	Platform     string
	PackageMgr   string
}

// Detectar información del sistema
func DetectSystem() SystemInfo {
	os := runtime.GOOS
	arch := runtime.GOARCH

	return SystemInfo{
		OS:           os,
		Architecture: arch,
		Platform:     getPlatformName(os),
		PackageMgr:   detectPackageManager(os),
	}
}

// Obtener nombre descriptivo de la plataforma
func getPlatformName(os string) string {
	switch os {
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "darwin":
		return "macOS"
	default:
		return os
	}
}

// Detectar gestor de paquetes
func detectPackageManager(os string) string {
	switch os {
	case "windows":
		return "winget/scoop"
	case "linux":
		// Por defecto asumimos debian/ubuntu
		return "apt"
	default:
		return "N/A"
	}
}

// Renderizar información del sistema
func (s SystemInfo) String() string {
	return fmt.Sprintf("%s %s (%s)", s.Platform, s.Architecture, s.PackageMgr)
}

// Renderizar información del sistema en estilo
func (s SystemInfo) Render() string {
	info := []string{
		fmt.Sprintf("OS: %s", HighlightStyle.Render(s.Platform)),
		fmt.Sprintf("Arch: %s", HighlightStyle.Render(s.Architecture)),
		fmt.Sprintf("PM: %s", HighlightStyle.Render(s.PackageMgr)),
	}
	
	return lipgloss.JoinHorizontal(lipgloss.Top, info...)
}

// Verificar si es Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// Verificar si es Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// Verificar si es macOS
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}
