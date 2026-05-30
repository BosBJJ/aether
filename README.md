# aether
[![Go Reference](https://pkg.go.dev/badge/github.com/BosBJJ/aether.svg)](https://pkg.go.dev/github.com/BosBJJ/aether)
Ethical hacking and security research tool for authorized environments only.


## Important Disclaimer

**Aether is intended for educational purposes and authorized security auditing only.**

- Only use this tool on systems you own or have **explicit written permission** to test.
- Unauthorized use may violate laws in your country (e.g., Computer Fraud and Abuse Act in the US).
- The author is not responsible for any misuse or damage caused by this tool.
- Users are fully responsible for complying with all applicable laws and regulations.

## Installation

### For Users
Install the binary directly to your `GOBIN`:
```bash
go install github.com/BosBJJ/aether@latest
```
### For Developers (Build from source)

```bash
git clone https://github.com/BosBJJ/aether.git
cd aether
go build -o aether .
```

## Global Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--output`, `-o` | `table` | Output format: `table`, `json`, `html` |
| `--save` | `false` | Save scan results to the database |
| `--threads` | `50` | Number of concurrent threads |
| `--timeout` | `5` | Timeout in seconds for network operations |
| `--quiet`, `-q` | `false` | Show minimal output |
| `--verbose`, `-v` | `false` | Show detailed output |
| `--pretty`, `-p` | `false` | Pretty print JSON output |


## Commands

### `crypto` — Cryptography utilities
| Command | Description |
|---------|-------------|
| `aether crypto hash -t myfile.txt` | Generate hash of text or a file |
| `aether crypto encrypt` | Encrypt file or text with AES256 |
| `aether crypto decrypt` | Decrypt file or text with AES256 |
| `aether crypto generate` | Generate a key or password |
| `aether crypto analyze` | Analyze password strength |
| `aether crypto hmac` | Generate or verify HMAC |

### `http` — HTTP analysis and inspection
| Command | Description |
|---------|-------------|
| `aether http info -t example.com` | Basic HTTP page information and metadata |
| `aether http headers -t example.com` | Analyze HTTP security headers |
| `aether http analyze -t example.com` | Analyze URL for suspicious parameters |
| `aether http fingerprint -t example.com` | Fingerprint web server technology stack |

### `recon` — Domain, DNS, and passive reconnaissance
| Command | Description |
|---------|-------------|
| `aether recon dns -t example.com` | DNS reconnaissance on a domain |
| `aether recon whois -t example.com` | WHOIS lookup for a domain or IP |
| `aether recon passive -t example.com` | Passive recon using certificate transparency logs |
| `aether recon subdomain -t example.com` | Subdomain enumeration using a DNS wordlist |

### `scan` — Host and port scanning
| Command | Description |
|---------|-------------|
| `aether scan ports -t example.com` | Scan ports on a target |
| `aether scan hosts -t 192.168.1.0/24` | Discover live hosts on a network |
| `aether scan banners -t example.com` | Grab banners from specified ports |
| `aether scan service -t example.com` | Full service scan (ports + banners) |

### `report` — Manage saved scans
| Command | Description |
|---------|-------------|
| `aether report list` | List all saved scans |
| `aether report view 42` | View full details of a saved scan |
| `aether report export 42 --file scan.html -o html` | Export a scan to a file |
| `aether report delete 42` | Delete a saved scan |



## Wordlists

The large wordlist (1M entries) is not included in this repository due to file size limits.

To use `--size large`, download it from SecLists:

https://github.com/danielmiessler/SecLists/blob/master/Discovery/DNS/subdomains-top1million-full.7z

Place the file at `data/wordlists/subdomains-top1million-full.txt`

## Configuration

Aether uses environment variables for paths when running outside the project directory.

| Variable | Default | Description |
|----------|---------|-------------|
| `AETHER_DB_PATH` | `aether.db` | Path to the SQLite database file |
| `AETHER_WORDLIST_PATH` | `data/wordlists` | Path to the wordlists directory |


Add these to your `~/.bashrc` or `~/.zshrc` to persist across sessions, then run `source ~/.bashrc` to reload.

`export AETHER_DB_PATH=$HOME/aether/aether.db`

`export AETHER_WORDLIST_PATH=$HOME/aether/data/wordlists`
