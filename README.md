# Setup

1. install `direnv`

```bash
    brew install direnv
```

2. configure `direnv`

```bash
    mkdir -p ~/.config/direnv
    touch ~/.config/direnv/config.toml
    echo "[whitelist]" >> ~/.config/direnv/config.toml
    echo "prefix = []" >> ~/.config/direnv/config.toml

```

```bash
    echo 'eval "$(direnv hook bash)"' >> ~/.zshrc
```

3. run `direnv allow` in the project root
    
```bash
    direnv allow .
```

