# Clickstream

<p align="center">
    <img src="docs/images/clickstream.png" width="256" height="256" />
</p>

<p align="center">
    <a href="/docs/services/clickstream/index.md">Project Documentation</a> | 
    <a href="docs/index.md">Service Documentation</a> |
    <a href="./CHANGELOG.md">Changelog</a>
</p>

# Usage
## Run the service
```bash
$ MODE=dev make up # will use "docker/modes/dev" environment file
```

## Down the service
```bash
$ MODE=dev make down
```

## Show the summary of the compose config
```bash
$ MODE=dev make config
```

## Stop the service
```bash
$ MODE=dev make stop
```
