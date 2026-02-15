// Package: commands
// Comandos principales del CLI XEBEC
// author: XebecCorporation
// version: 1.0.0

package commands

import (
	"fmt"
	"os"

	"github.com/XebecCorporation/XebecCorporation.Dots/internal/ui"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "xebec",
	Short: "XEBEC CORPORATION CLI - Configura y gestiona tu entorno de desarrollo",
	Long: `XebecCorporation.Dots es un CLI interactivo para instalar, 
configurar y mantener el ecosistema XEBEC CORPORATION.

Ejemplos de uso:
  xebec              - Inicia el menú interactivo
  xebec config       - Configura componentes
  xebec install      - Instala herramientas
  xebec version      - Muestra la versión`,
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Si no hay argumentos, ejecutar menú interactivo
		if len(args) == 0 {
			return runInteractiveMenu()
		}
		return nil
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(interactiveCmd)

	// Configure root command
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}

// runInteractiveMenu ejecuta el menú interactivo
func runInteractiveMenu() error {
	// Mostrar banner
	ui.ShowBanner()
	fmt.Println()

	// Ejecutar menú interactivo
	return ui.RunMenu(version)
}

// configCmd handles configuration of tools
var configCmd = &cobra.Command{
	Use:   "config [terminal|shell]",
	Short: "Configura componentes del ecosistema XEBEC",
	Long:  `Aplica configuraciones XEBEC a terminal, shell y herramientas.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(ui.RenderInfo("Usa uno de los subcomandos: terminal, shell"))
			fmt.Println()
			fmt.Println("  xebec config terminal  - Configura Alacritty")
			fmt.Println("  xebec config shell    - Configura Nushell + Starship")
			return
		}

		switch args[0] {
		case "terminal":
			fmt.Println(ui.RenderInfo("Configurando Alacritty..."))
			// TODO: Implementar configuración de terminal
		case "shell":
			fmt.Println(ui.RenderInfo("Configurando Nushell + Starship..."))
			// TODO: Implementar configuración de shell
		default:
			fmt.Println(ui.RenderError(fmt.Sprintf("Componente desconocido: %s", args[0])))
		}
	},
}

// installCmd handles installation of tools
var installCmd = &cobra.Command{
	Use:   "install [tools]",
	Short: "Instala herramientas del ecosistema XEBEC",
	Long:  `Instala herramientas según el sistema operativo detectado.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(ui.RenderInfo("Usa: xebec install tools"))
			return
		}

		switch args[0] {
		case "tools":
			fmt.Println(ui.RenderInfo("Instalando herramientas..."))
			// Detectar sistema
			sysInfo := ui.DetectSystem()
			fmt.Println(ui.NormalTextStyle.Render("Sistema detectado: " + sysInfo.String()))
			// TODO: Implementar instalación de herramientas
		default:
			fmt.Println(ui.RenderError(fmt.Sprintf("Comando desconocido: %s", args[0])))
		}
	},
}

// versionCmd shows version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Muestra la versión del CLI",
	Long:  `Muestra información de versión de XebecCorporation.Dots`,
	Run: func(cmd *cobra.Command, args []string) {
		// Mostrar banner con versión
		fmt.Println(ui.RenderBanner(version))
		fmt.Println()
		
		// Información del sistema
		sysInfo := ui.DetectSystem()
		fmt.Println(ui.InfoStyle.Render("Sistema: " + sysInfo.String()))
		fmt.Println()
		
		// Links
		fmt.Println(ui.MutedTextStyle.Render("Más información:"))
		fmt.Println(ui.MutedTextStyle.Render("  GitHub: https://github.com/XebecCorporation/XebecCorporation.Dots"))
	},
}

// interactiveCmd runs the interactive menu
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Inicia el menú interactivo",
	Long:  `Abre el menú interactivo del CLI`,
	Aliases: []string{"i", "menu"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInteractiveMenu(); err != nil {
			fmt.Println(ui.RenderError(fmt.Sprintf("Error: %v", err)))
			os.Exit(1)
		}
	},
}
