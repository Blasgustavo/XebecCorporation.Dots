# Skill: skill-sync

# Sincronización de Skills con OpenCode

Este skill proporciona instrucciones para sincronizar y recargar los skills en OpenCode.

## Problema Común

Al crear un nuevo skill, OpenCode puede no detectarlo inmediatamente debido a caché.

## Solución: Recargar Contexto

### Método 1: Nueva sesión (Recomendado)
Cierra y reinicia OpenCode para que cargue los nuevos skills.

### Método 2: Forzar recarga
En algunos casos, puedes forzar la recarga iniciando con flag fresco.

## Estructura de Skills

```
.opencode/skills/
├── skill-name/
│   └── SKILL.md           # Obligatorio
│   ├── scripts/           # Opcional
│   ├── references/        # Opcional
│   └── assets/           # Opcional
```

## Checklist para Nuevo Skill

1. ✅ Carpeta en `.opencode/skills/<nombre>/`
2. ✅ Archivo `SKILL.md` con frontmatter (name, description)
3. ✅ Nombre del directorio coincide con skill name
4. ✅ Reiniciar contexto de OpenCode

## Validar Skill Creado

Usa el comando de carga de skill:
```
Cargar skill: <nombre-del-skill>
```

Si aparece el contenido del SKILL.md, está disponible.

## Tips

- Los skills se cargan desde `.opencode/skills/`
- El config.json debe tener `"enabled": true` en skills
- Evita espacios y caracteres especiales en nombres de skills
- Usa minúsculas para nombres de skills

## Solución de Problemas

| Problema | Solución |
|----------|----------|
| Skill no detectado | Reiniciar OpenCode |
| Error de parseo | Verificar sintaxis markdown del SKILL.md |
| Skills vacíos | Verificar `config.json` tiene `"enabled": true` |

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/skill-sync
