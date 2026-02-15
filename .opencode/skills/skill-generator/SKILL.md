---
name: skill-generator
description: Guía completa para crear, desarrollar y empaquetar nuevos skills para OpenCode. Usa este skill cuando necesites crear un nuevo skill desde cero, actualizar uno existente, o entender las mejores prácticas para diseñar skills efectivos con progressive disclosure.
license: Apache-2.0
---

# Skill Generator

Este skill proporciona una guía completa para crear skills efectivos para OpenCode.

## Proceso de Creación

### Paso 1: Entender el Skill

Antes de crear un skill, comprende ejemplos concretos de uso:

- ¿Qué funcionalidad debe soportar el skill?
- ¿Cuáles son los casos de uso más comunes?
- ¿Qué dirá el usuario para activar este skill?

**No ask** demasiadas preguntas de una vez. Comienza con las más importantes y sigue según sea necesario.

### Paso 2: Planificar Contenidos Reutilizables

Analiza cada ejemplo para identificar:

- **Scripts**: Código determinista que se ejecuta repetidamente
- **References**: Documentación que Claude debe consultar
- **Assets**: Archivos que se usan en el output (plantillas, imágenes)

### Paso 3: Inicializar el Skill

Usa el script de inicialización:

```bash
python scripts/init_skill.py <nombre-del-skill> --path <ruta-salida>
```

Esto crea la estructura completa automáticamente.

### Paso 4: Editar el Skill

Personaliza el SKILL.md generado:

- **Frontmatter**: name y description claros
- **Body**: Instrucciones concisas, menos de 500 líneas
- **Recursos**: Scripts, references, assets según necesidad

### Paso 5: Empaquetar

Valida y empaqueta el skill:

```bash
python scripts/package_skill.py <ruta-skill> [directorio-salida]
```

### Paso 6: Iterar

Mejora el skill basándote en el uso real.

## Estructura de un Skill

```
skill-name/
├── SKILL.md           # Requerido
├── scripts/           # Opcional: código ejecutable
├── references/        # Opcional: documentación
└── assets/           # Opcional: archivos de output
```

## Principios Core

1. **Concisión**: Solo incluye lo que Claude necesita y no sabe
2. **Grados de libertad apropiados**: 
   - Alto: instrucciones textuales
   - Medio: pseudocódigo/parámetros
   - Bajo: scripts específicos
3. **Progressive disclosure**: Metadata → SKILL.md → Recursos

## Referencias

- **Workflows**: Ver [references/workflows.md](references/workflows.md)
- **Patrones de output**: Ver [references/output-patterns.md](references/output-patterns.md)
