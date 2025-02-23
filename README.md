# Build a Simple Full-stack App with Golang in 2 Hours!

This repository contains the code for the ["Build a Simple Full-stack App with Golang in 2 Hours!"](https://youtu.be/bem0bdDBs8A) tutorial, showcasing how to build a social interaction app called **GoCards** using the **GOAT Stack** (Go, a-h/templ, Alpine.js, Tailwind CSS) and **PocketBase** as a Backend-as-a-Service (BaaS). Users can sign up, create cards, like cards, and view/read cards created by others. This project is designed for learning and rapid prototyping, *not* for production-ready applications.

**Project Overview:**

GoCards leverages the speed and efficiency of Go for server-side rendering with `a-h/templ`, the dynamic capabilities of Alpine.js for user interactions, the styling power of Tailwind CSS, and the simplicity of PocketBase for user authentication, data storage, and real-time capabilities.

* **Backend:** Go (with server-side rendering using `a-h/templ`), PocketBase v0.22.23 (BaaS)
* **Frontend:** Alpine.js (for dynamic interactions)
* **Styling:** Tailwind CSS
* **Database:** PocketBase (handles user accounts, card data, and likes)
* **Editor:** Neovim (for development)

**Important Version Note:**
This project uses PocketBase v0.22.23. As PocketBase is still pre-1.0 and undergoes frequent breaking changes, newer versions may cause compatibility issues. Please stick to version 0.22.23 when following this tutorial:

```go
require github.com/pocketbase/pocketbase v0.22.23
```

**Important Disclaimer:**

This project uses a simplified architecture for educational purposes and rapid prototyping. It is *not* suitable for large, complex, or production-ready applications. For those, you'll need a more robust architecture, potentially involving custom backend logic, more sophisticated data modeling, and considerations for scaling and security that go beyond the scope of this tutorial.

## Getting Started

### Prerequisites

* [Go](https://go.dev/dl/) (version 1.20 or later recommended)
* [Bun](https://bun.sh/) (for managing JavaScript dependencies)
* [Neovim](https://neovim.io/) (our chosen editor)
* [Git](https://git-scm.com/downloads)
* [fswatch](https://github.com/emcrisostomo/fswatch)
* [templ](https://github.com/a-h/templ)

### Usage

1. First, create a new directory and clone the Makefile:

```bash
mkdir my-gocards
cd my-gocards
curl -O https://raw.githubusercontent.com/morethancoder/gocards/main/Makefile
```

2. View available commands:

```bash
make help
```

This will show all available make commands:
* `make init` → Initialize project
* `make build` → Build project
* `make css` → Generate CSS
* `make serve` → Start server
* `make watch` → Watch for changes
* `make clean` → Clean build files
* `make source` → Load environment variables
* `make sync` → Run browser-sync
* `make test` → Run tests

3. Initialize the project:

```bash
make init
```

The `make init` command will:
* Check all required dependencies
* Create the project directory structure
* Initialize the Go module
* Set up Tailwind CSS configuration
* Download Alpine.js

### Project Structure

```
.
├── LICENSE
├── Makefile
├── build/                 # Build outputs
├── cmd/
│   └── GoCards/
│       └── main.go       # Main application entry
├── internal/             # Internal packages
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── service/
│   └── util/
├── static/              # Static assets
│   ├── css/
│   │   ├── input.css
│   │   └── styles.css
│   └── js/
│       ├── alpine.min.js
│       └── pocketbase.min.js
├── views/               # templ templates
│   ├── components/
│   │   ├── Icons.templ
│   │   └── Navbar.templ
│   ├── layouts/
│   │   └── BaseLayout.templ
│   └── pages/
│       ├── Dashboard.templ
│       ├── HomePage.templ
│       ├── Login.templ
│       └── Signup.templ
└── pb_data/             # PocketBase data directory
    ├── data.db
    ├── logs.db
    └── storage/
```

### Running the Application

1. Start all services:

```bash
make serve
```

This command:
* Generates templ files
* Compiles CSS with Tailwind
* Runs the Go server

2. Watch for changes during development:

```bash
make watch
```

3. Access the applications:
* Main app: http://localhost:8090
* PocketBase admin: http://localhost:8090/_/

### Development Workflow

1. Start the development server with file watching:
```bash
make watch
```

2. Make changes in Neovim
3. The server will automatically rebuild when:
   * `.templ` files are modified
   * Go files are changed
   * CSS is updated

4. Run browser-sync for live reload (optional):
```bash
make sync
```

### Additional Commands

* `make css` - Regenerate Tailwind CSS
* `make clean` - Remove built artifacts and generated files
* `make test` - Run the test suite
* `make source` - Load environment variables from .env
* `make build` - Create a production build

## Related Resources

* [Tutorial Video Part 1](https://youtu.be/bem0bdDBs8A)
* More videos coming soon!

## Connect With Me

* Telegram Channel: [AliTahseenDev](https://t.me/+-aFYIKA-ZGU4M2Uy)
* Website: [MoreThanCoder](https://morethancoder.com)

## License

[MIT License](LICENSE)
