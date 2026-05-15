# envdiff

> Compare `.env` files across environments and surface missing or mismatched keys.

---

## Installation

```bash
go install github.com/yourusername/envdiff@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/envdiff.git
cd envdiff && go build -o envdiff .
```

---

## Usage

```bash
envdiff [flags] <base-file> <compare-file> [additional-files...]
```

### Example

```bash
envdiff .env.example .env.production
```

**Sample output:**

```
MISSING in .env.production:
  - DATABASE_URL
  - REDIS_HOST

MISMATCHED keys (present but values differ):
  ~ LOG_LEVEL  (.env.example: "debug" | .env.production: "info")
```

Compare multiple environments at once:

```bash
envdiff .env.example .env.staging .env.production
```

### Flags

| Flag | Description |
|------|-------------|
| `--keys-only` | Only compare key names, ignore values |
| `--strict` | Exit with non-zero status if any differences are found |
| `--json` | Output results as JSON |

---

## Contributing

Pull requests are welcome. Please open an issue first to discuss any significant changes.

---

## License

[MIT](LICENSE)