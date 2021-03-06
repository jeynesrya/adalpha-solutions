# AdAlpha Solutions Technical Task
## Problem
An investor has an account that is setup in a modelled portfolio with a current
value of `£220,000`. The model has the following purchasable assets in it.

|ISIN   | ASSET  | % Of portfolio  |
|---|---|---|
| IE00B52L4369  | BlackRock Institutional Cash Series Sterling Liquidity Agency Inc  | 20  |
| GB00BQ1YHQ70  | Threadneedle UK Property Authorised Investment Net GBP 1 Acc  | 20  | 
| GB00B3X7QG63  | Vanguard FTSE U.K. All Share Index Unit Trust Accumulation  | 10  |
| GB00BG0QP828  | Legal & General Japan Index Trust C Class Accumulation  | 5  |
| GB00BPN5P238  | Vanguard US Equity Index Institutional Plus GBP Accumulation  | 15  |
| IE00B1S74Q32  | Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation  | 30  |

The investor would like to withdraw an amount `£X`. Your challenge is to create
an application that programmatically creates an order sheet for this problem.

An order sheet can have the following instructions:

**BUY** - a buy instruction is placed in units specifing the number of units of
a particular fund that is wanted.

**INVEST** - an invest instruction is placed in a currency specifying how much
you would like to buy.

**SELL** - a sell instruction is placed in units specifying how many units to
sell.

**RAISE** - a raise instruction is placed in a currency specifying how much you
are trying to receive from your holding.

## Solution
- VueJS: User Interface for the investor to see their portfolio and alter as needed
- Go: Backend code used to take user input and make changes to database
- PostgreSQL: Database solution for storing the investors portfolio
- Elasticsearch/Kibana: Used for storing and visualising application error logs

### Prerequisites
- npm
- docker
- docker-compose
- go

### Running the application
`./run-app.sh`

Once running, the application should be available at `localhost:8080` and `127.0.0.1:8080`.
Kibana is available to view application logs at `localhost:5601` and `127.0.0.1:5601`
    The logs will appear in the index "logs" but you will only be able to find them when the first error has occurred.

### Testing the application
`./run-tests.sh` 
