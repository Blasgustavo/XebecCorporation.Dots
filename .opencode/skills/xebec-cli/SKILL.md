---
name: xebec-cli
description: Skill para desarrollar y mantener el CLI de XebecCorporation.Dots, incluyendo comandos, acciones y estructura del proyecto.
---

# Skill: XEBEC CLI Development

Este skill proporciona instrucciones especializadas para el desarrollo del CLI XebecCorporation.Dots.

## Estructura del CLI

```
cmd/xebec/           # Punto de entrada
internal/ui/          # Componentes TUI (menús, banners, estilos)
internal/os/          # Detección de SO, rutas, terminales
internal/actions/     # Acciones de instalación/config
configs/              # Archivos de configuración base
```

## Estructura del Proyecto

| Directorio | Descripción |
|------------|-------------|
| `cmd/xebec/` | Punto de entrada del CLI, comandos Cobra |
| `internal/ui/` | Menús interactivos, banner ASCII, estilos lipgloss |
| `internal/os/` | Detección de SO, terminales instalados |
| `assets/` | Recursos (branding.json, logo) |

## Comandos del CLI

### Menú Principal
- Configurar Terminal - Detecta y configura Alacritty
- Configurar Shell - Configura Nushell + Starship
- Instalar Herramientas - Instala fzf, zoxide, bat, delta, eza
- Ver Estado - Muestra estado de configuraciones
- Crear Backup - Respalda configuraciones
- Restaurar Backup - Restaura desde backup

## Tips de Implementación

1. **Detección de SO**: Usa `runtime.GOOS` para detectar Windows/Linux/macOS
2. **Rutas correctas**: 
   - Windows: `os.Getenv("APPDATA")`
   - Linux: `filepath.Join(os.Getenv("HOME"), ".config")`
3. **Branding**: Edita `assets/branding.json` para personalizar todo
4. **Detección de terminales**: Usa `internal/os/DetectTerminals()`
5. **Backup**: Siempre haz backup antes de sobrescribir archivos

## Detección de Terminales

Usa la función `os.DetectTerminals()` para detectar terminales instalados:

```go
terminals := os.DetectTerminals()
for _, t := range terminals {
    fmt.Printf("%s %s - Instalado: %v\n", t.Icon, t.Name, t.Installed)
}
```

## Branding Modular

El sistema de branding permite editar todo desde `assets/branding.json`:

- Logo ASCII
- Colores hexadecimales
- Textos de la UI
- Opciones del menú con iconos

## Ejemplo: Agregar nuevo comando

```go
// cmd/xebec/commands/terminal.go
package commands

import (
    "github.com/XebecCorporation/XebecCorporation.Dots/internal/os"
    "github.com/spf13/cobra"
)

var terminalCmd = &cobra.Command{
    Use:   "terminal",
    Short: "Configurar terminal",
    Run: func(cmd *cobra.Command, args []string) {
        terminals := os.DetectTerminals()
        // Mostrar terminales detectados
    },
}
```

## Tests

Ejecutar tests:
```bash
go test ./...
```

## Build

Compilar binario:
```bash
go build ./cmd/xebec
```

Ejecutar en modo desarrollo:
```bash
go run ./cmd/xebec
```

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/xebec-cli
