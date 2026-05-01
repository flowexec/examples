<p align="center"><a href="https://flowexec.io"><img src="https://flowexec.io/icon.png" alt="flow" width="200"/></a></p>

<br>

A collection of real-world flow automation examples showing how developers use
[flow](https://flowexec.io) to automate their daily workflows. Intended as a
living reference and starting point for your own workspace.

## Getting Started

```sh
# Register this repo as a flow workspace
flow workspace add flow-examples . --set

# Browse all available executables
flow browse
```

## Structure

| Directory | What it shows |
|---|---|
| [`basics/`](basics/) | Core feature reference — one file per executable type |
| [`go-project/`](go-project/) | Full developer lifecycle (build, test, lint, release) — same pattern applies to any ecosystem (Node, Rust, Python…) |
| [`git/`](git/) | Git workflow helpers (commit, fetch/rebase, branch cleanup) |
| [`api/`](api/) | HTTP automation — GitHub REST API, webhook dispatch |
| [`docker/`](docker/) | Container workflows (build, run, push, clean) |
| [`kubernetes/`](kubernetes/) | kubectl automation (context, apply, pods, logs, shell) + Helm shared-library pattern (reusable installer called by app deployers) |
| [`setup/`](setup/) | Project onboarding (prereq checks, tool install, env config) |
| [`assets/`](assets/) | Supporting scripts and templates referenced by examples |

## Template

`exec-template.flow.tmpl` is an executable template — run it to scaffold a
starter flow file for a new project:

```sh
flow template generate NAME exec-template.flow.tmpl
```

## Validation

All `.flow` files in this repo are schema-validated in CI:

```sh
flow validate
```
