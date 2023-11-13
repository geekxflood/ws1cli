# WS1CLI

`ws1cli` is a command-line interface (CLI) tool built in Go and leveraging the Cobra library, designed for interacting with VMware Workspace ONE UEM API. It enables administrators to manage and automate tasks within the Workspace ONE platform efficiently.

## Features

- Simplified Command Set: Easy-to-use commands for common administration tasks.
- Secure Configuration: Securely stores API credentials using advanced encryption.
- Go and Cobra: Built with Go for reliability and Cobra for a powerful command line interface.
- Commands implemented:
  - Test: Verify configuration and connectivity to the API.
  - Devices: Interact with device API.
    - Inventory: Output an array of JSON of devices in LGID.

## TODO

- [ ] Product List Interaction: Manage and list products within your Workspace ONE environment.
- [ ] Fetch Tag List: Retrieve a list of tags associated with devices or users.
- [ ] Device Management: Perform operations such as device lock, wipe, and query information.
- [ ] Comprehensive Reporting: Generate reports on user access, device compliance, and application usage.

## Environement Variables

`ws1cli` need the following environment variables:

| Variable | Description |
| --- | --- |
| `WS1_KEY` | Encryption key |

## Installation

Ensure Golang is installed on your system before installing `ws1cli`.

```bash
git clone https://github.com/christopherime/WS1CLI.git
cd ws1cli
go build
```

Add `ws1cli` to your system's PATH to use it from any directory.

## Configuration

Initialize your configuration with the `init` command:

```bash
ws1cli init
```

You'll be prompted to enter your API URL, username, and password.

## Commands

Here's how to get started with some basic commands:

- **Global flags**

| Flag | Description |
| --- | --- |
| `-h` `--help` | Help info |
| `-i` `--insecure` | Ignore TLS verification |
| `-p` `--pretty` | Pretty-print JSON output |

### Test

To verify API configuration and connectivity:

```bash
ws1cli test
```

### Devices

To interact with device API:

```bash
ws1cli device
```

| Flag | Description |
| --- | --- |
| `-d` `--inventory` | Output an array of JSON of devices in LGID |
| `-l` `--lgid` | LGID of device, mandatory when using `-d` flag |

## Contributing

Contributions are welcome. Fork the project, make your updates, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support

For support, please open an issue in the GitHub repository issue tracker.
