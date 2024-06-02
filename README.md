# AzteMarket
Discord bot application written in Go that manages the trading of benefits on the AzteMarket.

# Composing services
### Core
- `bot-service` (Handles Discord interactions like market-related slash commands, stock updates, price updates, etc.)

### Dependencies
- `aztebot-db` (Containerised MySQL instance for AzteBot DB in the local development)
- `aztemarket-db` (Containerised MySQL instance for AzteMarket DB in the local development)

# How to Run
## Prerequisites
In order to run the application, a few prerequisites must be met.
1. Have the repository cloned locally.
2. Have Docker installed.
3. Have Make installed.
4. Have a fully-configured `.env` file saved in the root of the repository. (contact [@RazvanBerbece](https://github.com/RazvanBerbece) for the configuration)
5. Additionally, for full local development capabilities and to run the database migrations on the development machine, have the [Aztebot-Infrastructure](https://github.com/RazvanBerbece/Aztebot-Infrastructure) repository cloned locally in a folder which also contains the `Aztebot` repository (**For example**, the folder `Project` should contain both the `Aztebot` and the `Aztebot-Infrastructure` repository folders) 

### Notes
At the moment, to propagate remote DB changes to the local dev environment, the `Infrastructure` submodule has to be updated manually when there are changes in the remote source (e.g. a new migration file). 
This can be done by running `git submodule update --recursive --remote --init` in the root of this repository.

## Running the full service composition
1. Run a freshly built full service composition (app, DBs, etc.) with the `make up` command.
    - This is required so the local development database is configured with all the necessary default data.   
2. Once the `mysql-db` service has sucessfully started, run the DB migrations locally by executing the following commands from the root of this repository (_requires [Aztebot-Infrastructure](https://github.com/RazvanBerbece/Aztebot-Infrastructure) as described in prerequisite #5_)
    - To execute a dryrun and double-check the to-be-applied migrations: `make migrate-up-dry` 
    - To apply the migrations `make migrate-up`

To bring down all the services, one can do so by running `make down`.

# CI/CD
This project will employ CI/CD through the use of GitHub Actions and Google Cloud. 

## CI
Continuous integration is implemented through a workflow script which sets up a containerised service composition containing the Go environment and other dependencies (MySql, etc.) and then runs the internal logic tests on all pull request and pushes to main. The workflow file for the AzteBot CI can be seen in [test.yml](.github/workflows/test.yml).

## CD
Continuous deployment is implemented through a workflow script which builds all the project artifacts and uploads them to Google Cloud Artifact Registry on pushes to the main branch. Additionally, a GKE pod is created with the new container image and ultimately executed upstream to run the apps. The workflow file for the AzteBot CD can be seen in [deploy.yml](.github/workflows/deploy.yml).

Notes:
- The production environment file is base64 encoded using `make update-envs` and decoded accordingly in the Actions workflows.

# Contribution Guidelines
TODO

### Merge Commit Messages
**Must** contain one of the commit message types below such that the bump and release strategy works as intended.
- feat(...): A new feature
- fix(...): A bug fix
- docs(...): Documentation only changes
- style(...): Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- refactor(...): A code change that neither fixes a bug nor adds a feature
- perf(...): A code change that improves performance
- test(...): Adding missing or correcting existing tests
- chore(...): Changes to the build process or auxiliary tools and libraries such as documentation generation