CREATE TABLE WEBURL (
    ID SERIAL PRIMARY KEY,
    URL text NOT NULL,
    CREATED_AT timestamp with time zone default current_timestamp--,
    --COUNTER
)