# URL Shortener

[![Go Report Card](https://goreportcard.com/badge/github.com/abdullahkabakk/url-shortener)](https://goreportcard.com/report/github.com/abdullahkabakk/url-shortener)
[![Build](https://github.com/abdullahkabakk/url-shortener/actions/workflows/go.yml/badge.svg)](https://github.com/abdullahkabakk/url-shortener/actions/workflows/go.yml)
[![Coverage](https://img.shields.io/codecov/c/github/abdullahkabakk/url-shortener)](https://codecov.io/github/abdullahkabakk/url-shortener)
[![Code Style](https://img.shields.io/badge/code%20style-gofmt-%2300ADD8)](https://golang.org/doc/effective_go.html#formatting)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/abdullahkabakk/url-shortener/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/abdullahkabakk/url-shortener)](https://golang.org/doc/devel/release.html)
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/29774798-07fc0dc3-e608-40cc-bb6a-d9adf3248971?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D29774798-07fc0dc3-e608-40cc-bb6a-d9adf3248971%26entityType%3Dcollection%26workspaceId%3D108d3a9c-b63c-4312-b0f2-f9ec95727ae4)

This project is a URL shortener built with Go and Echo framework, utilizing MySQL for data storage.

Feel free to check the initial design of the database schema [here](https://dbdocs.io/abdullahkabak322/URL-Shortener).

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technologies](#technologies)
- [API Endpoints](#api-endpoints)
- [Installation](#installation)
- [Usage](#usage)
- [Directory Structure](#directory-structure)
- [Commit Tag Meanings](#commit-tag-meanings)
- [Changelog](#changelog)
- [Security](#security)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The project is a URL shortener that allows users to shorten long URLs into short, easy-to-remember URLs. 
It also provides a redirection service that allows users to redirect to the original URL by visiting the short URL.
The project also includes user authentication to allow users to create accounts and manage their shortened URLs.

## Features

- User authentication
- URL shortening
- URL redirection

## Technologies

- Go
- Echo
- MySQL

## API Endpoints

The following are the API endpoints available in the application:

### Auth

- `POST /auth/register`: Register a new user
- `POST /auth/login`: Login a user

### URL

- `POST /url/shorten`: Shorten a URL

### Clicks

- `GET /clicks/:shortURL`: Redirect to the original URL

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

9. You can also make coverage reports using the following command:

    ```bash
    make report
    ```

## Database Migrations

The project uses Golang Migrate for database migrations. But doesn't include migration files. You need to create migration files for user, url, and click tables.

## Usage

To create a user, run the following command:

```bash
curl -X POST http://localhost:8080/auth/register -d '{"username": "user", "password": "password"}'
```

To log in a user, run the following command:

```bash
curl -X POST http://localhost:8080/auth/login -d '{"username": "user", "password": "password"}'
```

To shorten a URL, run the following command:

```bash
curl -X POST http://localhost:8080/url/shorten -d '{"url": "https://www.google.com"}' -H "Authorization
```

To redirect to the original URL, run the following command:

```bash
curl -X GET http://localhost:8080/clicks/<shortURL>
```

## Directory Structure

The project's directory structure is as follows:

```
.
├── internal
│   ├── app
│   │   ├── handler
│   │   │   ├── auth
│   │   │   └── url
│   │   ├── models
│   │   │   ├── url
│   │   │   └── user
│   │   ├── repository
│   │   │   ├── auth
│   │   │   └── url
│   │   └── services
│   │       ├── auth
│   │       └── url
│   ├── config
│   ├── infrastructure
│   │   ├── database
│   │   │   └── migrations
│   │   └── http
│   ├── mocks
│   └── utils
...
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

If you discover any security vulnerabilities within the URL shortener project, please follow these steps to report them:
[Security Policy](SECURITY.md)

## Contributing

Find out how to contribute to the URL Shortener project by following the guidelines provided in the [CONTRIBUTING](CONTRIBUTING.md) file.

## License

This project is open-source and is available under the [MIT License](LICENSE).