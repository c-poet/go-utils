# Repository guidelines

## CodeGraph

If a `.codegraph/` directory exists at the repository root, use CodeGraph before `rg`, `find`, or directly reading source files when locating or understanding code:

```bash
codegraph explore "<symbol name or code question>"
```

If `.codegraph/` is absent, proceed with normal repository tools.

## Go conventions

- Keep utilities grouped by responsibility in their own subpackages (for example, `cipher`, `json`, `os`, and `types`).
- Preserve exported API compatibility unless a change explicitly calls for a breaking change.
- Add or update focused tests for behavior changes.
- Format changed Go files with `gofmt`.
- Validate changes with `go test ./...`.

## Documentation

- Whenever a method is added or changed, update `README.md` in the same change. Update the relevant module description, public API documentation, and/or usage example so the documentation accurately reflects the current behavior.
- Update `README.md` when adding a public package.
- Document security-sensitive behavior and compatibility constraints near the relevant API.
