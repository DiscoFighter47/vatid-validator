# :oncoming_police_car: VATID Validator
Microservice to validate german VATID.

## :rocket: Quick Start

#### Source
* Clone this repository
    ```bash
    $ git clone https://github.com/DiscoFighter47/vatid-validator.git
    ```
* Fetch dependencies
    ```bash
    $ go mod download
    ```
* Edit [config.example.yml](config.example.yml) file [optional]
* Run server
    ```bash
    $ CONFIG_PROVIDER=file CONFIG_FILE=config.example.yml go run main.go
    ```

#### Docker
* Create a config file similar to [config.example.yml](config.example.yml)
* Run docker image
    ```bash
    $ docker run -p 8080:8080 \
        -v {config_file_path}:/config/config.yml \
        -e CONFIG_PROVIDER=file \
        -e CONFIG_FILE=/config/config.yml \
        discofighter47/vatid-validator
    ```

## :page_facing_up: API Schema

#### VATID check

**URL** : `/api/v1/vatcheck/{VATID}`

**Method** : `GET`

## :tada: Features

- [x] **Config Management**
    - *File System*: Read config from file
    - *Extendable*: Strategy pattern was used to implement the config package. Extendable to handle remote config providers (e.g. Redis, Consul etc).
- [x] **Test**
    - *Unit Test*: Most of the cases are covered.
- [x] **Deployment**
    - *Docker*: Used multi stage build.
- [X] **CI/CD**
    - *Testing*: Triggered before merging to master branch.
    - *Docker Build & Push*: Triggered after merging to master branch.

## :hammer: Improvements

- [ ] **EUVIES Client**
    - *Caching*: Can implement a proxy over the client to handle caching using Redis.
- [ ] **Config Management**
    - *Remote Config*: Add remote providers.
- [ ] **Security**
    - *X-APP-Key*: Need to add app key checker.
    - *Rate Limiting*: Need to add request rate limiter using Redis.
- [ ] **Deployment**
    - *Helm Chart*: For k8s deployment.
- [ ] **Metric**
    - *Service Health*: Need to track service health. Can use solutions like Prometheus, New Relic, InfluxDB etc.
    - *Error Tracking*: Can use tools like Sentry.

## :rose: Acknowledgements
:pray: Special thanks goes to **Home24** and the team for inviting me to this coding challenge. It was really an interesting problem to solve.