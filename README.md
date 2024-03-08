# URL Shortener

![Coverage](https://img.shields.io/badge/coverage-${COVERAGE_PERCENTAGE}%25-brightgreen)

This project is a URL shortener built with Go and Echo framework, utilizing MySQL for data storage.

Feel free to check the initial design of the database schema [here](https://dbdocs.io/abdullahkabak322/URL-Shortener).

## Table of Contents
- [Installation](#installation)
- [Commit Tag Meanings](#commit-tag-meanings)
- [Changelog](#changelog)
- [Security](#security)
- [Code Coverage](#code-coverage)
- [Contributing](#contributing)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/abdullahkabakk/url-shortener.git
    ```

2. Change directory to the project root:

    ```bash
    cd url-shortener
    ```

3. Create a `.env` file in the root directory and add the following environment variables:

    ```
    DB_USERNAME=<your_mysql_username>
    DB_PASSWORD=<your_mysql_password>
    DB_HOST=<mysql_host>
    DB_PORT=<mysql_port>
    DB_NAME=<database_name>
    HOST=<host_name>
    PORT=<port_name>
    MIGRATIONS_DIR=<migrations_dir>
    JWT_SECRET_KEY=<jwt_key>
    ```

4. Install the dependencies:

    ```bash
    go mod download
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

6. The application should now be running on `http://<HOST>:<PORT>`.
7. You can now access the application on your browser or using a tool like Postman.
8. You can also run the tests using the following command:

    ```bash
    go test ./...
    ```

## Commit Tag Meanings

Commit tags convey the nature of changes made in the codebase. Below are common commit tags and their meanings:

| Tag        | Description                                          | Example Commit Message                         |
|------------|------------------------------------------------------|------------------------------------------------|
| [feat]     | New feature or significant enhancement               | `[feat] Implement user authentication`         |
| [fix]      | Bug fix or correction to existing functionality      | `[fix] Resolve issue with user registration`   |
| [chore]    | Routine tasks, maintenance, or tooling changes       | `[chore] Update dependencies`                  |
| [docs]     | Changes or additions to documentation                | `[docs] Update installation guide`             |
| [style]    | Code style changes                                   | `[style] Format code according to style guide` |
| [refactor] | Code restructuring or optimization                   | `[refactor] Simplify user profile rendering`   |
| [test]     | Adding or modifying tests                            | `[test] Add unit tests for authentication`     |
| [ci]       | Changes to continuous integration (CI) configuration | `[ci] Update Travis CI configuration`          |
| [build]    | Changes affecting build system or dependencies       | `[build] Upgrade webpack to version 5`         |
| [perf]     | Performance improvements or optimizations            | `[perf] Optimize database queries`             |

## Changelog

All significant updates to this project will be meticulously documented in this [log](CHANGELOG.md).

## Security

The project is regularly scanned for security vulnerabilities using security tools like Gosec. The current number of security vulnerabilities detected is X.

## Code Coverage

The project's test coverage is tracked using tools like Golang Codecov. The current code coverage percentage is Y%. 

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.