# URL Shortener

This project is a URL shortener built with Go and Echo framework, utilizing MySQL for data storage.

Feel free to check initial design of database schema [here](https://dbdocs.io/abdullahkabak322/URL-Shortener).

## Table of Contents
- [Installation](#installation)
- [Commit Tag Meanings](#commit-tag-meanings)
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
    ```
   
4. Install the dependencies:

    ```bash
    go mod download
    ```
   
5. Run the application:

    ```bash
    go run main.go
    ```
   
6. The application should now be running on `http://localhost:1323`.
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

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.