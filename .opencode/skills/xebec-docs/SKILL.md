---
name: xebec-docs
description: Genera documentación completa para el proyecto XebecCorporation.Dots. Incluye estructura de directorios, configuración de herramientas, guías de uso y referencias técnicas.
---

# Xebec Docs Generator

Este skill proporciona instrucciones detalladas para generar documentación completa del proyecto XebecCorporation.Dots, siguiendo las mejores prácticas de documentación técnica y el estilo del proyecto Gentleman.Dots.

## Cuándo Usar

Usa este skill cuando:
- Necesites crear documentación desde cero
- Agregues nuevas funcionalidades al CLI
- Configures nuevas herramientas
- Actualices dependencias o versiones
- Necesites documentar la estructura del proyecto

## Estructura de Documentación

La documentación del proyecto debe seguir esta estructura:

```
docs/
├── index.md                    # Tabla de contenidos principal
├── getting-started.md          # Guía de inicio
├── installation.md             # Guía de instalación
├── configuration/
│   ├── terminal.md             # Configuración de Alacritty
│   ├── shell.md                # Configuración de Nushell
│   └── tools.md                # Herramientas adicionales
├── development/
│   ├── setup.md                # Configuración de desarrollo
│   ├── commands.md             # Comandos del CLI
│   └── architecture.md         # Arquitectura del proyecto
├── skills/
│   ├── overview.md             # Visión general de skills
│   └── guide.md                # Guía para crear skills
└── contribute.md               # Guía de contribuciones
```

## Reglas de Documentación

### 1. Encabezado de Archivo

Cada documento debe iniciar con:

```markdown
---
title: Título del Documento
description: Breve descripción del contenido (máx 200 caracteres)
---

# Título Principal

[Contenido introductorio...]
```

### 2. Descripción de Herramientas

Para cada herramienta documentada, incluye:

```markdown
## Nombre de la Herramienta

| Propiedad | Valor |
|-----------|-------|
| **Nombre** | nombre-herramienta |
| **Versión** | x.x.x |
| **Tipo** | terminal/shell/prompt/herramienta |
| **Repo** | [enlace](url) |

### Descripción
Breve descripción de qué hace y por qué se usa.

### Instalación

[Comandos de instalación según SO]

### Configuración

[Explicación de opciones de configuración]

### Uso

[Ejemplos de uso]
```

### 3. Descripción de Comandos CLI

Para cada comando del CLI:

```markdown
## Comando Nombre

`xebec nombre-comando [opciones]`

### Descripción
Descripción breve de qué hace el comando.

### Opciones

| Opción | Alias | Descripción | Default |
|--------|-------|-------------|---------|
| `--opcion` | `-o` | Descripción de la opción | valor |

### Ejemplos

```bash
# Ejemplo 1: Descripción
xebec nombre-comando --opcion valor

# Ejemplo 2: Descripción
xebec nombre-comando argumento
```
```

### 4. Estructura de README Principal

El README.md debe contener:

```markdown
# XebecCorporation.Dots

> Descripción corta del proyecto

## Tabla de Contenidos

- [Qué es esto?](#qué-es-esto)
- [Características](#características)
- [Instalación](#instalación)
- [Uso](#uso)
- [Documentación](#documentación)
- [Contribuir](#contribuir)
- [Licencia](#licencia)

## Qué es esto?

[Descripción extendida del proyecto]

## Características

- Feature 1
- Feature 2
- Feature 3

## Instalación

### Requisitos Previos

- Go 1.21+
- [otras dependencias]

### Pasos

[Instrucciones de instalación]

## Uso

[Ejemplos de uso]

## Documentación

Documento | Descripción
----------|-------------
[doc](docs/doc.md) | Descripción

## Contribuir

[Guía de contribuciones]

## Licencia

MIT
```

### 5. Progresión de Descripción

Cuando documentes algo nuevo, sigue este orden:

1. **Qué es**: Definición clara y concisa
2. **Por qué**: Beneficios y razones para usarlo
3. **Cómo**: Instrucciones paso a paso
4. **Ejemplos**: Casos de uso prácticos

### 6. Metadatos de Archivo

Cada archivo de código o configuración debe incluir un comentario de cabecera:

```go
// Package: nombre_del_paquete
// Descripción: Qué hace este paquete
// author: XebecCorporation
// version: 1.0.0
```

## Plantillas

### Template: Nueva Herramienta

```markdown
---
title: Herramienta Nombre
description: Descripción breve
---

# Herramienta Nombre

| Campo | Valor |
|-------|-------|
| Nombre | nombre |
| Versión | x.x.x |
| Tipo | categoria |
| Web | [sitio](url) |

## Por qué usamos esto

[Breve justificación]

## Instalación

### Windows

```powershell
winget install ...
```

### Linux

```bash
# Debian/Ubuntu
sudo apt install ...

# Arch
sudo pacman -S ...
```

## Configuración

[Detalles de configuración]

## Uso

[Ejemplos]
```

### Template: Nuevo Comando

```markdown
---
title: Comando xebec nombre
description: Descripción del comando
---

# Comando `xebec nombre`

## Sinopsis

```
xebec nombre [opciones] [argumentos]
```

## Descripción

[Descripción detallada]

## Opciones

| Opción | Alias | Tipo | Default | Descripción |
|--------|-------|------|---------|-------------|
| `--flag` | `-f` | bool | false | Descripción |

## Ejemplos

```bash
# Básico
xebec nombre

# Con opciones
xebec nombre --flag valor
```
```

## Recursos

- [Gentleman.Dots Reference](https://github.com/Gentleman-Programming/Gentleman.Dots)
- [Documentación de Go](https://go.dev/doc/)
- [Cobra CLI](https://github.com/spf13/cobra)
