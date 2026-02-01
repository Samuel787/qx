# qx

A command-line tool built with Go.

## Installation

### From Homebrew (coming soon)

```bash
brew install yourusername/tap/qx
```

### From Source

```bash
git clone https://github.com/yourusername/qx
cd qx
make install
```

## Usage

This tool internally uses GROQ to get the CLI command. Thus, GROQ API Key is required for this tool to work. 

### Getting Your GROQ API Key

1. Visit [https://console.groq.com](https://console.groq.com)
2. Sign up for a free account
3. Navigate to the **API Keys** section
4. Click **Create API Key** to generate a new key
5. Copy the API key and use it with the `qx set-key` command

### GROQ Free Tier Information

The GROQ API is **free to use** with the following limitations:

- **Rate Limit**: 30 requests per minute
- **Tokens**: Up to 8,000 input tokens per request
- **Models**: Access to Groq's fast LLM models (Llama, Mixtral, etc.)
- **No credit card required**: Use the free tier indefinitely

For more details, visit the [GROQ Documentation](https://console.groq.com/docs)

### Set Your GROQ API Key

Before using qx, you need to set your GROQ API key:

```bash
qx set-key YOUR_GROQ_API_TOKEN
```

The command will:
1. Add the API key to your shell configuration file (`~/.zshrc` or `~/.bashrc`)
2. Display the command to activate the environment variable
3. Copy the source command to your clipboard for easy pasting

Example:
```bash
qx set-key gsk_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
# Then simply paste and run the source command that was automatically copied to clipboard
```

## Development

### Building

```bash
make build
```

### Installing Locally

```bash
make install
```

### Running in Development Mode

```bash
make dev
```

### Running Tests

```bash
make test
```

## Commands

- `qx set-key [token]` - Set your GROQ API key in shell configuration

## Homebrew Formula

To publish this to Homebrew, you'll need to create a tap repository and add a formula. See the [Homebrew documentation](https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap) for more details.

### Quick Steps for Homebrew Publishing:

1. Create a repository named `homebrew-tap` on GitHub
2. Add a `Formula/qx.rb` file with the binary details
3. Users can install with: `brew tap yourusername/tap && brew install qx`

## License

MIT License
