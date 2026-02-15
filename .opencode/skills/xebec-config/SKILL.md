---
name: xebec-config
description: Skill para gestionar configuraciones de terminal, shell y herramientas del ecosistema XEBEC CORPORATION.
---

# Skill: XEBEC Configuration Management

Este skill proporciona instrucciones para gestionar y aplicar configuraciones del ecosistema XEBEC.

## Configuraciones Disponibles

### Alacritty (configs/alacritty.toml)
- Tema oscuro corporativo (#00AEEF, #0A0A0A)
- JetBrains Mono como fuente principal
- Transparencia 0.85
- Shell por defecto: Nushell

### Nushell (configs/config.nu)
- Banner deshabilitado
- PATH expandido automáticamente
- Integración con Starship
- Editor: nvim

### Starship (configs/starship.toml)
- Prompt minimal con indicadores Git
- Tema corporativo
- Soporte multi-shell

### Zellij (configs/zellij.kdl)
- Layouts predefinidos
- Plugins WebAssembly (futuro)

## Rutas de Configuración

### Windows
- Alacritty: `%APPDATA%/alacritty/alacritty.toml`
- Nushell: `%APPDATA%/nushell/config.nu`
- Starship: `%APPDATA%/starship.toml`

### Linux
- Alacritty: `~/.config/alacritty/alacritty.toml`
- Nushell: `~/.config/nushell/config.nu`
- Starship: `~/.config/starship.toml`

## Proceso de Aplicación

1. **Detectar** - Identificar SO y gestor de paquetes
2. **Verificar** - Comprobar si herramienta está instalada
3. **Backup** - Crear copia de seguridad de config existente
4. **Copiar** - Mover archivo de configs/ a ruta destino
5. **Validar** - Verificar que la configuración es válida
6. **Loguear** - Registrar operación y resultado

## Tips

- Siempre verifica el SO antes de determinar rutas
- Usa templates para generar configs dinámicas
- Mantén backups con timestamp: `config.backup.20240101.toml`
- Proporciona comando para revertir cambios
