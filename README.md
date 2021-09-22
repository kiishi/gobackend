# Go Backend

To run the project, use the below command after creating a .env file that contains the .env.sample format

## Endpoints

### InMemory

- Create In Memory Record
  ```bash
  curl --location --request POST 'https://getir-go-backend.herokuapp.com/in-memory/create' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "key":"dog",
    "value":"hsod"
  }'
  ```
- Get In Memory Record
  ```bash
  curl --location --request GET 'https://getir-go-backend.herokuapp.com/in-memory?key=dog' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "key":"dog",
    "value":"hsod"
  }'
  ```

### Records Collection

- Get Records
  ```bash
  curl --location --request POST 'https://getir-go-backend.herokuapp.com/get-records' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "startDate":"2016-05-31",
    "endDate":"2021-05-31",
    "minCount": 5500,
    "maxCount": 6000
  }'
  ```
