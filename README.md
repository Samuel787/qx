# ct

A command-line tool built with Go.

## Installation

### From Homebrew (coming soon)

```bash
brew install yourusername/tap/ct
```

### From Source

```bash
git clone https://github.com/yourusername/ct
cd ct
make install
```

## Usage

```bash
ct --help
ct version
ct hello World
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

- `ct version` - Display the version of ct
- `ct hello [name]` - Say hello to someone

## Homebrew Formula

To publish this to Homebrew, you'll need to create a tap repository and add a formula. See the [Homebrew documentation](https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap) for more details.

### Quick Steps for Homebrew Publishing:

1. Create a repository named `homebrew-tap` on GitHub
2. Add a `Formula/ct.rb` file with the binary details
3. Users can install with: `brew tap yourusername/tap && brew install ct`

## License

MIT License
