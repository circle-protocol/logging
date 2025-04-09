# logging

[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Release](https://github.com/Circle-Protocol/logging/workflows/Release/badge.svg)](https://github.com/Circle-Protocol/logging/actions)
[![license](https://badgen.net/github/license/Circle-Protocol/logging/)](https://github.com/Circle-Protocol/logging/blob/production/LICENSE)
[![release](https://badgen.net/github/release/Circle-Protocol/logging/stable)](https://github.com/Circle-Protocol/logging/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/Circle-Protocol/logging)](https://goreportcard.com/report/github.com/Circle-Protocol/logging)

## What Is It

A modern, structured logging library for Go applications built on top of the standard library's `log/slog` package. This library provides a flexible and powerful logging solution with support for:

- Structured logging with key-value pairs
- Context-aware logging
- HTTP request and response logging middleware
- Conditional logging with OnError
- Multiple output formats (JSON, text)
- Custom field mapping
- Caller information

## Getting Started

```go
package main

import (
    "github.com/Circle-Protocol/logging"
)

func main() {
    // Simple logging
    logging.Info("Hello, world!")
    
    // Structured logging with fields
    logging.WithFields("user", "john", "action", "login").Info("User logged in")
    
    // Conditional logging (only logs if err is not nil)
    err := someOperation()
    logging.OnError(err).Error("Operation failed")
}
```

## Supported Go Versions

For security reasons, we only support and recommend the use of one of the latest two Go versions (✅).

| Version | Supported          |
|---------|--------------------|
| <1.23   | ❌                |
| 1.23    | ✅                |
| 1.24    | ✅                |

## License

This library is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
