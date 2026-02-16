# ai-commitizen

AI-powered helper for writing conventional commit messages, integrated with JIRA and AI providers.

This tool will:
- Ask for a JIRA ticket number
- Fetch the ticket summary from JIRA
- Read your staged Git diff
- Ask you to pick a commit type (feat, fix, chore, etc.)
- Generate a commit message using the configured AI provider, or prompt you to type one
- Optionally run `git commit -m "<message>"`

## Requirements

- Go (compatible with this module)
- A JIRA account with API access
- Config file at `~/.ai-commitizen/config.yaml`
- For AI mode: an API key for the configured provider (Anthropic, Gemini, etc.)
- Git repository with changes already staged (`git add ...`)

## Installation

Recommended (uses installer + git alias):

```bash
cd scripts/commitizen
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
cd scripts/commitizen
go build -o ai-commitizen ./cmd
```

You can then move the binary somewhere on your PATH.

## Configuration

Configuration is read from:

- `~/.ai-commitizen/config.yaml`

Example `config.yaml`:

```yaml
WithAI: true
Provider: anthropic # or "gemini" (when wired up)
Jira:
  Username: your-jira-username
  Password: your-jira-api-token-or-password
  BaseURL: https://your-company.atlassian.net
Anthropic:
  APIKey: your-anthropic-api-key
Gemini:
  APIKey: your-gemini-api-key
  Model: models/gemini-2.5-flash
```

Fields:
- `WithAI`: set to `true` to use an AI model; `false` will ask you to type a commit message manually
- `Provider`: which AI backend to use (e.g. `anthropic`, `gemini`)
- `Jira.Username`: JIRA username or email
- `Jira.Password`: JIRA API token or password (depending on your setup)
- `Jira.BaseURL`: Base URL of your JIRA instance
- `Anthropic.APIKey`: Anthropic API key used by the SDK
- `Gemini.APIKey`: Google Gemini API key
- `Gemini.Model`: Gemini model name (defaults to `models/gemini-2.5-flash` if empty)

## Usage

After installation, from inside a Git repository with staged changes:

```bash
git add .
git cz
```

This runs `ai-commitizen` via the git alias. You can also run it directly:

```bash
ai-commitizen
# or, from source
cd scripts/commitizen
go run ./cmd
```

Flow:
- Input the JIRA ticket number when prompted (e.g. `STOL-6969`)
- Select a commit type from the list
- The tool reads your staged diff and calls the configured AI provider (if `WithAI: true`), or asks you to type a message
- You will see the full commit message printed, e.g.

  `STOL-6969: (feat) Generate commit message with AI`

- Confirm to perform `git commit` with that message, or cancel to abort.

## Notes

- Make sure your config file exists and is valid before running.
- The tool only looks at staged changes; use `git add` first.
- If `WithAI` is true but the provider or API key is missing/misconfigured, AI generation will fail and no commit will be made.
