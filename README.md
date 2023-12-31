# WS1CLI

![GitHub release](https://img.shields.io/github/release/geekxflood/WS1CLI.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/geekxflood/WS1CLI)](https://goreportcard.com/report/github.com/geekxflood/WS1CLI)
[![GoDoc](https://godoc.org/github.com/geekxflood/WS1CLI?status.svg)](https://godoc.org/github.com/geekxflood/WS1CLI)
[![License](https://img.shields.io/github/license/geekxflood/WS1CLI)](https://github.com/geekxflood/ws1cli/blob/main/LICENSE)

**DISCLAIMER**: This project is not affiliated with VMware or its parent/subsidiary companies. Any issues with this tool must be reported on the GitHub repository issue tracker and not VMware's support platforms or community forums.

`ws1cli` is a command-line interface (CLI) tool built with Go, leveraging the Cobra library to interact with VMware Workspace ONE UEM API. It streamlines the management and automation of tasks for administrators within the Workspace ONE platform.

This project aims for simplicity, with each command corresponding to a single API call. We strive to maintain this simplicity for ease of maintenance and extension. If you need to chain commands, parse the output from one command and input it into another.

The output is provided in JSON format, which can be parsed using tools like `jq`.

## Table of Contents

- [WS1CLI](#ws1cli)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [TODO](#todo)
  - [Installation](#installation)
  - [Configuration](#configuration)
    - [Environement Variables](#environement-variables)
  - [Commands](#commands)
    - [Version](#version)
    - [Test](#test)
    - [Devices](#devices)
    - [Products](#products)
    - [Remote](#remote)
  - [Tips](#tips)
  - [Contributing](#contributing)
  - [License](#license)
  - [Support](#support)

## Features

- Simplified Command Set: Easy-to-use commands for common administration tasks.
- Secure Configuration: Securely stores API credentials using advanced encryption.
- Go and Cobra: Built with Go for reliability and Cobra for a powerful command line interface.
- Commands implemented:
  - `version`: Display the version of `ws1cli`.
  - `test`: Verify configuration and connectivity to the API.
  - `device`: Interact with device API.
  - `product`: Interact with product provisionning API
  - `remote`: Generate remote control URL

## TODO

- Device Management: Perform devices operations.
  - [x] Get a list of devices.
  - [ ] Search a specific device.
  - [ ] Send command to a device.
  - [x] Send command to a list of devices.
  - [ ] Send a message to a device.
  - [ ] Lock a device.
  - [ ] Unlock a device.
  - [ ] Reboot a device (will depend if the OS allows it).
  - [ ] Wipe a device.
- Product List Interaction: Manage and list products within your Workspace ONE environment.
  - [x] Get a list of products.
  - [x] Start or stop a product.
  - [ ] Force Reprocess a product.
  - [ ] List devices assigned to a product and their status (will requiered several API call).
- Fetch Tag.
  - [ ] Get a list of tags.
  - [ ] Get a list of devices per tag.
  - [ ] Create a tag.
  - [ ] Delete a tag.
- Remote Session
  - [x] Generate an URL to start a remote session.
- Comprehensive Reporting: Generate reports on user access, device compliance, and application usage.

## Installation

Ensure Golang is installed on your system before installing `ws1cli`.

```bash
git clone https://github.com/geekxflood/WS1CLI.git
cd ws1cli
go build -ldflags="-X 'geekxflood/ws1cli/internal/version.Version=$(cat VERSION)'" -o ws1cli
```

Add `ws1cli` to your system's PATH to use it from any directory.

## Configuration

Initialize your configuration with the `init` command:

```bash
ws1cli init
```

You'll be prompted to enter your API URL, username, and password.
This will create a `~/.ws1cli` file in your home directory with your encrypted API credentials.
Your credentials and `aw-tenant-code` will be encrypted using AES-256-GCM with a passphrase you provide as an environment variable (see below).

### Environement Variables

`ws1cli` need the following environment variables:

| Variable | Description |
| --- | --- |
| `WS1_KEY` | Encryption passphrase |

This passphrase is used to encrypt your API credentials, the longer the better.
If you lost it, you will need to delete the `~/.ws1cli` file and reconfigure with `ws1cli init`.

## Commands

Here's how to get started with some basic commands:

- **Global flags**

| Flag | Description |
| --- | --- |
| `-h` `--help` | Help info |
| `-i` `--insecure` | Ignore TLS verification |
| `-p` `--pretty` | Pretty-print JSON output |

### Version

Display the version of `ws1cli`:

```bash
ws1cli version
```

### Test

To verify API configuration and connectivity:

```bash
ws1cli test
```

| Flag | type |Description |
| --- | --- | --- |
| `-d` `--show-details` | `N/A` | **INSECURE** Display the URL and headers for the test API call |

The `-d`or `--show-details` flag is mentionned as **INSECURE** because it will display your API credentials in plain text.

### Devices

To interact with device API:

```bash
ws1cli device
```

| Flag | type |Description |
| --- | --- | --- |
| `-d` `--inventory` | `N/A` | Output an array of JSON of devices in LGID |
| `-l` `--lgid` | `int` | LGID of device, mandatory when using `-d` flag |
| `-c` `--command` | `string` | Command to send, command accepted: `EnterpriseWipe`, `LockDevice`, `ScheduleOsUpdate`, `SoftReset`, `Shutdown`, |
| `-f` `--filter` | `string` | Filter to apply to the device list (for `command` flag) Value accepted: `Macaddress`, `Udid`, `Serialnumber`, `ImeiNumber` |
| `-inputJson` | `string` | JSON input for the command to send, it can be the output of the `--inventory` |

### Products

To interact with product provisionning API:

```bash
ws1cli product
```

| Flag | type |Description |
| --- | --- | --- |
| `-l` `--lgid` | `int` |LGID of device, mandatory when using `-d` flag |
| `-s` `--start-stop` | `bool` | Start (true/1) or stop (false/0) a product |
| `-o` `--product-id` | `int` | Product ID for starting or stopping a product |

### Remote

To interact with remote management API:

```bash
ws1cli remote
```

| Flag | type |Description |
| --- | --- | --- |
| `-d` `--device-uuid` | `string` | Device UUID |
| `-s` `--session-type` | `string` | type of session to be created `ScreenShare`, `FileManager`, `RemoteShell`, `RegistryEditor` |

The command return a JSON with an URL to open in a browser to start the session.

## Tips

JSON output can be parse using `jq`:

Find the number of devices in LGID:

```bash
./ws1cli device --inventory --lgid $LGID | jq 'length'
```

Execute a batch command to a list of devices in a specific LGID:

```bash
ws1cli --command "SoftReset" --filter "Serialnumber" --inputJson $(ws1cli --inventory --lgid $LGID)
```

In this example the LGID is 500, the command is `SoftReset` and the filter is `Serialnumber`.

Same command but on devices from a large LGID:

```bash
ws1cli --inventory --lgid $LGID > devices.json
ws1cli --command "SoftReset" --filter "Serialnumber" --inputJson "$(cat devices.json)"
```

In this example, we use a file to store the output of the first command and use it as input for the second command.
So that we are sure that we have all the devices in the LGID since the first command can take some time to finish

Not all command are supported by all platform, you can check the API documentation for more information.

## Contributing

Contributions are welcome. Fork the project, make your updates, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please open an issue in the GitHub repository issue tracker.
