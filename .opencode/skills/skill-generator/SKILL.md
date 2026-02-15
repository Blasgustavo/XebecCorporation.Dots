---
name: skill-generator
description: Guía completa para crear, desarrollar y empaquetar nuevos skills para OpenCode siguiendo el estándar Agent Skills de Anthropic. Usa este skill cuando necesites crear un nuevo skill desde cero, actualizar uno existente, o entender las mejores prácticas para diseñar skills efectivos.
license: Apache-2.0
---

# Skill Generator

Este skill proporciona una guía completa para crear skills efectivos para OpenCode siguiendo el estándar [Agent Skills](https://agentskills.io) de Anthropic.

## Cuándo Usar Este Skill

Usa este skill cuando:
- Necesitas crear un nuevo skill desde cero
- Quieres actualizar un skill existente
- Necesitas entender las mejores prácticas para diseñar skills
- Quieres seguir el estándar de Anthropic para skills

---

## Estructura de un Skill

Según el estándar de Anthropic, un skill tiene la siguiente estructura:

```
mi-skill/
├── SKILL.md           # REQUERIDO - Instrucciones y metadata
├── scripts/           # OPCIONAL - Código ejecutable
├── references/        # OPCIONAL - Documentación de referencia
└── assets/            # OPCIONAL - Archivos estáticos
```

### Progressive Disclosure

El sistema usa tres niveles de información:

1. **Metadata** (name, description): Primera capa - Claude decide cuándo invocar el skill
2. **SKILL.md body**: Segunda capa - Instrucciones detalladas
3. **Recursos** (scripts, references): Tercera capa - Código y documentación adicional

---

## Proceso de Creación

### Paso 1: Define el Propósito

Antes de crear un skill, responde:

- **¿Qué funcionalidad debe soportar?** Tarea específica y repetible
- **¿Cuándo debe usarse?** Descripción clara (máx 200 caracteres)
- **¿Qué dirá el usuario para activarlo?** Palabras clave y ejemplos

**Principio clave**: Un skill = una tarea específica. Skills pequeños se componen mejor.

### Paso 2: Crea la Estructura

Usa el script de inicialización:

```bash
python .opencode/skills/skill-generator/scripts/init_skill.py <nombre-del-skill> --path .opencode/skills/
```

Esto crea:
```
mi-skill/
├── SKILL.md
├── scripts/
│   └── example.py
├── references/
│   └── example.md
└── assets/
```

### Paso 3: Edita el SKILL.md

#### Frontmatter (Requerido)

```yaml
---
name: mi-skill-name
description: Descripción clara de qué hace y cuándo usarlo (máx 200 caracteres).
---
```

#### Frontmatter (Opcional)

```yaml
---
name: mi-skill
description: Descripción breve
dependencies: python>=3.8, pandas>=1.5.0
---
```

#### Body del SKILL.md

Estructura recomendada:

```markdown
# Nombre del Skill

[Descripción breve de qué hace el skill]

## Cuándo Usar

Describe específicamente cuándo Claude debe activar este skill.

## Instrucciones

Pasos detallados que Claude debe seguir...

## Ejemplos

### Ejemplo 1
Input: ...
Output: ...

### Ejemplo 2
Input: ...
Output: ...

## Recursos

- [Referencia 1](references/archivo.md)
- [Referencia 2](scripts/script.py)
```

### Paso 4: Empaqueta el Skill

1. Asegúrate que el nombre de la carpeta coincida con el `name` en el frontmatter
2. Crea un ZIP con la carpeta como raíz (no pongas archivos sueltos en la raíz del ZIP)
3. El ZIP debe contener la carpeta del skill como elemento raíz

**Estructura correcta:**
```
mi-skill.zip
└── mi-skill/
    ├── SKILL.md
    └── references/
```

### Paso 5: Prueba el Skill

1. Revisa que SKILL.md sea claro
2. Verifica que la descripción refleje cuándo usarlo
3. Confirma que los archivos referenciados existan
4. Prueba con prompts de ejemplo

---

## Mejores Prácticas

### 1. Enfócate en Una Tarea

- Un skill = una responsabilidad
- Skills pequeños se composición mejor que uno grande
- Claude puede usar múltiples skills juntos automáticamente

### 2. Escribe Descripciones Claras

La descripción es crítica - Claude la usa para decidir cuándo invocar el skill.

**Bueno:**
> "Aplica las guías de marca de Acme Corp a presentaciones y documentos, incluyendo colores oficiales, fuentes y uso del logo."

**Malo:**
> "Ayuda con documentos."

### 3. Comienza Simple

- Primero crea instrucciones básicas en Markdown
- Añade scripts solo si es necesario
- Itera y expande después

### 4. Incluye Ejemplos

```markdown
## Ejemplos

### Crear un archivo
Input: "Crea un archivo README.md para mi proyecto"
Output: Archivo con estructura básica de README

### Actualizar configuración
Input: "Actualiza el theme a oscuro"
Output: config.toml modificado con theme=dark
```

### 5. Usa Grados de Libertad Apropiados

| Nivel | Descripción | Cuándo Usar |
|-------|-------------|-------------|
| Alto | Instrucciones textuales genéricas | Tareas creativas |
| Medio | Pseudocódigo con parámetros | Tareas convariaciones |
| Bajo | Scripts específicos deterministas | Tareas repetibles |

---

## Errores Comunes a Evitar

1. **Skill muy grande**: Separa en skills más pequeños
2. **Descripción vaga**: Sé específico sobre cuándo usar
3. **Sin ejemplos**: Los ejemplos ayudan a Claude a entender el output esperado
4. **Hardcodear secrets**: No pongas API keys o passwords
5. **Skip testing**: Prueba después de cada cambio significativo

---

## Scripts Disponibles

### init_skill.py

Inicializa un nuevo skill con estructura básica.

```bash
python scripts/init_skill.py <nombre> --path <ruta>
```

### package_skill.py

Valida y empaqueta un skill.

```bash
python scripts/package_skill.py <ruta-skill> [directorio-salida]
```

---

## Referencias

- [Agent Skills Specification](https://agentskills.io)
- [Anthropic Skills Repository](https://github.com/anthropics/skills)
- [Documentación de Anthropic](https://support.claude.com/en/articles/12512198-creating-custom-skills)
