---
title: Guía de Contribuciones
description: Cómo contribuir al proyecto XebecCorporation.Dots
---

# Guía de Contribuciones

> ¡Gracias por tu interés en contribuir a XebecCorporation.Dots!

## Código de Conducta

Somos una comunidad abierta y respetuosa. Al participar, te comprometes a:

- Ser respetuoso con otros contribuyentes
- Aceptar críticas constructivas de manera positiva
- Enfocarse en lo que es mejor para la comunidad

## Cómo Contribuir

### 1. Reportar Bugs

Usa GitHub Issues para reportar bugs:

1. Busca issues existentes
2. Crea un nuevo issue con:
   - Título claro
   - Pasos para reproducir
   - Expected vs actual behavior
   - Screenshots si aplica

### 2. Sugerir Features

1. Busca issues relacionados
2. Crea un issue con:
   - Descripción clara del feature
   - Casos de uso
   - Alternativas consideradas

### 3. Pull Requests

#### Proceso

1. **Fork** el repositorio
2. Crea una **rama** feature: `git checkout -b feature/mi-feature`
3. Haz **commits** siguiendo Conventional Commits
4. **Push** a tu fork
5. Crea un **Pull Request**

#### Conventional Commits

```
<tipo>(<alcance>): <descripción>

[ cuerpo opcional ]

[ pies opcionales ]
```

**Tipos**:

| Tipo | Descripción |
|------|-------------|
| `feat` | Nueva característica |
| `fix` | Bug fix |
| `docs` | Documentación |
| `style` | Formateo |
| `refactor` | Refactorización |
| `test` | Tests |
| `chore` | Mantenimiento |

**Ejemplos**:

```bash
feat(cli): add install tools command
fix(config): resolve path detection on Windows
docs(readme): update installation instructions
refactor(actions): simplify terminal config flow
```

### 4. Revisión de Código

- Sé constructivo y respetuoso
- Enfócate en el código, no en la persona
- Sugiere soluciones, no solo problemas

## Requisitos de Desarrollo

### Entorno

| Requisito | Versión |
|------------|---------|
| Go | 1.21+ |
| Git | 2.0+ |
| Editor | VS Code / GoLand |

### Setup de Desarrollo

```bash
# Clonar
git clone https://github.com/XebecCorporation/XebecCorporation.Dots.git
cd XebecCorporation.Dots

# Instalar dependencias
go mod download

# Compilar
go build -o xebec.exe ./cmd/xebec

# Ejecutar tests
go test ./...

# Ejecutar linter
golangci-lint run
```

## Estándares de Código

### Go

- Usa `gofmt` para formateo
- Sigue [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Escribe tests unitarios

### Documentación

- Comenta funciones exportadas
- Mantén README.md actualizado
- Documenta cambios significativos

### Commits

- Usa mensajes claros y descriptivos
- Un cambio por commit
- Incluye contexto en el cuerpo

## Estructura de Proyecto

```
XebecCorporation.Dots/
├── cmd/xebec/           # Punto de entrada
├── internal/            # Paquetes internos
│   ├── actions/        # Lógica de acciones
│   ├── os/             # Detección de SO
│   └── ui/             # Componentes UI
├── configs/            # Archivos de configuración
├── docs/               # Documentación
└── .opencode/         # Configuración de OpenCode
```

## Recursos

### Documentación

- [Go Documentation](https://go.dev/doc/)
- [Cobra CLI](https://github.com/spf13/cobra)
- [Conventional Commits](https://www.conventionalcommits.org/)

### Comunidad

- [GitHub Issues](https://github.com/XebecCorporation/XebecCorporation.Dots/issues)
- [Discussions](https://github.com/XebecCorporation/XebecCorporation.Dots/discussions)

## Licencia

Al contribuir, aceptas que tus contribuciones serán licenciadas bajo MIT License.

---

¡Gracias por contribuir a XebecCorporation.Dots!
