# aether
Ethical hacking and security research tool for authorized environments only.


## Important Disclaimer

**Aether is intended for educational purposes and authorized security auditing only.**

- Only use this tool on systems you own or have **explicit written permission** to test.
- Unauthorized use may violate laws in your country (e.g., Computer Fraud and Abuse Act in the US).
- The author is not responsible for any misuse or damage caused by this tool.
- Users are fully responsible for complying with all applicable laws and regulations.

## Installation
```bash
git clone https://github.com/BosBJJ/aether.git
cd aether
go build -o aether .

# DNS reconnaissance
aether recon dns -t example.com

# Subdomain enumeration
aether recon subdomain -t example.com --size medium
aether recon subdomain -t example.com --size large --threads 100




## Wordlists

The large wordlist (1M entries) is not included in this repository due to file size limits.

To use `--size large`, download it from SecLists:

https://github.com/danielmiessler/SecLists/blob/master/Discovery/DNS/subdomains-top1million-full.7z

Place the file at:
data/wordlists/subdomains-top1million-full.txt