# WS1CLI

`ws1cli` is a command-line interface (CLI) tool built in Go and leveraging the Cobra library, designed for interacting with VMware Workspace ONE UEM API. It enables administrators to manage and automate tasks within the Workspace ONE platform efficiently.

## Features

- Simplified Command Set: Easy-to-use commands for common administration tasks.
- Secure Configuration: Securely stores API credentials using advanced encryption.
- Go and Cobra: Built with Go for reliability and Cobra for a powerful command line interface.
- Commands implemented:
  - `version`: Display the version of `ws1cli`.
  - `test`: Verify configuration and connectivity to the API.
  - `device`: Interact with device API.
    - `-d` or `--inventory`: Output an array of JSON of devices in LGID (`-l`, `--lgid` *mandatory* with it).
  - `product`: Interact with product provisionning API
    - `-l` or `--lgid`: **mandatory** LGID of products

## TODO

- [ ] Product List Interaction: Manage and list products within your Workspace ONE environment.
- [ ] Fetch Tag List: Retrieve a list of tags associated with devices or users.
- [ ] Device Management: Perform operations such as device lock, wipe, and query information.
- [ ] Comprehensive Reporting: Generate reports on user access, device compliance, and application usage.

## Environement Variables

`ws1cli` need the following environment variables:

| Variable | Description |
| --- | --- |
| `WS1_KEY` | Encryption passphrase |

## Installation

Ensure Golang is installed on your system before installing `ws1cli`.

```bash
git clone https://github.com/christopherime/WS1CLI.git
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

| Flag | Description |
| --- | --- |
| `-d` `--show-details` | **INSECURE** Display the URL and headers for the test API call |

The `-d`or `--show-details` flag is mentionned as **INSECURE** because it will display your API credentials in plain text.

### Devices

To interact with device API:

```bash
ws1cli device
```

| Flag | Description |
| --- | --- |
| `-d` `--inventory` | Output an array of JSON of devices in LGID |
| `-l` `--lgid` | LGID of device, mandatory when using `-d` flag |

### Products

To interact with product provisionning API:

```bash
ws1cli product
```

| Flag | Description |
| --- | --- |
| `-l` `--lgid` | LGID of device, mandatory when using `-d` flag |

## Tips

JSON output can be parse using `jq`:

Find the number of devices in LGID:

```bash
./ws1cli device --inventory -l $LGID | jq 'length'
```

## Contributing

Contributions are welcome. Fork the project, make your updates, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support

For support, please open an issue in the GitHub repository issue tracker.
