---
name: skill-sync
description: Skill para sincronizar y recargar skills en OpenCode, coordinar con el orquestador y gestionar agentes y subagentes del proyecto XebecCorporation.Dots.
---

# Skill: skill-sync

Skill para sincronizar y recargar skills en OpenCode, coordinar con el orquestador y gestionar agentes y subagentes del proyecto XebecCorporation.Dots.

## Cuándo Usar Este Skill

Usa este skill cuando:
- Necesitas sincronizar un nuevo skill con OpenCode
- Un skill no es detectado y necesitas diagnosticar el problema
- Necesitas coordinar el orquestador con agentes y subagentes
- Necesitas crear un skill que interactúe con los agentes del proyecto

---

## Arquitectura de Coordinación

### Orquestador XEBEC

El orquestador es el agente principal que coordina la ejecución de flujos completos de configuración. Se comunica con:

```
Orquestador
    ├── Agente: Configuración de Terminal
    ├── Agente: Configuración de Shell
    ├── Agente: Instalación de Herramientas
    └── Subagentes
        ├── Exploración de código
        ├── Análisis de configuraciones
        └── Generación de archivos
```

### Flujo de Comunicación

1. **Orquestador** recibe una tarea compleja
2. **Orquestador** detecta qué agentes necesitan ejecutarse
3. **Agentes** ejecutan tareas específicas
4. **Subagentes** ayudan con exploración y generación
5. **Resultados** se coordinan de vuelta al orquestador

---

## Sincronización de Skills

### Método 1: Nueva sesión (Recomendado)

Cierra y reinicia OpenCode para que cargue los nuevos skills.

### Método 2: Script de sincronización

Usa el script de sincronización:

```bash
python .opencode/skills/skill-sync/scripts/sync_skills.py
```

Este script:
- Valida la estructura de todos los skills
- Verifica frontmatter de cada SKILL.md
- Genera reporte de estado
- Sugiere correcciones si hay errores

### Método 3: Crear skill con integración automática

```bash
python .opencode/skills/skill-sync/scripts/create_skill.py <nombre> --integrate
```

Este comando:
- Crea la estructura del skill
- Registra el skill en el índice de skills
- Prepara la integración con agentes relevantes

---

## Estructura de Skills

```
.opencode/skills/
├── skill-name/
│   ├── SKILL.md           # Obligatorio
│   ├── scripts/           # Opcional
│   ├── references/        # Opcional
│   └── assets/            # Opcional
```

### Frontmatter Requerido

```yaml
---
name: nombre-del-skill
description: Descripción clara de qué hace y cuándo usarlo.
---
```

---

## Integración con Agentes y Subagentes

### Registro de Skills con Agentes

Cada skill puede especificar qué agentes del proyecto puede invocar:

```yaml
---
name: mi-skill
description: Descripción del skill
agents:
  - terminal-config
  - shell-config
  - tools-install
---
```

### Tipos de Agentes

| Tipo | Descripción | Ejemplos |
|------|-------------|----------|
| `terminal-config` | Configuración de terminal (Alacritty) | Theme, fonts, opacity |
| `shell-config` | Configuración de shell (Nushell, Starship) | Prompt, aliases, env |
| `tools-install` | Instalación de herramientas | fzf, zoxide, bat |

### Tipos de Subagentes

| Tipo | Descripción | Uso |
|------|-------------|-----|
| `explore` | Exploración de código | Buscar archivos, entender estructura |
| `analysis` | Análisis de configuraciones | Comparar configs, validar sintaxis |
| `generator` | Generación de archivos | Crear archivos base, templates |

---

## Checklist para Nuevo Skill

1. ✅ Carpeta en `.opencode/skills/<nombre>/`
2. ✅ Archivo `SKILL.md` con frontmatter (name, description)
3. ✅ Nombre del directorio coincide con skill name
4. ✅ Opcional: Especificar `agents` en frontmatter
5. ✅ Opcional: Crear scripts de soporte en `scripts/`
6. ✅ Ejecutar sincronización: `python scripts/sync_skills.py`
7. ✅ Reiniciar contexto de OpenCode

---

## Scripts Disponibles

### sync_skills.py

Sincroniza y valida todos los skills del proyecto.

```bash
python scripts/sync_skills.py [--verbose] [--fix]
```

Opciones:
- `--verbose`: Muestra detalles de cada skill
- `--fix`: Intenta corregir errores automáticamente

### create_skill.py

Crea un nuevo skill con estructura básica.

```bash
python scripts/create_skill.py <nombre> [descripcion] [--integrate]
```

Opciones:
- `--integrate`: Registra el skill para integración con agentes

---

## Integración con el Orquestador

### Cómo el Orquestador Usa los Skills

1. **Detección de tarea**: El orquestador analiza la tarea del usuario
2. **Selección de skill**: Determina qué skills son relevantes
3. **Invocación**: Carga el skill seleccionado
4. **Ejecución**: Aplica las instrucciones del skill
5. **Coordinación**: Coordina múltiples skills si es necesario

### Best Practices para Skills

1. **Skills pequeños y específicos**: Un skill = una tarea específica
2. **Descripciones claras**: El orquestador las usa para seleccionar
3. **Ejemplos concretos**: Ayudan al orquestador a entender el contexto
4. **Integración explícita**: Especifica qué agentes pueden usarse

---

## Validar Skill Creado

### Método 1: Carga manual

```
Cargar skill: <nombre-del-skill>
```

Si aparece el contenido del SKILL.md, está disponible.

### Método 2: Script de validación

```bash
python scripts/sync_skills.py --verbose
```

Busca tu skill en la salida del script.

---

## Tips

- Los skills se cargan desde `.opencode/skills/`
- Evita espacios y caracteres especiales en nombres
- Usa minúsculas para nombres de skills
- Especifica `agents` para integración automática
- Los scripts deben ser ejecutables (chmod +x si es necesario)

---

## Solución de Problemas

| Problema | Solución |
|----------|----------|
| Skill no detectado | Reiniciar OpenCode, luego ejecutar sync_skills.py |
| Error de parseo | Verificar sintaxis YAML del frontmatter |
| Skills vacíos | Verificar estructura de directorios |
| No integra con agentes | Añadir campo `agents` en frontmatter |
| Orquestador no usa skill | Revisar descripción - debe ser específica |

---

## Ejemplos

### Ejemplo 1: Crear skill básico

```bash
python scripts/create_skill.py xebec-terminal "Configuración de terminal XEBEC"
```

### Ejemplo 1: Crear skill con integración

```bash
python scripts/create_skill.py xebec-shell --integrate
```

### Ejemplo 3: Sincronizar todos los skills

```bash
python scripts/sync_skills.py --verbose --fix
```

---

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/skill-sync
