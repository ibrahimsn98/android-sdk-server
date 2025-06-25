class AndroidSdkServer < Formula
  desc "Service to manage Android SDK tasks over a server"
  homepage "https://github.com/ibrahimsn98/android-sdk-server"
  url "https://github.com/ibrahimsn98/android-sdk-server/archive/refs/tags/1.0.0.tar.gz"
  sha256 "8f42e29fedca2488c202cfdd3e029604203489ad1bbc71627018871ed2b335bf"
  license "Apache"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "bin/android-sdk-server", *std_go_args, "./cmd"
  end

  service do
    run opt_bin/"android-sdk-server"
    keep_alive true
    working_dir var
    log_path var/"log/android-sdk-server.log"
    error_log_path var/"log/android-sdk-server.log"
  end

end
