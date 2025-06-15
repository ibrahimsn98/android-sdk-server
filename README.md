# WORK IN PROGRESS

## Android SDK Command-line Tools HTTP Bridge

Android SDK Command-line Tools HTTP Bridge is an open-source HTTP server that acts as a bridge to the Android SDK's core command-line utilities — including sdkmanager, avdmanager, adb, and emulator.

This tool simplifies remote and programmatic access to Android SDK functionality by exposing these tools over a RESTful HTTP API. Ideal for use in CI/CD environments, cloud-based development, or remote emulation management.


- 📦 SDK Manager Support — Install, update, and manage Android SDK packages via HTTP
- 👤 AVD Manager — Create and manage Android Virtual Devices (AVDs)
- 📱 ADB Integration — Run ADB commands remotely (e.g., install APKs, logcat, shell commands)
- 🕹️ Emulator Support — Start, stop, and interact with emulators over HTTP
- 🛠️ Runs as Daemon — Works as a background service on Linux or macOS
