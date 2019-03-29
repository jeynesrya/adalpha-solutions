CREATE TABLE IF NOT EXISTS assets (
    isin    VARCHAR PRIMARY KEY,
    asset   VARCHAR
);

CREATE TABLE IF NOT EXISTS portfolio (
    isin    VARCHAR REFERENCES assets(isin),
    amount  VARCHAR
)