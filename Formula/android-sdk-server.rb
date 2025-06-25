class AndroidSdkServer < Formula
  desc "Service to manage Android SDK tasks over a server"
  homepage "https://github.com/ibrahimsn98/android-sdk-server"
  url "https://github.com/ibrahimsn98/android-sdk-server/archive/refs/tags/1.0.0.tar.gz"
  sha256 "197c57bda6a62c8e4b9de5cbd679da16d8ade63f3ad378231d186278e01cf03f"
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
