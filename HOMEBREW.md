# Publishing to Homebrew

This guide explains how to publish the `ct` CLI tool to Homebrew.

## Step 1: Create a Homebrew Tap Repository

1. Create a new GitHub repository named `homebrew-tap` in your GitHub account
2. Clone it locally: `git clone https://github.com/yourusername/homebrew-tap`

## Step 2: Create the Formula

In your `homebrew-tap` repository, create the directory structure and formula file:

```bash
mkdir -p Formula
cat > Formula/ct.rb << 'EOF'
class Ct < Formula
  desc "A command-line tool"
  homepage "https://github.com/yourusername/ct"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/ct/releases/download/v0.1.0/ct_0.1.0_darwin_arm64.tar.gz"
      sha256 "REPLACE_WITH_ARM64_SHA256"
    else
      url "https://github.com/yourusername/ct/releases/download/v0.1.0/ct_0.1.0_darwin_amd64.tar.gz"
      sha256 "REPLACE_WITH_AMD64_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/ct/releases/download/v0.1.0/ct_0.1.0_linux_arm64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_ARM64_SHA256"
    else
      url "https://github.com/yourusername/ct/releases/download/v0.1.0/ct_0.1.0_linux_amd64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_AMD64_SHA256"
    end
  end

  def install
    bin.install "ct"
  end

  test do
    system "#{bin}/ct", "--help"
  end
end
EOF
```

## Step 3: Using GoReleaser (Recommended)

The project includes a `.goreleaser.yml` configuration file. To use it:

1. Install GoReleaser: `brew install goreleaser`
2. Create a GitHub Personal Access Token (with `repo` and `workflow` scopes)
3. Export the token: `export GITHUB_TOKEN=your_token_here`
4. Tag a release: `git tag v0.1.0 && git push origin v0.1.0`
5. Run GoReleaser: `goreleaser release`

This will:
- Build binaries for multiple platforms
- Create a GitHub Release with the binaries
- Automatically update your Homebrew tap formula

## Step 4: Update GoReleaser Config

Update `.goreleaser.yml` with your GitHub username:

```yaml
brews:
  - repository:
      owner: YOUR_GITHUB_USERNAME
      name: homebrew-tap
```

## Step 5: Verify Installation

Once everything is set up, users can install with:

```bash
brew tap yourusername/tap
brew install ct
```

## For More Information

- [Homebrew Documentation](https://docs.brew.sh/How-to-Create-and-Maintain-a-Tap)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GoReleaser Homebrew Support](https://goreleaser.com/customization/homebrew/)
