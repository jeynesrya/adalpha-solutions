import axios from 'axios';

import { PortfolioItem } from './portfolio';

const api = 'http://localhost:8081';

class PortfolioService {
  public getPortfolio() {
    return axios.get<PortfolioItem[]>(`${api}/portfolio`);
  }
}

// Export a singleton instance in the global namespace
export const portfolioService = new PortfolioService();
