# fmoj

A simple concurrent job runner written in Go. It loads job definitions from a `jobs.json` file and executes them periodically using goroutines and context-based cancellation.

---

## 🚀 Features

- Load jobs from a JSON file
- Run commands concurrently
- Context-based cancellation
- Easy to extend with custom logic

---

## 📦 Installation

Download the latest binary from [Releases](https://github.com/fmo/jobs/releases) and make it executable:

```bash
chmod +x fmoj_v1.0.0_linux_arm64

./fmoj_v1.0.0_linux_arm64
```

Or build it manually:

`go build -o fmoj .`

## 🛠 Usage

After downloading the binary, you'll need to also copy the jobs.json file and the scripts/ directory into the same directory where the binary will run.
