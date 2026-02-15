# Workflows - Patrones de Proceso

Este documento contiene patrones para procesos multi-paso y lógica condicional en skills.

## Patrones de Workflow

### 1. Secuencial Simple

Para procesos que siguen una secuencia lineal:

```markdown
## Proceso

1. **Paso 1**: Descripción del primer paso
2. **Paso 2**: Descripción del segundo paso
3. **Paso 3**: Descripción del tercer paso
```

### 2. Con Ramificaciones

Cuando hay decisiones basadas en condiciones:

```markdown
## Flujo

- Si condición A: ejecutar acción 1
- Si condición B: ejecutar acción 2
- Si ninguna: ejecutar acción por defecto
```

### 3. Iterativo

Para procesos que se repiten:

```markdown
## Iteración

Para cada elemento en la lista:
  1. Procesar elemento
  2. Verificar resultado
  3. Si error: continuar o detener según gravedad
```

### 4. Paralelo

Para tareas independientes que pueden ejecutarse simultaneamente:

```markdown
## Ejecución Paralela

- Tarea A: ejecutar de forma independiente
- Tarea B: ejecutar de forma independiente
- Combinar resultados al final
```

## Manejo de Errores

### Retry Pattern

```markdown
## Retry

1. Ejecutar acción
2. Si error y intentos < max:
   - Esperar delay
   - Reintentar (incrementar delay)
3. Si error y intentos >= max:
   - Reportar error
   - Sugerir solución
```

### Fallback Pattern

```markdown
## Fallback

1. Intentar método preferido
2. Si falla:
   - Intentar método alternativo
   - Si también falla:
     - Usar default
     - Reportar situación
```

## Validación

Siempre validar antes de proceder:

```markdown
## Validación

1. Verificar precondiciones
2. Si no se cumplen:
   - Explicar qué falta
   - Proporcionar solución
3. Si se cumplen: continuar
```
