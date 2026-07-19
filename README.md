# Go Social Network

A study project built to explore **Hexagonal Architecture (Ports & Adapters)**
in Go, using a blog/social-network domain (users, posts, comments) as the
playground.

The core idea: the application core (domain + use cases) never depends on any
external technology. New ways of reaching the application — GraphQL today,
gRPC/HTTP/WebSocket/CLI later — are added purely as new **adapters**, without
touching the core.

## Goals

This repository exists primarily as a learning exercise for:

- Hexagonal Architecture / Ports & Adapters in practice
- GraphQL in Go (schema-first, resolvers, dataloaders, N+1 handling)
- Clean separation between business logic and delivery mechanisms
- Incrementally adding new adapters (gRPC, REST, WebSocket, CLI) without
  modifying the domain or use cases
- Swapping infrastructure (in-memory → Postgres, adding cache, messaging,
  etc.) without changing business rules

## Architecture

See [`ARCHITECTURE.md`](./ARCHITECTURE.md) for the full breakdown of the
directory structure, the responsibility of each layer, and the roadmap of
adapters to be added.

Quick summary:

```
Outside world → Adapter (in) → Port (driving) → Usecase → Port (driven) → Adapter (out) → Outside world
```

- `internal/app/domain` — pure business entities and rules, no external dependencies
- `internal/app/usecase` — application services, orchestrate domain logic
- `internal/app/ports` — interfaces the core exposes (driving) and requires (driven)
- `internal/infra` — configuration loading and shared resource wiring (DB connections, etc.)
- `internal/adapters/in` — driving adapters (currently: GraphQL)
- `internal/adapters/out` — driven adapters (currently: in-memory storage)

## Tech stack

| Concern | Choice |
|---|---|
| Language | Go |
| API | GraphQL ([gqlgen](https://github.com/99designs/gqlgen)) |
| Storage (initial) | In-memory |
| Storage (planned) | PostgreSQL |

This stack will grow as new topics are studied — see the roadmap in
[`ARCHITECTURE.md`](./ARCHITECTURE.md).

## Getting started

### Prerequisites

- Go 1.22+

### Running the project

```bash
go run ./cmd/server
```

The GraphQL Playground will be available at `http://localhost:<port>/` (port
configurable via environment variables — see `internal/infra/config`).

### Project layout at a glance

```
/blog
├── /cmd/server            # Application entry point
├── /internal
│   ├── /app                # Domain, use cases and ports (the core)
│   ├── /infra               # Config loading and resource wiring
│   └── /adapters
│       ├── /in/graphql      # GraphQL adapter
│       └── /out/db          # Storage adapters
├── /pkg                     # Generic, reusable utilities
└── ARCHITECTURE.md
```

## Contributing

This is a personal study project — not currently open for external
contributions, but feel free to fork it for your own learning.

## License

Not defined yet.