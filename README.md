# ws1-cli

`ws1-cli` is a command-line interface (CLI) tool built in Go and leveraging the Cobra library, designed for interacting with VMware Workspace ONE UEM API. It enables administrators to manage and automate tasks within the Workspace ONE platform efficiently.

## Features

- Simplified Command Set: Easy-to-use commands for common administration tasks.
- Secure Configuration: Securely stores API credentials using advanced encryption.
- Go and Cobra: Built with Go for reliability and Cobra for a powerful command line interface.

## Planned Features

- Product List Interaction: Manage and list products within your Workspace ONE environment.
- Fetch Tag List: Retrieve a list of tags associated with devices or users.
- Device Management: Perform operations such as device lock, wipe, and query information.
- Automated Actions: Schedule and automate routine tasks for device and application management.
- Comprehensive Reporting: Generate reports on user access, device compliance, and application usage.

## TODO

- Integrate More API Endpoints: Expand command functionality to cover all aspects of the API provided by Workspace ONE.
- Interactive CLI Wizard: Create an interactive CLI wizard for complex workflows.
- Enhanced Security Features: Additional security checks and configuration validation.
- Custom Report Building: Tools for building custom reports via the CLI.
- Scripting and Automation: Guides and templates for common scripting and automation use cases.
- Bulk Operations: Efficient methods for bulk updates or deletions.
- Compliance and Reporting: Tools for conducting and reporting compliance checks.
- Event Log Monitoring: Capabilities to monitor and extract event logs.

## Installation

Ensure Golang is installed on your system before installing `ws1-cli`.

```bash
git clone https://github.com/christopherime/ws1-cli.git
cd ws1-cli
go build
```

Add `ws1-cli` to your system's PATH to use it from any directory.

## Configuration

Initialize your configuration with the `init` command:

```bash
ws1-cli init
```

You'll be prompted to enter your API URL, username, and password.

## Usage

Here's how to get started with some basic commands:

### Get Devices

To fetch a list of devices:

```bash
ws1-cli getDevices
```

### Test Configuration

To verify API configuration and connectivity:

```bash
ws1-cli test
```

## Contributing

Contributions are welcome. Fork the project, make your updates, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Support

For support, please open an issue in the GitHub repository issue tracker.

## Author

- Christophe Rime - [christopherime@me.com](mailto:christopherime@me.com)

