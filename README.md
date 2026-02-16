# ai-commitizen

AI-powered helper for writing conventional commit messages, integrated with JIRA and Anthropic Claude.

This tool will:
- Ask for a JIRA ticket number
- Fetch the ticket summary from JIRA
- Read your staged Git diff
- Ask you to pick a commit type (feat, fix, chore, etc.)
- Generate a commit message using Anthropic Claude
- Optionally run `git commit -m "<message>"`

## Requirements

- Go (compatible with this module)
- A JIRA account with API access
- An Anthropic API key
- Git repository with changes already staged (`git add ...`)

## Installation

From the project root:

```bash
# Build a binary
cd scripts/commitizen

go build -o ai-commitizen ./cmd

# Or run directly
go run ./cmd
```

You can also move the built binary somewhere on your PATH, e.g. `~/bin`.

## Configuration

Configuration is read from:

- `~/.ai-commitizen/config.yaml`

Example `config.yaml`:

```yaml
with_ai: true
jira:
  username: your-jira-username
  password: your-jira-api-token-or-password
  base_url: https://your-company.atlassian.net
anthropic:
  api_key: your-anthropic-api-key
```

Fields:
- `with_ai`: set to `true` to use Anthropic for commit message generation
- `jira.username`: JIRA username or email
- `jira.password`: JIRA API token or password (depending on your setup)
- `jira.base_url`: Base URL of your JIRA instance
- `anthropic.api_key`: Anthropic API key used by the SDK

## Usage

Run the tool from inside a Git repository after staging your changes:

```bash
ai-commitizen
# or, if running via Go
cd scripts/commitizen
go run ./cmd
```

Flow:
- Input the JIRA ticket number when prompted (e.g. `STOL-6969`)
- Select a commit type from the list
- The tool reads your staged diff and calls Anthropic to generate a message
- You will see the full commit message printed, e.g.

  `STOL-6969: (feat) Generate commit message with AI`

- Confirm `Y` to perform `git commit` with that message, or cancel to abort.

## Notes

- Make sure your config file exists and is valid before running.
- The tool only looks at staged changes; use `git add` first.
- If `with_ai` is false or the API key is missing, generation will fail.