
# Country API Assignment - Golang

This API service provides information about countries using the REST Countries API (<https://restcountries.com>).

## This project is Live on <https://assignment.snifyak.com/>

## Swagger Documentation can be accessed on <https://assignment.snifyak.com/swagger/index.html>

## Installation

To use this API, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/aayushsonu/country_assignment_api.git
    ```

2. Navigate to the project directory:

    ```bash
    cd country_assignment_api
    ```

3. Install dependencies:

    ```bash
    go mod download
    ```

4. Build the application:

    ```bash
    go build
    ```

5. Run the application:

    ```bash
    ./country_assignment_api
    ```

## Swagger Documentation

Explore the API interactively using Swagger UI:

[Swagger UI](http://localhost:8080/swagger/index.html) : <http://localhost:8080/swagger/index.html>

- After extracting auth token from /auth endpoint you have to put that token in other endpoint's Authorization Header. So, Don't forget to use the keyword `Bearer` to mention before the token like `Authorization: Bearer <AUTH_TOKEN>`

- You can also do this by using curl command. For this follow the below steps. (cURL utility must be installed in your system)

## BaseURL Endpoint: /api/v1

## Endpoints

### 1. Authentication

**Endpoint:** `/auth`

**Method:** `POST`

**Description:** Generates a valid auth token based on user credentials (username/password).

**Example:**

- Deployed Instance

```bash
curl -X POST -d "username=<USERNAME>&password=<PASSWORD>" http://assignment.snifyak.com/api/v1/auth
```

- Development

```bash
curl -X POST -d "username=<USERNAME>&password=<PASSWORD>" http://localhost:8080/api/v1/auth
```

- Replace `<USERNAME>` and `<PASSWORD>` with your credentials
- Valid Credential is USERNAME=snifyak and PASSWORD=123@snifyak@123

### 2. Fetch Country Information

**Endpoint:** `/country`

**Method:** `GET`

**Description:** Fetches detailed information about a specific country by providing its name as a parameter.

**Example:**

- Deployed Instance

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/country?name=India
```

- Development

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/country?name=India
```

### 3. Retrieve List of Countries based on filter

**Endpoint:** `/countries/filter`

**Method:** `GET`

**Description:** Retrieves a list of all countries' names based on filters (population/area/language) and sorting (asc/desc). Supports pagination.

**Example:**

- Default Page = 1
- Per Page = 20 countries
- Default sorting = asc
- Sorting option = asc/desc, if any other value provided then it will sort using ascending order, no error generated
- Sorting technique = Dictionary Based
- Optional paramters = population,area,lang,page,sort
- If no filter provided then it will response with 20 coutries name in ascending order

## Deployed Instance

**Filter using Population:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?population=2500000&sort=asc&page=1
```

**Filter using Area:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?area=948
```

**Filter using Language:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?lang=eng&sort=asc&page=1
```

**Filter using Population & Language:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?population=10000&lang=eng&sort=asc&page=1
```

**Pagination:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?population=50000000&page=2
```

**Sorting:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries/filter?lang=eng&sort=desc
```

## Development

**Filter using Population:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?population=2500000&sort=asc&page=1
```

**Filter using Area:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?area=948
```

**Filter using Language:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?lang=eng&sort=asc&page=1
```

**Filter using Population & Language:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?population=10000&lang=eng&sort=asc&page=1
```

**Pagination:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?population=50000000&page=2
```

**Sorting:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries/filter?lang=eng&sort=desc
```

### 4. Fetch All Countries (Optional) - for testing purpose only

**Endpoint:** `/countries`

**Method:** `GET`

**Description:** Fetches all countries details - For testing purpose only. Because it fetch all the details that's why it takes time to load (approx. 5-10 seconds)

**Example:**

- Deployed Instance

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://assignment.snifyak.com/api/v1/countries
```

- Development

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/api/v1/countries
```

## Error Handling

The API handles errors gracefully and returns appropriate error responses in case of failures.

## Notes

- For authentication, use the auth token obtained from the `/auth` endpoint in subsequent requests.
- Include the auth token in the `Authorization` header for protected endpoints.
