# Circulate Certstream

Monitor Certstream in real-time with keywords that are important to you and your business. Stay up to date on changes made to your certs, expiring certs, potential phishing domains, typo squatting, and more. Stay a step ahead of the adversary by proactively monitoring Certstream.

<br />

Simply add keywords that you care about to the KEYWORDS environment variable and start running the app in a matter of minutes.


## Roadmap:

- Add Terraform ECS Fargate IaC for simple production deployment
- Add Levenshtein distance to account for strings that are within a few characters of your keywords (e.g. domain squatting)
- Add other Alert destinations (create an issue & vote)


## Running locally:

1. If you have enabled alerting to ArangoDB, ensure that you correctly configured the Circulate database. If you would like to see other alerting destinations, open an issue! 

1. Build the Circulate Certstream docker container: <br />
    ```
    docker build . -t circulate-certstream
    ```

1. Run the docker container locally:
    ```
    docker run \
        -e KEYWORDS='["com","test","circulate","flrefox"]' \
        -p 8080:8080 circulate-certstream
    ```
