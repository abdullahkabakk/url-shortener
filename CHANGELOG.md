# Changelog
All significant updates to this project will be meticulously documented in this log.

## 0.7.0 - 13/03/2024

### Added

- **Get User Clicks Endpoint:** Added a new route to the echo server to handle the user's clicks.

- **Get User Clicks Handler:** Added a new handler to handle the user's clicks.

- **Get User Clicks Service:** Added a new service to handle the user's clicks.

- **Get User Clicks Repository:** Added a new repository to handle the user's clicks.

- **Get User Short URL Service:** Added a new service to handle the user's short URL.

- **Get User Short URL Repository:** Added a new repository to handle the user's short URL.

- **Token Service inside Click Handler:** Added the token service to the click handler.
  - ***Reason:*** This change was made to ensure that the token service is used to validate the token.
  - ***Impact:*** This change will affect the click handler.

- **Clicks Model:** Added a clicks model to the project.

## 0.6.2 - 13/03/2024

### Added

- **Validation:** Added validation to the auth and url handler.

### Changed

- **Database Migration:** Updated the database migration removed the `id` from url table.

## 0.6.1 - 11/03/2024

### Changed

- **Error handling:** Added error handling for create_url in tests

## 0.6.0 - 11/03/2024

### Added

- **User's URL Endpoint:** Added a new route to the echo server to handle the user's URL.

## 0.5.1.alpha.2 - 11/03/2024

### Added

- **Get User URL Functionality:** Added functionality to retrieve the user's URL.

- **Mock URL Repository:** Added a mock URL repository to the project.

## 0.5.1.alpha.1 - 11/03/2024

### Added

- **Get User URL Test:** Added a test to ensure that the user's URL is retrieved.

### Changed

- **Parse Time:** Updated the parse time for database.

- **Renamed RegistrationDate to CreatedAt:** Renamed the `registration_date` field to `created_at` in the user model.

- **Renamed CreationDate to CreatedAt:** Renamed the `creation_date` field to `created_at` in the click model.

## 0.5.1 - 10/03/2024

### Added

- **Empty Token Test:** Added a test to ensure that empty token are not sent to the server.

### Changed

- **Refactored Auth Repository Test:** Refactored the auth repository test to use the testify library for better assertions.

## 0.5.0 - 10/03/2024

### Added

- **Refresh Token Route:** Added a new route to the echo server to handle refresh token requests.

## 0.4.1.alpha.2 - 10/03/2024

### Added

- **Auth Handler:** Added a new handler to handle refresh token requests. 

### Changed

- **Auth Query:** Updated the auth query from `registration_date` to `created_at`.
  - ***Reason:*** This change was made to ensure that the correct field is used to store the registration date.
  - ***Impact:*** This change will affect the database schema and the auth repository.
- **Console Log:** Removed console logs from url_handler

## 0.4.1.alpha.1 - 10/03/2024

### Added

- **Refresh Token Test:** Added a test to ensure that the refresh token is generated.

### Docs

- **Postman Collection:** Added a postman collection to the project.

## 0.4.1 - 10/03/2024

### Changed

- **Refactored tests:** Refactored the tests to use the `testify` library for better assertions.

## 0.4.0 - 10/03/2024

### Added

- **Echo Click Route:** Added a new route to the echo server to handle click tracking.

## 0.3.2.alpha.2 - 10/03/2024

### Added

- **Click Handler:** Added a click handler to the project.

- **Mock Click Repository:** Added a mock click repository to the project.

- **Click Repository:** Added a click repository to the project.

## 0.3.2.alpha.1 - 10/03/2024

### Added

- **Click Handler:** Added a click handler to the project.

## 0.3.2 - 10/03/2024

### Changed

- **Auth Service Return Type:** Updated the return type of the auth service to return a pointer to the user model.
  - ***Reason:*** This change was made to ensure that the token generation function can access the user ID.

- **URL Parsing:** Added the URL parsing to use the `url.Parse` function.
  - ***Reason:*** This change was made to ensure that the URL is parsed correctly.

- **Gitignore:** Updated the `.gitignore` file to include the `migrations` file.

### Docs

- **Swagger Documentation:** Updated the swagger documentation to change url shortening endpoint

- **README:** Updated the README file to include instructions on how to use the URL shortening endpoint.

## 0.3.1 - 09/03/2024

### Changed

- **CI/CD Pipeline:** Updated the CI/CD pipeline actions/go version @v2 to actions/setup-go@v5.
  - **set-output:** Removed the set-output from the CI/CD pipeline used GITHUB_OUTPUT instead.

- **Server Test:** Updated the server test with better assertions.

- **`/`**: Updated the root route to return a welcome message.

## 0.3.0 - 09/03/2024

### Added

- **URL Repository:** Added a URL repository to the project.

- **URL Shorten API:** Added a new route to the echo server to handle URL shortening.

## 0.2.0.alpha.2 - 09/03/2024

### Added

- **URL Service:** Added a URL service to the project.

- **Generate Short URL:** Added a function to generate a short URL.

- **Mock URL Repository:** Added a mock URL repository to the project.

## 0.2.0.alpha.1 - 09/03/2024

### Added

- **URL Handler:** Added a URL handler to the project.

- **URL Model:** Added a URL model to the project.

## 0.2.0 - 09/03/2024

### Added

- **API Documentation:** Added API documentation to the project.

### Changed

- **Directory of User Model:** Moved the user model to the `internal/app/models/user` directory.

- **Directory of Auth Handler:** Moved the auth handler to the `internal/app/handler/auth` directory.

- **Directory of Auth Repository:** Moved the auth repository to the `internal/app/repository/auth` directory.

- **Directory of Auth Service:** Moved the auth service to the `internal/app/services/auth` directory.

- **Make Token as a Service:** Refactored the code to make token generation a service.

- **Handler Generates Token**: Refactored the code to make the handler generate tokens instead of auth service.

- **Auth Handler Takes Token Service**: Refactored the code to make the auth handler take the token service as a dependency.

## 0.1.0 - 08/03/2024

### Changed

- **Versioning:** Updated the versioning system to a stable release.

## 0.0.0.alpha.4 - 08/03/2024

### Added

- **User Authentication:** Added user authentication to the project.

- **JWT Library:** Added the JWT library to the project.

- **Auth Repository:** Added an auth repository to handle user authentication.

- **Auth Service:** Added an auth service to handle user authentication.

- **Auth Handler:** Added an auth handler to handle user authentication.

## 0.0.0.alpha.3 - 08/03/2024

### Added

- **User Model:** Added a user model to the project.

- **Migration:** Added a migration to create the users table.

- **Echo Server:** Added a new route to the echo server to handle user registration.

- **Echo Library:** Added the echo library to the project.

- **Added New Instructions For Makefile:** Added new instructions to the Makefile to handle the new changes.

### Changed

- **Database:** Updated the database schema to include the users table.

- **README:** Updated the instructions in the README file.

## 0.0.0.alpha.2 - 07/03/2024

### Added

- **Pipeline:** Added a pipeline to the project.

## 0.0.0.alpha.1 - 07/03/2024

### Added

- **Changelog File:** Added a new changelog file to keep track of the changes made to the project.

- **Versioning:** Added a versioning system to the project.