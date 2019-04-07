import axios from 'axios';

import { PortfolioItem } from './portfolio';
import { Order } from './order';

const api = 'http://localhost:8081';

class PortfolioService {
  public getPortfolio() {
    return axios.get<PortfolioItem[]>(`${api}/portfolio`);
  }
  public updatePortfolio(order: Order, type: string) {
    return axios.post(`${api}/instruction/${type}`, order);
  }
}

// Export a singleton instance in the global namespace
export const portfolioService = new PortfolioService();
