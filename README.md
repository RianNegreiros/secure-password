# Secure Password Validation

This is a challenge from [BackEndBrasil](https://github.com/backend-br/desafios/blob/master/secure-password/PROBLEM.md)

The project provides a simple password validation API with basic security criteria.

## Getting Started

To run the project locally, follow these steps:

1. Clone the repository:

```bash
  git clone https://github.com/RianNegreiros/secure-password.git
```

2. Navigate to the project directory:

```bash
  cd secure-password
```

3. Run the application:

```bash
  go run cmd/main.go
```

The API will be accessible at `http://localhost:8080/validate-password`.

## API Endpoints

### Validate Password

- Endpoint: POST `/validate-password`
- Request Body: JSON with a single field password
- Response:
  - 204 No Content if the password meets the security criteria.
  - 400 Bad Request with error messages if the password is invalid.

## Running Tests

To run the tests, use the following command:

```bash
go test ./...
```

To manually test, you can use [curl](https://curl.se/download.html):

```bash
  curl -X POST -H "Content-Type: application/json" -d '{"password": "yourPASSWORD123!"}' http://localhost:8080/validate-password
```

## Contributing

If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
