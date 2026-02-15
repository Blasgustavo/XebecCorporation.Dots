---
title: Inicio Rápido
description: Guía de inicio rápido para comenzar con XebecCorporation.Dots
---

# Inicio Rápido

> Comienza a usar XebecCorporation.Dots en minutos

## Qué es XebecCorporation.Dots?

XebecCorporation.Dots es un CLI interactivo, moderno y multiplataforma (Windows + Linux) escrito en Go. Su propósito es instalar, configurar y mantener el ecosistema XEBEC CORPORATION en cualquier máquina mediante un binario único con menús TUI.

## Requisitos Previos

| Requisito | Versión Mínima | Descripción |
|------------|----------------|-------------|
| Go | 1.21+ | Compilador de Go |
| Git | 2.0+ | Control de versiones |
| Windows 11 / Linux | - | Sistema operativo |

### Dependencias Adicionales (Opcionales)

| Herramienta | Propósito | Instalador |
|-------------|-----------|-------------|
| Alacritty | Terminal | Windows: winget / Linux: pacman/apt |
| Nushell | Shell moderno | winget / pacman / apt |
| Starship | Prompt | Instalador automatico |

## Instalación Rápida

### 1. Clonar el Repositorio

```bash
git clone https://github.com/XebecCorporation/XebecCorporation.Dots.git
cd XebecCorporation.Dots
```

### 2. Compilar el CLI

```bash
# Usando Go
go build -o xebec.exe ./cmd/xebec

# O ejecutar en modo desarrollo
go run ./cmd/xebec
```

### 3. Verificar Instalación

```bash
./xebec.exe --help
```

Deberías ver:

```
XebecCorporation.Dots es un CLI interactivo para instalar, 
configurar y mantener el ecosistema XEBEC CORPORATION.

Usage:
  xebec [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Configura componentes del ecosistema XEBEC
  help        Help about any command
  install     Instala herramientas del ecosistema XEBEC
  version     Muestra la versión del CLI

Flags:
  -h, --help      help for xebec
  -v, --version   version for xebec
```

## Uso Básico

### Ver Versión

```bash
xebec version
# Output: XebecCorporation.Dots v0.1.0
```

### Configurar Terminal

```bash
xebec config terminal
```

### Configurar Shell

```bash
xebec config shell
```

### Instalar Herramientas

```bash
xebec install tools
```

## Próximos Pasos

1. **[Instalación Detallada](installation.md)** - Guía completa de instalación
2. **[Configuración de Terminal](configuration/terminal.md)** - Configura Alacritty
3. **[Configuración de Shell](configuration/shell.md)** - Configura Nushell + Starship
4. **[Arquitectura](development/architecture.md)** - Entiende el código

## Solución de Problemas

### Error: "go: command not found"

Asegúrate de tener Go instalado y en tu PATH. Verifica con:

```bash
go version
```

### Error: "No se encontró el comando"

Asegúrate de que el binario `xebec.exe` esté en tu PATH o usa la ruta completa.

---

*Para más ayuda, consulta la [guía de contribuciones](contribute.md) o abre un issue en GitHub.*
