# GoTK - Modern Go Toolkit for Rapid Service Development

🚀 Accelerate backend development with modular Golang components.

## Features

- 🛠️ Comprehensive toolkit for Go service development
- 🔧 Modular design with independent components
- 🚀 High performance and low latency
- 🔒 Built-in security features
- 📊 Monitoring and metrics support
- 🔄 Easy integration with existing systems

## Installation

Install the GoTK CLI tool and framework:

```bash
# Install the CLI tool
go install github.com/sunshinexcode/gotk/vcli@latest

# Add GoTK to your project
go get github.com/sunshinexcode/gotk
```

## Quick Start

Create and run a new HTTP service project:

```bash
# Create a new HTTP service project
vcli init test-http-example

# Navigate to the project and run it
cd ./test-http-example/http
make go    # Install dependencies
make run   # Start the service
```

The service will be available at `http://localhost:8080`.

## Modules

GoTK consists of the following modules:

| Module      | Description                      |
| ----------- | -------------------------------- |
| vaes        | AES encryption utilities         |
| valarm      | Alarm and notification system    |
| vapi        | API development tools            |
| vauth       | Authentication and authorization |
| vbase64     | Base64 encoding/decoding         |
| vbootstrap  | Application bootstrap helpers    |
| vcache      | Caching utilities                |
| vcli        | Command-line interface tools     |
| vcode       | Error code management            |
| vconfig     | Configuration management         |
| vcontroller | MVC controller helpers           |
| vconv       | Data conversion utilities        |
| vcron       | Cron job scheduling              |
| vdebug      | Debugging utilities              |
| venv        | Environment variable management  |
| verror      | Error handling utilities         |
| ves         | Elasticsearch integration        |
| vfile       | File operations utilities        |
| vfx         | Effects and animations           |
| vhmac       | HMAC authentication              |
| vhttp       | HTTP client and server           |
| vhttp       | HTTP server and client           |
| vjson       | JSON processing utilities        |
| vlimit      | Rate limiting utilities          |
| vlog        | Logging system                   |
| vmap        | Map operations utilities         |
| vmask       | Mask operations utilities        |
| vmd5        | MD5 hashing utilities            |
| vmetric     | Metrics collection               |
| vmiddleware | HTTP middleware                  |
| vmock       | Mocking utilities                |
| vmodel      | Data model helpers               |
| vmongodb    | MongoDB integration              |
| vmysql      | MySQL database integration       |
| voutput     | Output formatting utilities      |
| vpb         | Protocol Buffers utilities       |
| vpprof      | Performance profiling            |
| vqueue      | Queue management utilities       |
| vrand       | Random number generation         |
| vredis      | Redis integration                |
| vreflect    | Reflection utilities             |
| vreq        | HTTP request utilities           |
| vsafe       | Thread-safe utilities            |
| vshell      | Shell command execution          |
| vstr        | String manipulation utilities    |
| vstruct     | Struct operations utilities      |
| vtcp        | TCP networking utilities         |
| vtest       | Testing utilities                |
| vtime       | Time manipulation utilities      |
| vtrace      | Tracing utilities                |
| vuuid       | UUID generation utilities        |
| vvalid      | Validation tools                 |
| vvar        | Variable management utilities    |
| vversion    | Version management utilities     |
| vwebsocket  | WebSocket utilities              |

## Examples

Check out the [examples](./examples) directory for usage examples:

- HTTP server example
- Database integration examples
- Authentication examples
- And more...

## License

`GoTK` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.
