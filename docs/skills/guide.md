---
title: Guía de Skills
description: Guía completa para crear y usar skills en XebecCorporation.Dots
---

# Guía de Skills

> Aprende a crear, usar y gestionar skills en XebecCorporation.Dots

## Introducción

Los skills son extensiones especializadas que añaden capacidades específicas al agente OpenCode. Esta guía te enseñará cómo crearlos y usarlos efectivamente.

## Qué es un Skill?

Un skill es un paquete autocontenido que incluye:

- **SKILL.md**: Instrucciones y metadata
- **scripts/**: Código ejecutable opcional
- **references/**: Documentación de referencia
- **assets/**: Archivos estáticos

## Crear un Nuevo Skill

### Paso 1: Define el Propósito

Antes de crear un skill, responde:

- ¿Qué tarea específica resolverá?
- ¿Cuándo debe activarse?
- ¿Qué información necesita?

### Paso 2: Inicializa el Skill

Usa el generador de skills:

```bash
python .opencode/skills/skill-generator/scripts/init_skill.py mi-nuevo-skill --path .opencode/skills/
```

Esto crea:

```
mi-nuevo-skill/
├── SKILL.md
├── scripts/
│   └── example.py
├── references/
│   └── example.md
└── assets/
```

### Paso 3: Edita SKILL.md

#### Metadata (Requerido)

```yaml
---
name: mi-nuevo-skill
description: Descripción clara de qué hace y cuándo usarlo (máx 200 caracteres)
---
```

#### Cuerpo del Documento

```markdown
# Mi Nuevo Skill

[Descripción breve]

## Cuándo Usar

- Cuando el usuario necesite...
- Cuando se pide...

## Instrucciones

1. **Paso 1**: [Descripción]
2. **Paso 2**: [Descripción]

## Ejemplos

### Ejemplo 1
Input: "..."
Output: "..."

## Recursos

- [Referencia](references/archivo.md)
```

### Paso 4: Añade Recursos

#### Scripts

```python
#!/usr/bin/env python3
"""Script del skill."""

def main():
    print("Hello from skill!")

if __name__ == "__main__":
    main()
```

#### Referencias

```markdown
# Documentación de Referencia

## Sección 1

Contenido...

## Sección 2

Contenido...
```

### Paso 5: Empaqueta (Opcional)

Para distribuir el skill:

```bash
python .opencode/skills/skill-generator/scripts/package_skill.py mi-nuevo-skill
```

## Mejores Prácticas

### 1. Enfócate en Una Tarea

Un skill debe hacer una cosa y hacerla bien.

**Bueno**: "Configurar Alacritty"  
**Malo**: "Configurar todo el entorno"

### 2. Descripciones Claras

La descripción es crítica para la activación automática.

**Bueno**:  
> "Aplica la configuración de Alacritty con el tema XEBEC, incluyendo colores, fuentes y atajos de teclado."

**Malo**:  
> "Ayuda con configuración"

### 3. Incluye Ejemplos

Los ejemplos ayudan a entender el comportamiento esperado.

```markdown
## Ejemplos

### Configurar Terminal
Input: "Configura mi terminal con Alacritty"
Output: Copia alacritty.toml a ~/.config/alacritty/
```

### 4. Usa Progressive Disclosure

No pongas todo en SKILL.md. Usa referencias:

```markdown
## Instrucciones

1. [Pasos principales...]

## Detalles Técnicos

Ver [references/config-detailed.md](references/config-detailed.md)
```

### 5. Mantén Actualizado

Los skills deben evolucionar con el proyecto:

- Agrega nuevos ejemplos
- Actualiza referencias
- Refina instrucciones

## Estructura de Archivo de Referencia

### SKILL.md Template

```markdown
---
name: nombre-del-skill
description: Descripción breve (máx 200 caracteres)
---

# Nombre del Skill

[Introducción breve]

## Cuándo Usar

- [Caso 1]
- [Caso 2]

## Instrucciones

1. **Paso 1**: [Descripción]
2. **Paso 2**: [Descripción]

## Ejemplos

### Ejemplo 1
[Descripción del ejemplo]

## Errores Comunes

- [Error 1]: [Solución]
- [Error 2]: [Solución]

## Recursos

- [Enlace 1]
- [Enlace 2]
```

## Testing de Skills

### Prueba Manual

1. Activa el skill explícitamente
2. Ejecuta tareas de prueba
3. Verifica el comportamiento

### Verificación de Estructura

```bash
# Verificar estructura
ls -la mi-skill/

# Verificar SKILL.md
cat mi-skill/SKILL.md | head -20
```

## Integración con OpenCode

### Configuración

Edita `.opencode/config.json`:

```json
{
  "skills": {
    "enabled": ["xebec-cli", "xebec-config"],
    "autoActivate": true
  }
}
```

### Activación

```
Usa el skill xebec-cli para crear un comando de instalación
```

## Ejemplo: Crear un Skill para Git

### SKILL.md

```markdown
---
name: xebec-git
description: Configura Git con las mejores prácticas de XEBEC CORPORATION
---

# Xebec Git

Configura Git con alias, hooks y configuración profesional.

## Cuándo Usar

- "Configura mi Git"
- "Mejora mi configuración de Git"

## Instrucciones

1. **Detectar SO**: Windows o Linux
2. **Localizar config**: `~/.gitconfig`
3. **Aplicar configuración**: Alias, hooks, preferencias

## Alias Include

```ini
[alias]
st = status
co = checkout
br = branch
ci = commit
lg = log --graph --oneline --decorate
```

## Configuración Recomendada

```ini
[user]
name = Tu Nombre
email = tu@email.com

[pull]
rebase = false

[init]
defaultBranch = main
```

## Recursos

- [Git Documentation](https://git-scm.com/doc)
- [Oh My Git](https://ohmygit.org)
```

---

## Referencias

- [Agent Skills Specification](https://agentskills.io)
- [Anthropic Skills Repository](https://github.com/anthropics/skills)
- [Guía de Documentación](xebec-docs.md)

---

*Consulta también: [Visión General de Skills](overview.md)*
