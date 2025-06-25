class AndroidSdkServer < Formula
  desc "Service to manage Android SDK tasks over a server"
  homepage "https://github.com/ibrahimsn98/android-sdk-server"
  url "https://github.com/ibrahimsn98/android-sdk-server/archive/refs/tags/1.0.0.tar.gz"
  sha256 "63c70e7a05ed3aafc5a65d88a7893ab8a2da372b2504542a1546901d194077ba"
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
