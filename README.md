<p align="center"><a href="https://flowexec.io"><img src="https://flowexec.io/_media/logo.png" alt="flow" width="200"/></a></p>

<br>

This repository contains example flow files demonstrating various executable types and workflows.

### Generated Files

- **`*.flow.yaml`** - Flow file examples demonstrating different executable types
- **`assets/`** - Supporting files (scripts, templates, data files) referenced by the examples

### Examples Include

- **Exec Examples** - Running shell commands and scripts
- **Parallel Examples** - Executing multiple tasks concurrently
- **Serial Examples** - Sequential task execution
- **Request Examples** - HTTP requests and API interactions
- **Render Examples** - Template rendering with data
- **Launch Examples** - Launching processes
- **Executable Template Example** - Define flowfile templates for reusable workflows patterns

## Generating Examples

The examples are generated using the builder tool in `cmd/builder/`. To regenerate all examples:

```bash
go run cmd/builder/*.go
# or run via flow with
flow validate
```
