CREATE TABLE transactions
(
    id                  SERIAL PRIMARY KEY,
    user_id             INT         NOT NULL,
    user_email          VARCHAR(40) NOT NULL,
    amount              NUMERIC     NOT NULL,
    currency            VARCHAR(40) NOT NULL,
    date_of_creation    TIMESTAMP,
    date_of_last_change TIMESTAMP,
    status              VARCHAR(40) NOT NULL
);