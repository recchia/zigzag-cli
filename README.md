# ZigZag CLI

ZigZag is a command-line tool designed to list and download builds from Expo Application Services (EAS). It provides an interactive interface to select a build, specify a filename, and automatically organizes downloads into timestamped directories.

## Features

- **Interactive Build Selection**: Fetches and lists the latest 20 builds from EAS for selection.
- **Smart File Naming**: Provides sensible default filenames based on platform and build profile.
- **Timestamped Directories**: Automatically creates a subdirectory named with the current date and time (`YYYY-MM-DD_HH-MM-SS`) for each download to prevent file collisions.
- **Customizable Output**: Prompts for the base output directory and target filename.

## Project Structure

```text
.
├── cmd/
│   └── zigzag/
│       └── main.go         # CLI entry point (Cobra commands)
├── internal/
│   ├── downloader/
│   │   ├── downloader.go    # HTTP download logic and directory management
│   │   └── downloader_test.go
│   ├── eas/
│   │   ├── eas.go           # EAS CLI interaction and JSON parsing
│   │   └── eas_test.go
│   └── ui/
│       └── ui.go            # Interactive prompts using Survey
├── Makefile                 # Build and test automation
├── go.mod                   # Project dependencies
└── go.sum
```

## Prerequisites

- **Go 1.26+**: The project is built using modern Go 1.26 features.
- **EAS CLI**: You must have `eas-cli` installed and be logged in to an Expo account.
  ```bash
  npm install -g eas-cli
  eas login
  ```
- **Active Expo Project**: The tool must be run from within an initialized Expo project directory (where `app.json` or `eas.json` exists).

## Installation & Building

### Using Makefile (Recommended)

To build the project and create the executable:

```bash
make build
```

This will create a `build/` directory and place the `zigzag` binary inside it.

### Manual Build

```bash
mkdir -p build
go build -o build/zigzag ./cmd/zigzag/main.go
```

## Usage

### Running the Binary

Run the built binary from your Expo project root:

```bash
./build/zigzag
```

### Steps:
1. **Fetching Builds**: The tool runs `eas build:list --json` to fetch recent builds.
2. **Selection**: Select the desired build from the list.
3. **Output Directory**: Enter the base directory where you want to save the build (default is current directory `.`).
4. **Filename**: Enter the desired filename (a default is provided based on the build metadata).
5. **Download**: The tool downloads the build into a timestamped subfolder within your chosen directory.

### CLI Help

```bash
./build/zigzag --help
```

## Development

### Running Tests

To run the unit tests for the project:

```bash
make test
```

### Cleaning Build Artifacts

To remove the `build/` directory:

```bash
make clean
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra): CLI framework.
- [Survey](https://github.com/AlecAivazis/survey): Interactive terminal prompts.
- [Go-Dotenv](https://github.com/joho/godotenv): (Optional) For environment variable management.
