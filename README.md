# ai-commitizen

AI-powered helper for writing conventional commit messages, integrated with JIRA and multiple AI providers.

This tool will:
- Ask for a JIRA ticket number
- Fetch the ticket summary from JIRA
- Detect unstaged changes and offer to stage them for you (`git add .`)
- Read your staged Git diff
- Ask you to pick a commit type (feat, fix, chore, etc.)
- Generate a commit message using the configured AI provider, or prompt you to type one
- Optionally run `git commit -m "<message>"`

## Requirements

- Go (compatible with this module)
- A JIRA account with API access
- Config file at `~/.ai-commitizen/config.yaml`
- For AI mode: access to the configured provider (Anthropic API key, Gemini API key, or an Ollama instance)
- Git repository with changes (staged or unstaged)

## Installation

Recommended (uses installer + git alias):

```bash
make install
# or
sh install.sh
```

This will:
- Build the `ai-commitizen` binary into `$HOME/.local/bin` (or `$AI_COMMITIZEN_INSTALL_DIR`)
- Configure a global git alias `cz` that runs `ai-commitizen`

Make sure the install directory is in your `PATH`, for example (zsh):

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Manual build (alternative):

```bash
go build -o ai-commitizen ./cmd
```

You can then move the binary somewhere on your PATH.

## Configuration

Configuration is read from:

- `~/.ai-commitizen/config.yaml`

Example `config.yaml`:

```yaml
WithAI: true
Provider: anthropic # or "gemini" or "ollama"
Jira:
  Username: your-jira-username
  Password: your-jira-api-token-or-password
  BaseURL: https://your-company.atlassian.net
Anthropic:
  APIKey: your-anthropic-api-key
Gemini:
  APIKey: your-gemini-api-key
  Model: models/gemini-2.5-flash
Ollama:
  BaseURL: http://localhost:11434
  Model: qwen2.5-coder:7b
```

Fields:
- `WithAI`: set to `true` to use an AI model; `false` will ask you to type a commit message manually
- `Provider`: which AI backend to use: `anthropic`, `gemini`, or `ollama` (default is `anthropic` if empty)
- `Jira.Username`: JIRA username or email
- `Jira.Password`: JIRA API token or password (depending on your setup)
- `Jira.BaseURL`: Base URL of your JIRA instance
- `Anthropic.APIKey`: Anthropic API key used by the SDK
- `Gemini.APIKey`: Google Gemini API key
- `Gemini.Model`: Gemini model name (defaults to `models/gemini-2.5-flash` if empty)
- `Ollama.BaseURL`: Ollama server base URL (e.g. `http://localhost:11434`)
- `Ollama.Model`: Ollama model name (e.g. `qwen2.5-coder:7b`)

## Usage

After installation, from inside a Git repository:

```bash
# Recommended: stage explicitly
git add .
git cz

# Or: let ai-commitizen stage all changes when prompted
git cz
```

This runs `ai-commitizen` via the git alias. You can also run it directly:

```bash
ai-commitizen
# or, from source
go run ./cmd
```

Flow:
- If you have unstaged changes, the tool lists them and asks whether to stage all files (`git add .`).
- Input the JIRA ticket number when prompted (e.g. `STOL-6969`).
- Select a commit type from the list.
- The tool reads your staged diff and calls the configured AI provider (if `WithAI: true`), or asks you to type a message.
- You will see the full commit message printed, e.g.

  `STOL-6969: (feat) Generate commit message with AI`

- Confirm to perform `git commit` with that message, or cancel to abort.

## Notes

- Make sure your config file exists and is valid before running.
- The generated commit message is always based on staged changes. If you choose not to auto-stage, make sure you have staged the files you want included.
- If `WithAI` is true but the provider or its configuration (API key, Ollama URL/model, etc.) is missing or misconfigured, AI generation will fail and no commit will be made.
