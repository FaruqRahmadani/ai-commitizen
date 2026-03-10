# AI Commitizen 🚀

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8.svg?style=flat&logo=go)

> **AI-powered CLI tool for writing conventional commit messages, integrated with JIRA and multiple AI providers.**

Stop writing generic commit messages. Let AI analyze your code changes and generate meaningful, conventional commits for you—contextualized with your Jira tickets.

---

## ✨ Features

- 🤖 **Multi-LLM Support**: Works with **Anthropic (Claude)**, **Google Gemini**, and **Ollama** (for local privacy-focused models).
- 🎫 **Jira Integration**: Fetches ticket titles and types directly from Jira to prefix your commits (e.g., `STOL-123: (feat) ...`).
- 📝 **Conventional Commits**: Enforces industry-standard commit formats (`feat`, `fix`, `chore`, etc.).
- 🔄 **Smart Git Workflow**:
  - Detects unstaged changes and offers to stage them (`git add .`).
  - Reads staged diffs automatically.
  - Connects to git hooks (pre-commit, linters) seamlessly.
- ✏️ **Interactive Editing**: Not happy with the AI's suggestion? Edit the message in your default `$EDITOR` (Vim, Nano, VS Code, etc.) before committing.
- ⚡ **Git Alias**: Installs as `git cz` for a native git experience.

---

## 📦 Installation

### Quick Install (Recommended)

You can install the binary and configure the git alias automatically using the provided script:

```bash
make install
# OR
sh install.sh
```

This will:
1. Build the binary to `$HOME/.local/bin` (ensure this is in your `PATH`).
2. Set up `git cz` as a global alias.

### Manual Installation

If you prefer to build it yourself:

```bash
go build -o ai-commitizen ./cmd
mv ai-commitizen /usr/local/bin/ # or any directory in your PATH
git config --global alias.cz '!ai-commitizen'
```

---

## ⚙️ Configuration

Create a configuration file at `~/.ai-commitizen/config.yaml`.

```yaml
# Enable AI generation (set to false for manual mode)
WithAI: true

# AI Provider: "anthropic", "gemini", or "ollama"
Provider: anthropic

# Jira Integration (Optional)
Jira:
  Username: your-email@company.com
  Password: your-api-token # Use API Token, not account password
  BaseURL: https://your-company.atlassian.net

# Provider Configuration
Anthropic:
  APIKey: sk-ant-...

Gemini:
  APIKey: AIza...
  Model: models/gemini-2.5-flash

Ollama:
  BaseURL: http://localhost:11434
  Model: qwen2.5-coder:7b
```

> **Note**: If `WithAI` is `false`, the tool will skip AI generation and just prompt you for the message, but you still get the Jira integration benefits.

---

## 🚀 Usage

Inside any git repository, simply run:

```bash
git cz
```

### The Workflow

1.  **Stage Changes**: If you haven't staged files, the tool will list them and ask to `git add .`.
2.  **Ticket Number**: Enter your Jira ticket (e.g., `PROJ-123`). The tool fetches the summary.
3.  **Commit Type**: Select the type of change:
    - `feat`: A new feature
    - `fix`: A bug fix
    - `chore`: Maintenance
    - ...and more.
4.  **AI Generation**: The tool analyzes your `git diff --staged` and generates a commit message.
5.  **Review & Edit**:
    - **Yes**: Commit immediately.
    - **Edit**: Opens your default editor to tweak the message.
    - **No**: Cancel the operation.

### Example Output

```text
? Ticket Number: STOL-6969
[INFO] Fetching Jira ticket...
You're working on [Story] Add AI generation feature (In Progress)

? Commit Type: feat
[INFO] Generating commit message...

Generated commit message:
  STOL-6969: (feat) implement AI-powered commit generation using Anthropic API

? Are you sure you want to commit with this message? [Yes/Edit/No]
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to submit a Pull Request.

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/AmazingFeature`).
3.  Commit your changes.
4.  Push to the branch.
5.  Open a Pull Request.

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
