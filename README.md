# eventually-eats

Your food. Exactly Once. Eventually.

<!-- toc -->

* [How To run](#how-to-run)
* [Contributing](#contributing)
  * [Open in a container](#open-in-a-container)
  * [Commit style](#commit-style)

<!-- Regenerate with "pre-commit run -a markdown-toc" -->

<!-- tocstop -->

## How To run

Running is very simple:

```sh
docker compose up
```

This will start a new Temporal server, the workflow and the web application:

* [Temporal](http://localhost:8233)
* [Web app](https://localhost:9999)

## Contributing

### Open in a container

* [Open in a container](https://code.visualstudio.com/docs/devcontainers/containers)

### Commit style

All commits must be done in the [Conventional Commit](https://www.conventionalcommits.org)
format.

```git
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```
