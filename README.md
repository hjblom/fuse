# fuse
Fuse is a golang microservice generator.

> The project is in active development. It does not generate any code yet.


## Installation

To install this cli, run the following command:

```bash
$: go install github.com/hjblom/fuse@latest
```

Verify that the binary installed correctly by running the `--help` command.

```bash
$: fuse --help

A go microservice generator framework.

Usage:
  fuse [command]

Available Commands:
  add         Add components to the project
  completion  Generate the autocompletion script for the specified shell
  generate    Generate code for the project based on the configuration file
  help        Help about any command
  init        Initialize a new fuse project
  visualize   Visualize the project dependency graph

Flags:
  -c, --config string   Path to the config file (default ".fuse.yaml")
  -h, --help            help for fuse
  -t, --toggle          Help message for toggle

Use "fuse [command] --help" for more information about a command.
```

## Features

### Project initialization (uncompleted)

The `init` command initializes a new project. It will create a new directory with the name of the project and generate a default project structure.

```bash
$: fuse init my-project
```

### Add components

The `add` command adds a new component to the project. It will create a new directory with the name of the component and generate a default component structure.

```bash
$: fuse add client
```

#### Options

The `add` command has a few options that can be used to customize the component.

| Option | Description |
| --- | --- |
| `--path` | Path to where the package should be placed |
| `--require` | Add a requirement to the component that will be injected |
| `--tag` | Add a tag to the component, indicating what additional methods should be generated. |

### Visualize

The `visualize` command generates a dependency graph of the project. It will generate a `graph.svg` file in the current directory.

```bash
fuse visualize
```
