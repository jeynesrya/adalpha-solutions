export interface InternalOrder {
    Isin: string;
    Type: string;
    Amount: number;
    Currency: string;
}

export interface UnitOrder {
    Isin: string;
    Amount: number;
}

export interface CurrencyOrder extends UnitOrder {
    Currency: string;
}

export type Order = UnitOrder | CurrencyOrder;
