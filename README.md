# 🌐 Websites Health Checker

A simple and efficient CLI tool to monitor website availability and HTTP status in real time.

Designed for developers who want a lightweight way to track uptime, log responses, and manage a custom list of websites directly from the terminal.

---

## 🚀 Features

- Add websites to monitoring list
- List all registered websites
- Remove websites by ID
- Perform HTTP health checks
- Log status history locally
- Cross-platform support (Windows, Linux, macOS)

---

## 🧠 How it works

The tool reads a `websites.txt` file stored in your user home directory:

```
C:\Users\<user>\websites.txt
```

or

```
/home/<user>/websites.txt
```

Each line represents a website to be monitored.

Logs are stored in:

```
websites.txt (input)    logs.txt (output)
```

---

## 📦 Installation

### Option 1 — Download release

Go to the latest release:

👉 [https://github.com/danilowskii/websites-health-checker/releases](https://github.com/danilowskii/websites-health-checker/releases)

Download the binary for your system:

| Platform | Binary          |
| -------- | --------------- |
| Windows  | `checker.exe`   |
| Linux    | `checker-linux` |
| macOS    | `checker-mac`   |

---

## ▶️ Usage

Run the executable:

### Windows

```bash
checker.exe
```

### Linux / macOS

```bash
./checker
```

---

## 📋 Menu

When running the CLI:

```
1 - Check websites
2 - View logs
3 - Add website
4 - List websites
5 - Delete website
0 - Exit
```

---

## 📝 Example workflow

**Add a website:**

```
Enter a website: https://google.com
```

**Run health check:**

```
1 - Monitoring your websites...
```

**View logs:**

```
2 - Checking your last log...
```

---

## 📊 Status rules

| Status Code | Indicator | Result                                            |
| ----------- | --------- | ------------------------------------------------- |
| 2xx         | 🟢        | HEALTHY                                           |
| 4xx / 5xx   | 🟡        | UNHEALTHY                                         |
| 999         | ⚪        | UNAUTHORIZED / BLOCKED (e.g. LinkedIn protection) |

---

## ⚙️ Tech Stack

- **Go (Golang)**
- `net/http`
- File I/O (`os`, `bufio`)
- ANSI terminal colors

---

## 📁 Project Structure

```
websites-health-checker/
├── main.go
├── websites/
│   └── websites.go
├── releases/
└── README.md
```

---

## 🧩 Design decisions

- File-based storage for simplicity
- No external dependencies
- User home directory for portability
- CLI-first experience

---

## 🧠 Future improvements

- [ ] Global CLI install (`checker` command)
- [ ] Config file support (`.yaml` / `.json`)
- [ ] Parallel health checks
- [ ] Retry system for unstable websites
- [ ] Better CLI UX (colors + loading states)

---

## 👨‍💻 Author

Built by **Danilo Paiva**  
LinkedIn: [https://www.linkedin.com/in/paivadanilo/](https://www.linkedin.com/in/paivadanilo/)

---

## 📄 License

This project is open-source for learning and personal use.
