# PostgreSQL Test Application

Application for testing PostgreSQL Patroni Cluster

### Usage

1. Create test table
    ```sql
    CREATE TABLE test_table (
        t text
    );
    ```
1. Edit haproxy.cfg
1. Make `.env` file
    ```
    DB_HOST=localhost
    DB_PORT_RW=5432
    DB_PORT_RO=5433
    DB_USER=dbuser
    DB_PASS=Qwerty123
    DB_NAME=db
    ```
1. Run `make run`
