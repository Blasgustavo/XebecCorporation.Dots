---
title: Instalación
description: Guía completa de instalación de XebecCorporation.Dots y sus dependencias
---

# Instalación

> Guía detallada para instalar XebecCorporation.Dots y todas sus dependencias

## Requisitos del Sistema

| Sistema Operativo | Arquitectura | Requisitos |
|-------------------|-------------|------------|
| Windows 11 | x86_64 | Go 1.21+, PowerShell/CMD |
| Linux (Ubuntu/Debian) | x86_64, ARM64 | Go 1.21+, Bash |
| Linux (Arch) | x86_64 | Go 1.21+, Bash |
| Linux (Fedora) | x86_64 | Go 1.21+, Bash |

## Instalación de Go

### Windows

#### Opción 1: Winget (Recomendado)

```powershell
winget install GoLang.Go --accept-source-agreements --accept-package-agreements
```

#### Opción 2: Chocolatey

```powershell
choco install golang
```

#### Opción 3: Instalador Manual

1. Descarga el instalador desde [go.dev/dl](https://go.dev/dl/)
2. Ejecuta el instalador `.msi`
3. Sigue las instrucciones del asistente

### Linux

#### Debian/Ubuntu

```bash
# Descargar
wget https://go.dev/dl/go1.26.0.linux-amd64.tar.gz

# Instalar
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz

# Agregar al PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### Arch Linux

```bash
sudo pacman -S go
```

#### Fedora

```bash
sudo dnf install golang
```

### Verificar Instalación de Go

```bash
go version
# Output: go version go1.26.0 linux/amd64
```

## Instalación de XebecCorporation.Dots

### Clonar el Repositorio

```bash
git clone https://github.com/XebecCorporation/XebecCorporation.Dots.git
cd XebecCorporation.Dots
```

### Compilar el CLI

```bash
# Compilar binario
go build -o xebec.exe ./cmd/xebec

# O usar go install
go install ./cmd/xebec
```

### Agregar al PATH (Opcional)

#### Windows

Copia `xebec.exe` a una carpeta en tu PATH o agrega el directorio:

```powershell
# Agregar al PATH de forma temporal
$env:PATH += ";C:\ruta\a\XebecCorporation.Dots"

# Agregar permanentemente
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\ruta\a\XebecCorporation.Dots", "User")
```

#### Linux

```bash
sudo cp xebec /usr/local/bin/
# O agrega a tu PATH
echo 'export PATH=$PATH:/ruta/a/XebecCorporation.Dots' >> ~/.bashrc
```

## Dependencias Opcionales

### Alacritty (Terminal)

#### Windows

```powershell
winget install Alacritty.Alacritty
```

#### Linux

```bash
# Ubuntu/Debian
sudo apt install alacritty

# Arch
sudo pacman -S alacritty

# Fedora
sudo dnf install alacritty
```

### Nushell (Shell)

#### Windows

```powershell
winget install Nushell.Nushell
```

#### Linux

```bash
# Ubuntu/Debian
sudo apt install nushell

# Arch
sudo pacman -S nushell

# Fedora
sudo dnf install nushell
```

### Starship (Prompt)

#### Todos los Sistemas

```bash
# Bash, Zsh, Fish
curl -sS https://starship.rs/install.sh | sh
```

## Verificación de Instalación

```bash
# Ver versión
xebec version

# Ver ayuda
xebec --help

# Probar comandos
xebec config
xebec install
```

## Actualización

```bash
# Actualizar código
git pull origin main

# Recompilar
go build -o xebec.exe ./cmd/xebec
```

## Desinstalación

### Eliminar CLI

```bash
# Linux
sudo rm /usr/local/bin/xebec

# Windows
Remove-Item C:\ruta\a\xebec.exe
```

### Eliminar Configuraciones

```bash
# Eliminar configuraciones de usuario
rm -rf ~/.config/xebec
rm -rf ~/.xebec
```

---

*Para más información, consulta [Arquitectura del Proyecto](development/architecture.md).*
