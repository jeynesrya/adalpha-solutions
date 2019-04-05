CREATE TABLE IF NOT EXISTS assets (
    isin    VARCHAR PRIMARY KEY,
    asset   VARCHAR
);

CREATE TABLE IF NOT EXISTS portfolio (
    isin    VARCHAR REFERENCES assets(isin),
    units  NUMERIC
);

INSERT INTO assets VALUES
    ('IE00B52L4369', 'BlackRock Institutional Cash Series Sterling Liquidity Agency Inc'),
    ('GB00BQ1YHQ70', 'Threadneedle UK Property Authorised Investment Net GBP 1 Acc'),             
    ('GB00B3X7QG63', 'Vanguard FTSE U.K. All Share Index Unit Trust Accumulation'),
    ('GB00BG0QP828', 'Legal & General Japan Index Trust C Class Accumulation'),
    ('GB00BPN5P238', 'Vanguard US Equity Index Institutional Plus GBP Accumulation'),
    ('IE00B1S74Q32', 'Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation');
    
INSERT INTO portfolio VALUES
    ('IE00B52L4369', 44000),
    ('GB00BQ1YHQ70', 37931.0345),             
    ('GB00B3X7QG63', 109.420074),
    ('GB00BG0QP828', 17549.4576),
    ('GB00BPN5P238', 175.298805),
    ('IE00B1S74Q32', 676.229508);