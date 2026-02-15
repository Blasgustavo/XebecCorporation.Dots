---
title: Configuración de Terminal
description: Guía completa para configurar Alacritty con el tema XEBEC
---

# Configuración de Terminal

> Configura Alacritty con el tema XEBEC CORPORATION

## Alacritty

| Campo | Valor |
|-------|-------|
| **Nombre** | Alacritty |
| **Versión** | 0.14.0+ |
| **Tipo** | Terminal Emulator |
| **Repo** | [alacritty/alacritty](https://github.com/alacritty/alacritty) |
| **Web** | [alacritty.org](https://alacritty.org) |

## Por qué Alacritty?

Alacritty es un terminal moderno con las siguientes ventajas:

- **Rendimiento**: Rendering GPU acelerado
- **Ligero**: Minimum dependencies, máximo rendimiento
- **Cross-platform**: Windows, Linux, macOS
- **Configuración simple**: Archivo TOML intuitivo
- **Comunidad activa**: Desarrollo constante

## Instalación

### Windows

```powershell
winget install Alacritty.Alacritty
```

### Linux

```bash
# Ubuntu/Debian
sudo apt install alacritty

# Arch
sudo pacman -S alacritty

# Fedora
sudo dnf install alacritty
```

## Configuración

### Ubicación de Configuración

| Sistema | Ruta |
|---------|------|
| Windows | `%APPDATA%\alacritty\alacritty.toml` |
| Linux | `~/.config/alacritty/alacritty.toml` |

### Configuración Base

El proyecto incluye una configuración base en `alacritty/alacritty.toml`:

```toml
[window]
opacity = 0.95
padding = { x = 10, y = 10 }

[window.dimensions]
columns = 120
lines = 30

[font]
size = 12.0

[font.normal]
family = "JetBrains Mono"
style = "Regular"

[colors]
# Tema XEBEC
[colors.primary]
background = '#1a1b26'
foreground = '#a9b1d6'

[colors.cursor]
text = '#1a1b26'
cursor = '#c0caf5'

[colors.vi_mode_cursor]
text = '#1a1b26'
cursor = '#bb9af7'

[colors.search.matches]
foreground = '#1a1b26'
background = '#7aa2f7'

[colors.search.focused_match]
foreground = '#1a1b26'
background = '#9ece6a'

[colors.hints.start]
foreground = '#1a1b26'
background = '#e0af68'

[colors.hints.end]
foreground = '#1a1b26'
background = '#f7768e'

[colors.selection]
text = '#1a1b26'
background = '#c0caf5'

[colors.normal]
black = '#15161e'
red = '#f7768e'
green = '#9ece6a'
yellow = '#e0af68'
blue = '#7aa2f7'
magenta = '#bb9af7'
cyan = '#7dcfff'
white = '#a9b1d6'

[colors.bright]
black = '#414868'
red = '#f7768e'
green = '#9ece6a'
yellow = '#e0af68'
blue = '#7aa2f7'
magenta = '#bb9af7'
cyan = '#7dcfff'
white = '#c0caf5'

[colors.dim]
black = '#15161e'
red = '#f7768e'
green = '#9ece6a'
yellow = '#e0af68'
blue = '#7aa2f7'
magenta = '#bb9af7'
cyan = '#7dcfff'
white = '#a9b1d6'

[env]
TERM = "tmux-256color"
```

## Aplicar Configuración con CLI

```bash
# Copiar configuración base
xebec config terminal
```

Este comando:
1. Detecta el sistema operativo
2. Localiza el directorio de configuración
3. Copia `alacritty.toml` a la ubicación correcta
4. Aplica el tema XEBEC

## Personalización

### Cambiar Fuente

```toml
[font]
size = 14.0

[font.normal]
family = "FiraCode Nerd Font"
style = "Regular"
```

### Cambiar Tamaño de Ventana

```toml
[window.dimensions]
columns = 100
lines = 40
```

### Habilitar Transparacia

```toml
[window]
opacity = 0.90
```

## Atajos de Teclado

| Atajo | Acción |
|-------|--------|
| `Ctrl+Shift+T` | Nueva pestaña |
| `Ctrl+Shift+W` | Cerrar pestaña |
| `Ctrl+Shift+L` | Dividir horizontal |
| `Ctrl+Shift+V` | Dividir vertical |
| `Ctrl+Shift+Arrow` | Moverse entre panes |

## Solución de Problemas

### Error: "Font not found"

Instala la fuente especificada o cambia la familia de fuentes:

```toml
[font.normal]
family = "Courier New"
```

### Pantalla en blanco

Verifica que tu terminal soporte rendering acelerado:

```toml
[renderer]
enable_ligatures = false
```

---

*Consulta también: [Configuración de Shell](shell.md)*
