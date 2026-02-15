---
title: Visión General de Skills
description: Introducción a los skills de XebecCorporation.Dots
---

# Visión General de Skills

> Introducción a los skills disponibles en el proyecto

## Qué son los Skills?

Los skills son conjuntos de instrucciones especializadas que potencian las capacidades del agente OpenCode para tareas específicas del proyecto XebecCorporation.Dots.

Siguiendo el estándar [Agent Skills](https://agentskills.io) de Anthropic, cada skill proporciona:

- Instrucciones detalladas para tareas específicas
- Plantillas y ejemplos
- Referencias a recursos adicionales

## Skills Disponibles

### Skills Principales

| Skill | Descripción |
|-------|-------------|
| [xebec-cli](xebec-cli.md) | Desarrollo del CLI en Go |
| [xebec-config](xebec-config.md) | Configuración de herramientas |
| [skill-generator](skill-generator.md) | Creación de nuevos skills |
| [xebec-docs](xebec-docs.md) | Generación de documentación |

### Skills de Utility

| Skill | Descripción |
|-------|-------------|
| skill-sync | Sincronización de skills |
| skill-commits | Conventional commits |

## Estructura de un Skill

```
.opencode/
├── skills/
│   ├── skill-name/
│   │   ├── SKILL.md           # Instrucciones principales
│   │   ├── scripts/           # Scripts ejecutables
│   │   ├── references/        # Documentación de referencia
│   │   └── assets/            # Archivos estáticos
│   └── ...
├── agent/                     # Definición de agentes
└── config.json               # Configuración de OpenCode
```

## Cómo Usar los Skills

### Activación Automática

OpenCode activa automáticamente los skills relevantes según el contexto de la conversación.

### Activación Manual

Puedes invocar un skill explícitamente:

```
Usa el skill xebec-cli para crear un nuevo comando
```

## Progressive Disclosure

Los skills usan un sistema de revelación progresiva:

1. **Metadata** (name, description): OpenCode decide cuándo invocar
2. **SKILL.md body**: Instrucciones detalladas
3. **Recursos**: Scripts y referencias adicionales

---

## Skills del Proyecto

### xebec-cli

Desarrollo del CLI XebecCorporation.Dots:

- Comandos con Cobra
- Estructura del proyecto
- Mejores prácticas de Go

### xebec-config

Configuración de herramientas:

- Alacritty
- Nushell
- Starship
- Herramientas adicionales

### skill-generator

Creación de nuevos skills:

- Estructura según Anthropic
- Plantillas
- Empaquetado

### xebec-docs

Generación de documentación:

- Estructura de docs
- Plantillas
- Estilos de escritura

---

## Crear un Nuevo Skill

Para crear un nuevo skill:

```bash
python .opencode/skills/skill-generator/scripts/init_skill.py <nombre> --path .opencode/skills/
```

Luego edita el `SKILL.md` generado con las instrucciones específicas.

---

## Referencias

- [Agent Skills Specification](https://agentskills.io)
- [Anthropic Skills](https://github.com/anthropics/skills)
- [Documentación de OpenCode](https://opencode.ai)

---

*Consulta también: [Guía de Skills](guide.md)*
