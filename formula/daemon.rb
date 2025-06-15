class AndroidSdkServer < Formula
  desc "My always-on macOS background service"
  homepage "https://github.com/ibrahimsn98/android-sdk-server"
  url "https://github.com/ibrahimsn98/android-sdk-server/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "<sha256sum>"
  license "Apache"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "bin/android-sdk-server", *std_go_args, "./cmd"
  end

  plist_options manual: "android-sdk-server"

  def plist
    <<~EOS
      <?xml version="1.0" encoding="UTF-8"?>
      <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
       "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
      <plist version="1.0">
      <dict>
          <key>Label</key>
          <string>com.android.sdkserver</string>

          <key>ProgramArguments</key>
          <array>
              <string>#{opt_bin}/android-sdk-service</string>
          </array>

          <key>RunAtLoad</key>
          <true/>

          <key>KeepAlive</key>
          <true/>

          <key>StandardOutPath</key>
          <string>/tmp/android-sdk-service.log</string>

          <key>StandardErrorPath</key>
          <string>/tmp/android-sdk-service.err.log</string>
      </dict>
      </plist>

    EOS
  end

  test do
    system "#{bin}/android-sdk-server", "--version"
  end
end