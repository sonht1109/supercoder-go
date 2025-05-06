# Supercoder-go
A fork of [huytd/supercoder](https://github.com/huytd/supercoder) but in Golang.

<img width="550" alt="SCR-20250506-mrnh" src="https://github.com/user-attachments/assets/ed374f46-e0e8-4523-a56f-42a2e9767d7a" />

### Installation
1. Download file that works with your OS:

Visit https://github.com/sonht1109/supercoder-go/releases

Download file and rename it to `supercoder-go` or whatever you want.

2. Make that file executable:

For example, with Linux/MacOS
```bash
mv supercoder-go /usr/local/bin
chmod +x /usr/local/bin/supercoder-go
```

3. Set ENV variables:

Make sure that you setup following variables
```bash
export OPENAI_API_KEY=<API_KEY>
export MODEL=<MODEL>
```

If you want to perform web search tool, please add
```bash
export SEARXNG_BASE_URL=<SEARXNG_URL> // default value is "https://searx.be" but it is mostly failed to call. Best choice is to self host searxng by your self.
```

### Usage
In your terminal, type:
```bash
supercoder-go
```

And enjoy.
