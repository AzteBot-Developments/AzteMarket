# AzteMarket
Discord bot application written in Go that manages the trading of benefits on the AzteMarket.

# Composing services
### Core
- `bot-service` (Handles Discord interactions like market-related slash commands, stock updates, price updates, etc.)

### Dependencies
- `mysql-db` (Containerised MySQL instance for local development)