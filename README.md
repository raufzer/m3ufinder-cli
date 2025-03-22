# M3U Finder CLI

M3U Finder CLI is a command-line tool designed for extracting and validating M3U links from various IPTV providers. The tool allows users to scrape IPTV links, save them in different formats, and interact with the application in an interactive mode.

![M3U Finder CLI](assets/m3ufinder.png)

## Features

- Scrape IPTV links from various providers.
- Save scraped links in `.m3u`, `.json`, or `.csv` formats.
- Interactive mode for easy command execution.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/raufzer/m3ufinder-cli.git
   ```
2. Navigate to the project directory:
   ```sh
   cd m3ufinder-cli
   ```
3. Build the CLI application:
   ```sh
   go build -o m3ufinder-cli
   ```

## Usage

1. Run the CLI application:
   ```sh
   ./m3ufinder run
   ```

2. Use the interactive mode to execute commands:
   - **collect [category]**: Extracts IPTV links from websites. Category is optional.
   - **save [format]**: Exports extracted links in `.m3u`, `.json`, or `.csv` formats.
   - **exit**: Exit interactive mode.

### Example

```sh
./m3ufinder run
```

Interactive mode commands:
```
1. collect [category] - Extracts IPTV links from websites. Category is optional.
2. save [format] - Exports extracted links in .m3u, .json, or .csv formats.
3. exit - Exit interactive mode.
```

## Development Status

This project is still in development and may not have all features fully implemented. Contributions and feedback are welcome.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.