class Qx < Formula
  desc "{{ .Description }}"
  homepage "{{ .Homepage }}"
  version "{{ .Version }}"
  license "{{ .License }}"

  {{range .Binaries}}
  url "{{.URL}}"
  sha256 "{{.SHA256}}"
  {{end}}

  def install
    bin.install "qx"
  end

  test do
    # Ensure the binary runs and prints the version string
    assert_match version.to_s, shell_output("#{bin}/qx --version")
  end
end