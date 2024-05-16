Class Thothica < Formula
  desc "A simple cli tool to manage semantic search"
  homepage "https://github.com/Thothica/thothica"
  version "v0.1.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/Thothica/thothica/releases/download/v0.1.0/thothica_Darwin_arm64.tar.gz"
      sha256 "6c03c98170b4751e7307b84c07982766f92eed339cb25aaebb333c78ba57bf9a"
    end
    on_intel do
      url "https://github.com/Thothica/thothica/releases/download/v0.1.0/thothica_Darwin_x86_64.tar.gz"
      sha256 "1ec788ad57eb0b2fab3f0ab09581309a4f71c45db19b42cbbe636d9d5a41d112"
    end
  end

  on_linux do
    on_intel do
      if !Hardware::CPU.is_64_bit?
        url "https://github.com/Thothica/thothica/releases/download/v0.1.0/thothica_Linux_i386.tar.gz"
        sha256 "546952ea0a8977af43a038ba0e0996346d17093f77055bf71304f04c2313d36e"
      end
    end
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/Thothica/thothica/releases/download/v0.1.0/thothica_Linux_x86_64.tar.gz"
        sha256 "dc82ad058483926d794901d2a3156f9cf3924b8382686acbf67c08829dcce7ac"
      end
    end
    on_arm do
        url "https://github.com/Thothica/thothica/releases/download/v0.1.0/thothica_Linux_arm64.tar.gz"
      sha256 "311193cead25d8008ea200d0fe3b9a17261b1679c78e6922dc6609cc011e567c"
    end
  end

  def install
    bin.install "thothica"
  end

  def post_install
    thothica_file = File.join(Dir.home, ".thothica")
    unless File.exist?(thothica_file)
      FileUtils.touch(thothica_file)
    end
  end
end
