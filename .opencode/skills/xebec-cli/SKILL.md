---
name: xebec-cli
description: Skill para desarrollar y mantener el CLI de XebecCorporation.Dots, incluyendo comandos, acciones y estructura del proyecto.
---

# Skill: XEBEC CLI Development

Este skill proporciona instrucciones especializadas para el desarrollo del CLI XebecCorporation.Dots.

## Estructura del CLI

```
cmd/xebec/           # Punto de entrada
internal/ui/         # Componentes TUI
internal/os/         # Detección de SO y rutas
internal/actions/    # Acciones de instalación/config
configs/             # Archivos de configuración base
```

## Comandos del CLI

### install
Instala herramientas y dependencias según el SO detectado.
- Windows: usa winget/scoop
- Linux: usa apt/pacman/dnf

### config
Aplica configuraciones XEBEC a:
- Alacritty (terminal)
- Nushell (shell)
- Starship (prompt)

### configure
Configura componentes específicos:
- `configure terminal` - Configura Alacritty
- `configure shell` - Configura Nushell + Starship
- `configure tools` - Instala herramientas adicionales

## Implementación de Comandos

Usa Cobra para crear comandos:

```go
var installCmd = &cobra.Command{
    Use:   "install",
    Short: "Instala herramientas XEBEC",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Tu código aquí
        return nil
    },
}
```

## Tips de Implementación

1. Siempre detecta el SO primero con `internal/os/detect.go`
2. Usa rutas correctas: `os.UserConfigDir()`, `os.Getenv("APPDATA")`
3. Realiza backup antes de sobrescribir
4. Registra operaciones en log
5. Proporciona feedback claro al usuario
