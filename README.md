# config-read

[![Build Status](https://github.com/cumulus13/config-read/workflows/Build/badge.svg)](https://github.com/cumulus13/config-read/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/cumulus13/config-read)](https://goreportcard.com/report/github.com/cumulus13/config-read)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A cross-platform CLI tool for reading configuration files with beautiful syntax highlighting and automatic sensitive data masking.

## Features

- 🎨 **Beautiful Syntax Highlighting** - Default Fruity theme inspired by Python's Pygments
- 🔒 **Sensitive Data Masking** - Automatically masks passwords, tokens, API keys
- 📄 **Multiple Format Support** - YAML, JSON, TOML, INI, .env, and more
- 📟 **Built-in Pager** - Smart pagination with `less` (or `more` on Windows)
- ⚙️ **Configurable** - Custom themes, patterns, and behavior via config file
- 🚀 **Cross-Platform** - Linux, macOS, and Windows support
- 🎯 **Zero Dependencies** - Single binary distribution

## Installation

### Using Go

```bash
go install github.com/cumulus13/config-read/cmd/config-read@latest
```

### Binary Downloads

Download the latest binary for your platform from the [releases page](https://github.com/cumulus13/config-read/releases).

### Using Homebrew (macOS/Linux)

```bash
brew tap cumulus13/tap
brew install config-read
```

### Using Package Managers

**Debian/Ubuntu:**
```bash
curl -LO https://github.com/cumulus13/config-read/releases/latest/download/config-read_amd64.deb
sudo dpkg -i config-read_amd64.deb
```

**RHEL/CentOS/Fedora:**
```bash
sudo rpm -i https://github.com/cumulus13/config-read/releases/latest/download/config-read_amd64.rpm
```

## Quick Start

```bash
# Read a config file with default settings
config-read config.yaml

# Read without pager
config-read .env --no-pager

# Use a different theme
config-read appsettings.json --theme dracula

# Add custom sensitive patterns
config-read config.toml --sensitive-patterns "api_key,secret_token"
```

## Configuration

Create a `.config-read.yaml` file in your home directory:

```yaml
theme: fruity
pager_enabled: true
sensitive_patterns:
  - custom_secret
  - internal_token
  - private_key
```

## Supported File Types

- YAML (.yaml, .yml)
- JSON (.json)
- TOML (.toml)
- INI (.ini, .cfg, .conf)
- Environment files (.env)
- XML (.xml)
- HCL (.hcl, .tf)
- Dockerfile
- Makefile
- Properties files (.properties)

## Available Themes

- **fruity** (default) - Inspired by Python Pygments
- monokai
- dracula
- nord
- solarized-dark
- solarized-light
- github
- vs

## Sensitive Data Masking

By default, config-read automatically masks:

- `password`, `passwd`, `pwd`
- `secret`, `secret_key`
- `token`, `access_token`, `refresh_token`
- `api_key`, `apikey`, `api_secret`
- AWS credentials and session tokens
- Database connection strings
- JWT tokens
- OAuth client secrets
- And more...

Values are replaced with asterisks (`*`) matching the length of the original value.

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author
        
[Hadi Cahyadi](mailto:cumulus13@gmail.com)
    

[![Buy Me a Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cumulus13)

[![Donate via Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/cumulus13)
 
[Support me on Patreon](https://www.patreon.com/cumulus13)

## Acknowledgments

- [Chroma](https://github.com/alecthomas/chroma) - Syntax highlighting engine
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
