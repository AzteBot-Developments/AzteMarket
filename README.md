# AzteMarket
Discord bot application written in Go that manages the trading of benefits on the AzteMarket.

# Composing services
### Core
- `bot-service` (Handles Discord interactions like market-related slash commands, stock updates, price updates, etc.)

### Dependencies
- `aztebot-db` (Containerised MySQL instance for AzteBot DB in the local development)
- `aztemarket-db` (Containerised MySQL instance for AzteMarket DB in the local development)

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