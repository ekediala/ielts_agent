# 🎯 IELTS Agent

> An intelligent AI-powered assistant designed to help you excel in the IELTS exam through personalized feedback and guidance.

[![Go Version](https://img.shields.io/badge/Go-1.24.1+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Claude AI](https://img.shields.io/badge/Powered%20by-Claude%20AI-FF6B35?style=flat)](https://www.anthropic.com/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## ✨ Features

- 🤖 **AI-Powered Coaching**: Leverages Claude AI for intelligent IELTS preparation assistance
- 📝 **Real-time Feedback**: Get instant, actionable feedback on your writing and speaking practice
- 📄 **PDF Document Analysis**: Extract and analyze text from IELTS practice materials
- 🎯 **Focused Learning**: Specialized exclusively for IELTS exam preparation
- 💬 **Interactive Chat Interface**: Natural conversation flow for personalized learning
- 🔧 **Extensible Tool System**: Built-in file operations and content analysis tools

## 🚀 Quick Start

### Prerequisites

- Go 1.24.1 or higher
- Anthropic API key (Claude AI)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/ekediala/ielts_agent.git
   cd ielts_agent
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env and add your Anthropic API key
   echo "ANTHROPIC_API_KEY=your_api_key_here" > .env
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Build and run**
   ```bash
   make start
   ```

   Or manually:
   ```bash
   go build -o bin ./...
   ./bin
   ```

## 💡 Usage

Once started, the IELTS Agent provides an interactive chat interface where you can:

- **Get writing feedback**: Share your IELTS essays for detailed analysis
- **Practice speaking**: Receive guidance on speaking test responses
- **Analyze materials**: Upload PDF handouts and practice materials for study
- **Ask questions**: Get expert advice on IELTS strategies and techniques

### Example Interaction

```
You: Can you help me improve my Task 2 essay about technology in education?

Claude: I'd be happy to help you improve your IELTS Task 2 essay! Please share your essay, and I'll provide detailed feedback on:
- Task response and argument development
- Coherence and cohesion
- Lexical resource and vocabulary usage
- Grammatical range and accuracy

You can either type your essay directly or use a file if you have it saved.
```

## 🏗️ Architecture

The IELTS Agent follows a clean, modular architecture:

```
ielts_agent/
├── main.go           # Application entry point
├── agent.go          # Core agent logic and conversation handling
├── tools.go          # Tool definition framework
├── read_file.go      # File reading with PDF support
├── edit_file.go      # File editing capabilities
├── list_files.go     # Directory listing functionality
├── rules.xml         # Agent behavior rules and constraints
└── handout.pdf       # Example IELTS materials
```

### Core Components

- **Agent Engine**: Manages conversation flow and Claude AI integration
- **Tool System**: Extensible framework for adding new capabilities
- **File Operations**: Handle various file types including PDFs
- **Safety Rules**: Ensures the agent stays focused on IELTS preparation

## 🛠️ Available Tools

The agent comes with built-in tools for enhanced functionality:

| Tool | Description | Usage |
|------|-------------|-------|
| `read_file` | Read and analyze files (including PDFs) | Extract text from practice materials |
| `edit_file` | Create and modify text files | Save practice essays and notes |
| `list_files` | Browse directory contents | Organize study materials |

## 📋 Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `ANTHROPIC_API_KEY` | Your Claude AI API key | Yes |

### Rules Configuration

The agent's behavior is governed by `rules.xml`, which defines:

- **Language Settings**: Default to English, adapt to user preferences
- **Task Focus**: Strictly IELTS-related assistance only
- **Safety Guidelines**: Privacy protection and ethical AI usage
- **Response Format**: Professional, educational tone

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

For specific components:

```bash
go test -v ./read_file_test.go
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Add tests for new functionality
- Update documentation for any API changes
- Ensure the agent remains focused on IELTS preparation

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [IELTS Official Website](https://www.ielts.org/)
- [Claude AI Documentation](https://docs.anthropic.com/)
- [Go Documentation](https://golang.org/doc/)

## 🆘 Support

- **Issues**: Report bugs or request features via [GitHub Issues](https://github.com/ekediala/ielts_agent/issues)
- **Discussions**: Join the conversation in [GitHub Discussions](https://github.com/ekediala/ielts_agent/discussions)

## 🙏 Acknowledgments

- **Anthropic** for providing the Claude AI API
- **IELTS Community** for inspiration and feedback
- **Go Community** for excellent tooling and libraries

---

<div align="center">
<strong>Ready to ace your IELTS exam? Start chatting with your AI coach today! 🎓</strong>
</div>