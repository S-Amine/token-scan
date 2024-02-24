# Token Scan Prototype

Token-Scan is a prototype decentralized exchange (DEX) ERC20 Token scanner tailored for the Ethereum blockchain. This tool aims to simplify the complexities of navigating the decentralized ecosystem by offering users efficient scanning capabilities across multiple data sources. Please note that this version of the tool is a prototype.

## Introduction

Token-Scan prototype provides a fundamental functionality for scanning ERC20 tokens. It's built with modularity in mind, facilitating seamless integration of new scanning methods and offering insights into tokens listed on decentralized exchanges.

## Problem Statement

Decentralized exchanges (DEXs) have transformed the cryptocurrency landscape, providing users with the freedom to trade assets without reliance on traditional centralized intermediaries. However, the vast array of tokens listed on DEX platforms can present challenges. Identifying legitimate tokens, detecting potential scams or honeypots, and obtaining comprehensive intelligence on tokens of interest are essential tasks for developers and investors alike. Token-Scan prototype endeavors to tackle these challenges by providing a basic scanning tool capable of analyzing ERC20 tokens across various data sources.

## Features

- **Multi-Source Scanning**: Token-Scan prototype aggregates data from various sources to provide basic insights into ERC20 tokens.
- **Modular Architecture**: The project is structured with modularity in mind, allowing for easy integration of new scanning methods.
- **CLI Tool**: Token-Scan prototype offers a basic command-line interface (CLI) for convenient and straightforward usage.
- **GoLang Package**: Users can integrate Token-Scan prototype as a GoLang package to leverage its basic functionality within their projects.

## Installation

To utilize Token-Scan prototype, follow the installation instructions below based on your preferred method:

### CLI Usage & Installation (for linux)

1. Clone the repository:

```
git clone https://github.com/s-Amine/token-scan.git
```

2. Navigate to the project directory:

```
cd token-scan
```

3. Run the executable:

```
./token-scan -mode <mode> -token <token_hash>
```

Replace `<mode>` with the desired scanning mode (`multiscan`, `goplus`, `ishoneypot`, or `quickIntel`) and `<token_hash>` with the hash of the token you wish to scan.


### GoLang Package Integration

1. Install the package:

```
go get github.com/s-Amine/token-scan
```

2. Import the package in your GoLang code:

```go
import "github.com/s-Amine/token-scan/scanners/<mode>"
```

Replace `<mode>` with the desired scanning mode (`multiscan`, `goplus`, `ishoneypot`, or `quickIntel`)

3. Utilize the package to instantiate structs and access scanning methods as needed.



### GoLang Package Usage

#### Multiscan Scan Usage
```go
import "github.com/s-Amine/token-scan/scanners/multiscan"

...

result, err := multiscan.Scan("<token_hash>")
if err != nil {
    fmt.Println("Error occurred during scan:", err)
    return
}
fmt.Println(result)
```

#### Quickintel Scan Usage

```go
import "github.com/s-Amine/token-scan/scanners/quickintel"

...

result, err := quickintel.Scan("<token_hash>")
if err != nil {
    fmt.Println("Error occurred during scan:", err)
    return
}
fmt.Println(result)
```

#### Is Honeypot Scan Usage

```go
import "github.com/s-Amine/token-scan/scanners/ishoneypot"

...

result, err := ishoneypot.Scan("<token_hash>")
if err != nil {
    fmt.Println("Error occurred during scan:", err)
    return
}
fmt.Println(result)
```
#### Goplus Scan Usage

```go
import "github.com/s-Amine/token-scan/scanners/goplus"

...

result, err := goplus.Scan("<token_hash>")
if err != nil {
    fmt.Println("Error occurred during scan:", err)
    return
}
fmt.Println(result)
```

## Project Structure

```
token-scan/
├── go.mod
├── go.sum
├── main.go
├── scanners/
│   ├── goplus/
│   │   └── scan.go
│   ├── ishoneypot/
│   │   └── scan.go
│   ├── multiscan/
│   │   └── scan.go
│   └── quickintel/
│       └── scan.go
└── token/
    └── model.go
```

- **go.mod, go.sum**: Go module files managing dependencies.
- **main.go**: Entry point of the Token-Scan CLI tool.
- **scanners/**: Directory containing modules for different scanning methods.
- **token/**: Directory containing token-related models.

## Contributing

Contributions to Token-Scan are welcome! If you have ideas for improvements, new features, or bug fixes, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. 

## Acknowledgments

Token-Scan's functionality is indebted to various open-source libraries and data sources. We extend our gratitude to the developers and contributors of the following projects and platforms:

- The GoPlusSecurity team for their exceptional work on the [goplus-sdk-go](https://github.com/GoPlusSecurity/goplus-sdk-go/) library.
- The team behind [Honeypot API](https://api.honeypot.is/) for providing valuable data sources integral to Token-Scan's capabilities.
- [QuickIntel](https://app.quickintel.io/api/) for their API services, which significantly enhance Token-Scan's functionality.
- The creators and maintainers of the Go programming language, whose innovative language features and robust ecosystem have been instrumental in the development of Token-Scan.



---
**Disclaimer**: This project is provided as-is with no warranties or guarantees. Use at your own discretion.
