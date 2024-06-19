# uptimectl: A command-line interface for interacting with Better Uptime

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

## Overview

**uptimectl** is a command-line tool designed to interact seamlessly with the [Better Uptime](https://betterstack.com/uptime) API. It supports various functionalities such as managing incidents, monitors, and monitor groups.

## Features

- **Incident Management:** List and manage incidents.
- **Monitor Management:** View and control monitors and monitor groups.
- **On-Call Management:** Show on-call contacts.

## Installation

### Homebrew

```bash
brew install uptime-cli/cli/uptimectl
```

### Download

Download the latest binary from [GitHub releases](https://github.com/uptime-cli/uptimectl/releases/latest).

## Documentation

See [docs](/docs/)

## Usage

Authenticate using an API Token (see [uptime docs](https://betterstack.com/docs/uptime/api/getting-started-with-uptime-api/#obtaining-an-uptime-api-token) for how to get an token):
```bash
❯ uptimectl auth login --token <TOKEN>
```

List recent incidents for your team:
```bash
❯ uptimectl incidents ls -a
```

Show on-call contacts:
```bash
❯ uptimectl on-call
```

View monitors and their status:

```bash
❯ uptimectl monitors ls
```
