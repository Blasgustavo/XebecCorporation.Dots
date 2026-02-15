# Output Patterns - Patrones de Salida

Este documento contiene plantillas y ejemplos para formatos de salida consistentes.

## Plantillas Comunes

### 1. Archivo JSON

```json
{
  "skill": "nombre-del-skill",
  "version": "1.0.0",
  "timestamp": "2024-01-01T00:00:00Z",
  "data": {
    // Contenido específico
  }
}
```

### 2. Archivo YAML

```yaml
skill: nombre-del-skill
version: 1.0.0
timestamp: 2024-01-01T00:00:00Z

data:
  # Contenido específico
```

### 3. Markdown con Frontmatter

```markdown
---
title: Título del Documento
date: 2024-01-01
tags: [tag1, tag2]
---

# Contenido
```

## Patrones de Código

### Go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Python

```python
def main():
    print("Hello, World!")

if __name__ == "__main__":
    main()
```

### Shell Script

```bash
#!/bin/bash

echo "Hello, World!"
```

## Formatos de Reporte

### Éxito

```markdown
## Resultado

- **Status**: ✓ Completado
- **Archivos creados**: N
- **Tiempo**: Xms

### Detalles

| Archivo | Estado |
|---------|--------|
| file1   | ✓     |
| file2   | ✓     |
```

### Error

```markdown
## Error

- **Código**: ERR001
- **Mensaje**: Descripción del error
- **Sugerencia**: Cómo resolverlo

### Contexto

```
// Stack trace o información adicional
```
```

## Tablas

### Tabla Simple

```markdown
| Columna 1 | Columna 2 | Columna 3 |
|----------|-----------|-----------|
| Dato 1   | Dato 2    | Dato 3    |
```

### Tabla con Alignment

```markdown
| Left | Center | Right |
|:-----|:------:|------:|
| L    |   C    |     R |
```

## Listas

### Checklist

```markdown
- [ ] Tarea pendiente
- [x] Tarea completada
```

### Definiciones

```markdown
Término 1
: Definición del término 1

Término 2
: Definición del término 2
```
