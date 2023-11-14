
# Country API

This API service provides information about countries using the REST Countries API (https://restcountries.com).

<!-- http://localhost:8080/countries/filter?population=100000000&area=500000&sort=desc&page=2&itemsPerPage=5 -->

## Endpoints

### 1. Authentication

**Endpoint:** `/auth`

**Method:** `POST`

**Description:** Generates a valid auth token based on user credentials (username/password).

**Example:**

```bash
curl -X POST -d "username=snifyak&password=123@snifyak@123" http://localhost:8080/auth
```

### 2. Fetch Country Information

**Endpoint:** `/country`

**Method:** `GET`

**Description:** Fetches detailed information about a specific country by providing its name as a parameter.

**Example:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/country?name=India
```

### 3. Retrieve List of Countries

**Endpoint:** `/countries`

**Method:** `GET`

**Description:** Retrieves a list of all countries' names based on filters (population/area/language) and sorting (asc/desc). Supports pagination.

**Example:**

```bash
curl -H "Authorization: Bearer <your_auth_token>" http://localhost:8080/countries/filter?population=250000&sort=asc&page=1
```

## Error Handling

The API handles errors gracefully and returns appropriate error responses in case of failures.

## Notes

- For authentication, use the auth token obtained from the `/auth` endpoint in subsequent requests.
- Include the auth token in the `Authorization` header for protected endpoints.
