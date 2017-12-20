# Migrations

DSN format: `un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci`

## Supported DB Drivers

MySQL (or MariaDB) only at the moment.

## Testing

Create `.env` file for testing with the following content (replacing all values needed to connect):

```
DB_DSN="un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci"
```
