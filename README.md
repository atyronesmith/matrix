# Matrix Visual Effect

A Go application using the Fyne library that creates a classic Matrix movie-style falling green character effect.

## Features

- Falling green characters in multiple columns
- Varying fall speeds and column heights
- Bright white characters at the head of each stream
- Fading green characters in the trail
- Random character changes
- **LLM names fall vertically as part of the Matrix character streams**
- Classic Matrix-style visual effects with tech company references

## Prerequisites

- Go 1.21 or later
- The Fyne library (automatically downloaded via go mod)

## Installation & Running

1. Clone or download this project
2. Navigate to the project directory
3. Install dependencies:

```bash
make deps
```

4. Run the application using one of these methods:

**Using Makefile (recommended):**
```bash
# Build and run
make quick

# Or run directly
make run

# Or build first, then run
make build
./matrix
```

**Using Go directly:**
```bash
go run main.go
```

## Controls

- The window will open showing the Matrix effect
- Close the window or press Ctrl+C to exit

## Customization

You can modify the following constants in `main.go` to customize the effect:

- `windowWidth` and `windowHeight`: Window dimensions (default: 1000×600)
- `charWidth` and `charHeight`: Character spacing (default: 12×18)
- `charSet`: Characters used in the effect
- Animation speed: Modify the ticker duration in the main function (default: 100ms)
- LLM name falling: Names appear every 0.5-2 seconds and fall vertically in bright yellow-white
- LLM names: 25 different AI model and company names randomly selected

## Makefile Commands

The project includes a comprehensive Makefile with the following targets:

- `make build` - Build the Matrix application
- `make run` - Run the Matrix application
- `make quick` - Build and run in one command
- `make clean` - Clean build artifacts and binaries
- `make deps` - Download and install dependencies
- `make build-all` - Build for multiple platforms (Linux, macOS, Windows)
- `make release` - Build optimized release version
- `make dev` - Build development version with debugging
- `make fmt` - Format Go source code
- `make lint` - Lint Go source code
- `make info` - Show project information
- `make help` - Show all available targets

**Examples:**
```bash
make help       # Show all available commands
make build      # Compile the application
make run        # Run the Matrix effect
make build-all  # Create binaries for all platforms
make clean      # Clean up build files
```

## Dependencies

- [Fyne v2](https://fyne.io/) - Cross-platform GUI toolkit for Go