# qx (Quick Execute)

 A terminal command prefiller that uses Groq AI to generate shell commands from natural language queries. 

Eg:
```bash
qx search for pattern in all files
# → grep -r pattern .
```

Instead of trying to remember complex command syntax, just describe what you want to do in plain English, and `qx` will generate the command for you. The command is automatically copied to your clipboard, ready to paste and execute.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Setup](#setup)
- [Usage](#usage)
  - [Basic Usage with qx](#basic-usage-with-qx)
  - [Using the qq() Function](#using-the-qq-function)
- [Commands](#commands)
- [Examples](#examples)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Features

- 🤖 **AI-Powered**: Uses Groq's fast LLM to understand your command intent
- ⚡ **Instant**: Commands are generated in milliseconds
- 📋 **Auto-Copy**: Generated commands are automatically copied to your clipboard
- 🔧 **Shell Integration**: Install a `qq()` shell function for seamless integration
- 🔐 **Secure**: Your API key is stored locally in your shell profile
- 🎯 **Cross-Platform**: Works on macOS, Linux, and Windows
- ✨ **Free**: Uses Groq's free tier (no credit card required)

## Installation

### From Homebrew

```bash
brew tap Samuel787/tap
brew install qx
```

### From Source

```bash
git clone https://github.com/Samuel787/qx
cd qx
make build
make install
```

This builds the binary and installs it to `/usr/local/bin/qx`.

## Setup

### Getting Your Groq API Key

1. Visit [https://console.groq.com](https://console.groq.com)
2. Sign up for a free account
3. Navigate to the **API Keys** section
4. Click **Create API Key** to generate a new key
5. Copy the API key

### About Groq's Free Tier

The Groq API is **completely free** with the following limits:

- **Rate Limit**: 30 requests per minute
- **Input Tokens**: Up to 8,000 tokens per request
- **Models**: Access to fast LLM models (Llama 3.3 70B, Mixtral, etc.)
- **No Credit Card**: Use indefinitely at no cost

For more details, see the [Groq Documentation](https://console.groq.com/docs)

### Set Your API Key

```bash
qx set-key YOUR_GROQ_API_TOKEN
```

Example:
```bash
qx set-key gsk_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

The command will:
1. Add the API key to your shell configuration (`~/.zshrc` or `~/.bashrc`)
2. Copy the activation command to your clipboard
3. Display instructions for reloading your shell

After running this command, follow the on-screen instructions to activate the API key.

## Usage

### Basic Usage with qx

Query Groq AI for a command using natural language:

```bash
qx <your query here>
```

The generated command is:
- Displayed in your terminal
- **Automatically copied to your clipboard**
- Ready to paste and execute

#### Examples:

```bash
# List all files in the current directory
qx list all files in current directory
# Output: ls -la
# ✅ Copied to clipboard!

# Find files named "test"
qx find files named test
# Output: find . -name test
# ✅ Copied to clipboard!

# Count lines in a Python file
qx count lines in a Python file
# Output: wc -l *.py
# ✅ Copied to clipboard!

# Git: Show commit history
qx show git commit history
# Output: git log --oneline
# ✅ Copied to clipboard!

# Git: Create a new branch
qx create a new git branch named feature-x
# Output: git checkout -b feature-x
# ✅ Copied to clipboard!
```

### Using the qq() Function

For an even more seamless experience, install the `qq()` shell function. This allows you to use `qq` directly from your terminal without even typing `qx`.

#### Installing qq()

```bash
qx qq-install
```

This will:
1. Add a `qq()` function to your shell profile (`~/.zshrc` or `~/.bashrc`)
2. Copy the reload command to your clipboard
3. Display instructions for activating the function

After installation, reload your shell by one of these methods:

```bash
# Option 1: Run the source command (copied to clipboard)
source ~/.zshrc
# or
source ~/.bashrc

# Option 2: Close and reopen your terminal

# Option 3: Start a new shell
exec $SHELL
```

#### Using qq()

Once installed, you can use `qq` directly:

```bash
qq <your query>
```

**On zsh**: The command is prefilled on your command line, ready to execute with Enter

**On bash**: The command is copied to clipboard and added to history. Press `Cmd+V` (or `Ctrl+V`) to paste, or press the Up arrow to recall from history.

#### qq() Examples:

```bash
# List files
qq list all files
# zsh: Command appears on command line
# bash: Command copied to clipboard

# Show disk usage
qq show disk usage
# Output: df -h

# Kill a process by name
qq kill process named node
# Output: pkill node

# Change git branch
qq change git branch
# Output: git checkout

# Search for files
qq find all javascript files
# Output: find . -name "*.js"

# View recent git changes
qq show recent git changes
# Output: git log --oneline -n 10
```

#### Uninstalling qq()

If you want to remove the `qq()` function:

```bash
qx qq-uninstall
```

This will:
1. Remove the `qq()` function from your shell profile
2. Copy the reload command to your clipboard
3. Display instructions for reloading your shell

After running this, reload your shell using one of the methods above.

## Commands

### `qx <query>`

Generate a command from a natural language query.

**Aliases**: None (this is the default behavior)

**Example**:
```bash
qx how to list files recursively
```

**Output**: `ls -laR` (copied to clipboard)

---

### `qx --raw <query>` or `qx -r <query>`

Generate a command and output it as raw text without formatting or clipboard copy.

Useful for scripting or when you want just the command text without the pretty formatting.

**Example**:
```bash
qx --raw count lines in all python files
```

**Output**: `wc -l *.py`

---

### `qx set-key <api-token>`

Set your Groq API key in your shell configuration.

**Usage**:
```bash
qx set-key YOUR_GROQ_API_TOKEN
```

**What it does**:
1. Stores the API key as `QX_GROQ_KEY` environment variable
2. Updates your shell profile (`~/.zshrc` or `~/.bashrc`)
3. Copies the reload command to clipboard
4. Shows instructions for activation

**Example**:
```bash
qx set-key gsk_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
# API key has been set. Now run the following to activate:
# source ~/.zshrc
# ✅ Command copied to clipboard!
```

---

### `qx qq-install`

Install the `qq()` shell function for seamless integration.

**Usage**:
```bash
qx qq-install
```

**What it does**:
1. Adds the `qq()` function to your shell profile
2. Copies the reload command to clipboard
3. Shows instructions for different reload methods

**Outputs instructions for**:
- Running `source ~/.zshrc` (copied to clipboard)
- Closing/reopening your terminal
- Running `exec $SHELL`

---

### `qx qq-uninstall`

Remove the `qq()` function from your shell profile.

**Usage**:
```bash
qx qq-uninstall
```

**What it does**:
1. Removes the `qq()` function from your shell profile
2. Cleans up extra blank lines
3. Copies the reload command to clipboard
4. Shows instructions for reloading

---

### `qx --help` or `qx -h`

Display help information about qx and available commands.

**Usage**:
```bash
qx --help
qx -h
```

---

### `qx hello <text>`

Copy text to your clipboard and display it.

Mostly a utility command for testing clipboard functionality.

**Usage**:
```bash
qx hello "Hello, World!"
```

---

## Examples

### Basic Usage Examples

```bash
# List files with details
qx list files with details
# → ls -la

# Show current directory
qx print working directory
# → pwd

# View file contents
qx view contents of file.txt
# → cat file.txt

# Create a directory
qx create a new directory named myproject
# → mkdir myproject
```

### Advanced Examples

```bash
# Find large files
qx find files larger than 100MB
# → find . -size +100M

# Search text in files
qx search for pattern in all files
# → grep -r pattern .

# Compress files
qx create tar gz archive of current directory
# → tar -czf archive.tar.gz .

# View system info
qx show system information
# → uname -a

# List installed packages
qx list installed packages
# → npm list -g (for Node)
# → pip list (for Python)
```

### Git Examples

```bash
# View commit log
qx show git commit history
# → git log --oneline

# Create new branch
qx create git branch named feature-auth
# → git checkout -b feature-auth

# Show changes
qx show git changes
# → git diff

# Stage changes
qx stage all changes in git
# → git add .

# Commit changes
qx commit changes with message
# → git commit -m "message"

# Push to remote
qx push changes to git remote
# → git push origin main
```

### Using with qq()

```bash
# Simple command generation
qq how to find files
# → find . -name pattern

# Complex queries
qq recursively find all json files and count them
# → find . -name "*.json" | wc -l

# Git workflows
qq show my unpushed commits
# → git log origin/main..HEAD

# System administration
qq kill all node processes
# → pkill -f node
```

## Troubleshooting

### Error: "command not found: qx"

**Solution**: Make sure qx is installed and in your PATH.

```bash
# If installed from source
make install

# Verify installation
which qx
qx --help
```

### Error: "qq: command not found"

**Solution**: The `qq()` function hasn't been installed yet, or your shell hasn't been reloaded.

```bash
# Install the function
qx qq-install

# Reload your shell
source ~/.zshrc  # or ~/.bashrc
```

### Error: "QX_GROQ_KEY not set"

**Solution**: Your API key hasn't been set. Set it with:

```bash
qx set-key YOUR_GROQ_API_TOKEN
# Then reload your shell:
source ~/.zshrc  # or ~/.bashrc
```

### Error: "unknown command" with qq()

**Solution**: Make sure you set the API key and reloaded your shell:

```bash
# Check if API key is set
echo $QX_GROQ_KEY

# If not set, do this:
qx set-key YOUR_GROQ_API_TOKEN
source ~/.zshrc

# Then try again:
qq list files
```

### Generated Command is Wrong

The AI does its best, but sometimes might misinterpret queries. Try:

1. Being more specific: `qx find files named test.js` instead of `qx find test files`
2. Using common command patterns: `qx grep pattern in files` instead of `qx search for text`
3. Checking the Groq API status at [https://console.groq.com](https://console.groq.com)

## License

MIT License - See LICENSE file for details

## Contributing

Found a bug or want to contribute? Feel free to open an issue or pull request on [GitHub](https://github.com/Samuel787/qx).

---

**Built by [Samuel787](https://github.com/Samuel787) and [Mohamedirfansh](https://github.com/Mohamedirfansh)**
