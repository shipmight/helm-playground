# helm-playground

This repository contains the source code for https://helm-playground.com

## How it works

A piece of Go code is compiled to a WASM module which can be ran in the browser. This code implements a simple function which takes two inputs:

- YAML template
- YAML values

Then it simply renders the given template with the given values using [Sprig](https://github.com/Masterminds/sprig), which is also what Helm uses.

The WASM module is compiled in a GitHub action. You can find the workflow in [`.github/workflows/compile.yaml`](.github/workflows/compile.yaml). When a commit is pushed to `master`, the workflow is triggered, the code is compiled and committed back to `master` with the commit message `[GitHub action] WASM module`. The `master` branch is hosted live via GitHub Pages at https://helm-playground.com.

## Development

### Test the golang code

```bash
make test
```

### Build

```bash
make build
```

### Test the built WASM code in browser

```bash
make browser-test
```

### Locally develop in browser

```bash
npx http-server -c-1
```
