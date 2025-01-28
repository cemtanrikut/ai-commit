# AI-Commit

AI-Commit is a command-line tool designed to help developers generate meaningful Git commit messages using OpenAI's GPT models. It automatically analyzes your staged changes and suggests a commit message based on the context of the changes.

## Features

- Automatically fetches `git diff` for staged changes.
- Integrates with OpenAI API to generate meaningful commit messages.
- Supports Conventional Commit types like `feat`, `fix`, `docs`, and more.
- Allows manual editing of commit messages before finalizing.
- Includes test coverage for critical components.

---

## Installation

### Prerequisites

1. **Git**: Ensure you have Git installed on your machine.
   ```bash
   git --version
   ```

2. **Go**: Install Go (1.18 or later)
   ```bash
   go version
   ```

3. **OpenAI API Key**: Obtain an API key from OpenAI's platform.

### Clone the Repository

   ```bash
   git clone https://github.com/yourusername/ai-commit.git
   cd ai-commit
   ```

## Setup

### Configure Environment Variables

Set the OPENAI_API_KEY environment variable to your OpenAI API key.

**Linux / macOS**:

```bash
export OPENAI_API_KEY=your_api_key
```

**Windows (PowerShell)**:

```bash
$env:OPENAI_API_KEY="your_api_key"
```

## Usage

### Generate a Commit Message

1. Stage your changes:

``` bash
git add yourfile.go
```

2. Run the tool:

```bash
go run main.go
```

3. View the generated commit message:

```bash
Generated Commit Message:
feat: Add feature X

This commit introduces feature X with enhancements to Y and Z.
```

4. Approve or edit the message:
- If approved, the message will be committed automatically.
- If not approved, you can edit the message before committing.

## Testing

### Run Unit Tests

To execute all tests, run:

```bash
go test ./... -v
```

### Example Output

```bash
=== RUN   TestGenerateCommitMessage_Success
--- PASS: TestGenerateCommitMessage_Success (0.01s)
=== RUN   TestGenerateCommitMessage_Error
--- PASS: TestGenerateCommitMessage_Error (0.01s)
PASS
ok  	ai-commit/api	0.02s
```

## Future Improvements
- Authentication System: Add a credit/token system for using the OpenAI API.
- Customizable Templates: Allow users to define their own commit message templates.
- Support for Multiple Files: Generate commit messages for multiple staged files individually.

## Contributing

Contributions are welcome! Follow these steps to contribute:
1. Fork the repository.
2. Create a new branch:

```bash
git checkout -b feature/your-feaute-name
```

3. Commit your changes:

```bash
git commit -m "feat: Add your feature"

4. Push your branch:

```bash
git push origin feature/your-feature-name
```

5. Open a pull request

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- OpenAI for providing powerful language models.
- The Go programming language for its simplicity and performance.