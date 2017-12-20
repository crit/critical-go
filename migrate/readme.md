# Migrate

DSN format: `un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci`

## Build

`make build`

## Usage

- Linux: `./bin/linux/migrate "dsn" "scriptsDirectory"`
- Mac: `./bin/darwin/migrate "dsn" "scriptsDirectory"`
- Windows: TBD

Example: `./bin/linux/migrate "un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci" "/code/folder_of_sql_files"`

## Supported DB Drivers

MySQL (or MariaDB) only at the moment.

## Testing

Create a file named `.env` (in the same folder as `migrate_test.go`) with the following content (replacing all values needed to connect):

```
DB_DSN="un:pw@tcp(server:port)/db_name?multiStatements=true&collation=utf8mb4_general_ci"
```

Then run `make tests` which will attempt to attach to the above DSN, create a migration folder in the OS temp directory, add some SQL 
files to the new migration folder, and then run the Migrate function against these files.
