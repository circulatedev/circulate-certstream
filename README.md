# Circulate Certstream

Monitor Certstream in real-time with keywords that are important to you and your business. Stay up to date on changes made to your certs, expiring certs, potential phishing domains, typo squatting, and more. Stay a step ahead of the adversary by proactively monitoring Certstream.

<br />

## How To Use This App:

Simply add keywords that you care about (e.g. domains you own, other org domains that you commonly interface with, etc.) to the KEYWORDS environment variable and run the docker container. <br />
`Hint - for a quick test, if you leave "com" in the list of keywords then you will see many matches!`

<br />

## Running locally:

1. If you have enabled alerting to ArangoDB, ensure that you correctly configured the Circulate database. If you would like to see other alerting destinations, open an issue! 

1. Build the Circulate Certstream docker container: <br />
    ```
    docker build . -t circulate-certstream
    ```

1. Run the docker container locally using one of the following methods:
    - With Circulate's ArangoDB running and configured:
        ```
        docker run \
            -e KEYWORDS='["com","test","circulate","flrefox"]' \
            -e ENABLE_ARANGODB='true' \
            -e DB_ENDPOINT="tcp://172.17.0.2:8529" \
            -e ROOT_PASSWORD="password" \
            -e CA_CERT="" \
            -p 8080:8080 circulate-certstream
        ```
    - Print to console (quickest way to test):
        ```
        docker run \
            -e KEYWORDS='["com","test","circulate","flrefox"]' \
            -p 8080:8080 circulate-certstream
        ```

<br />

## Current Targets:

- Circulate's Database
    ```
    -e ENABLE_ARANGODB='true' \
    -e DB_ENDPOINT="tcp://172.17.0.2:8529" \
    -e ROOT_PASSWORD="password" \
    -e CA_CERT="" \
    ```
- Console
    ```
    -e ENABLE_CONSOLE_OUTPUT='true' \
    ```

Interested in seeing more targets? Create an issue

<br />

## Roadmap:

- Graph Relationships between Keywords and Matches in Circulate DB
- Terraform ECS Fargate IaC for simple production deployment
- Levenshtein distance to account for strings that are within a few characters of your keywords (e.g. domain squatting)
- ArangoDB batch support for higher volume data collection
- Other Alert destinations (create an issue / vote with a thumbs up)
- Could be interesting to take any matches and enumerate further data (e.g. network enumeration of DNS names that have been matched)
