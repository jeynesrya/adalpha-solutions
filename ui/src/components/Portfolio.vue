<template>
<div class=columns>
    <div class="column is-1"></div>
    <div class="column is-7">
    <h1 class="title">Portfolio</h1>
    <table class="table" v-if="showTable === true">
        <thead>
            <tr>
                <th>Isin</th>
                <th>Units</th>
                <th>Current Price (GBP)</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="p in portfolio" :key="p.Isin">
                <td>{{p.Isin}}</td>
                <td>{{p.Units}}</td>
                <td>{{p.CurrentWorth}}
                <td>
                    <a class="button is-info" @click="Buy(p)">Buy</a>&nbsp
                    <a class="button is-success" @click="Sell(p)">Sell</a>&nbsp
                    <a class="button is-warning" @click="Invest(p)">Invest</a>&nbsp
                    <a class="button is-danger" @click="Raise(p)">Raise</a>&nbsp
                </td>
            </tr>
        </tbody>
    </table>
    <div class="column is-12" v-else>
        <p>{{loadingMsg}}</p>
        <progress class="progress is-small is-primary" max="100"></progress>
    </div>
    </div>
    <div class="column is-3">
        <OrderSheet :order="order" v-on:update-values="getPortfolio()"/>
    </div>
</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { PortfolioItem } from '../portfolio';
import { portfolioService } from '../portfolio.service';
import { InternalOrder } from '../order';
import OrderSheet from '@/components/OrderSheet.vue';


@Component({components: {
    OrderSheet,
  },
})
export default class Portfolio extends Vue {
    private portfolio: PortfolioItem[] = [];
    private order: InternalOrder = {Isin: '', Type: '', Amount: 0, Currency: ''};
    private showTable = false;
    private loadingMsg = 'Loading data... Establishing current price of your asset(s)';

    private created() {
        this.getPortfolio();
    }
    private getPortfolio() {
        this.showTable = false;
        this.portfolio = [];
        return portfolioService.getPortfolio().then((response) => {
            this.portfolio = response.data;
            this.showTable = true;
        });
    }
    private Buy(item: PortfolioItem) {
        this.order = {Isin: item.Isin, Type: 'Buy', Amount: 0, Currency: 'N/A'};
    }
    private Sell(item: PortfolioItem) {
        this.order = {Isin: item.Isin, Type: 'Sell', Amount: 0, Currency: 'N/A'};
    }
    private Invest(item: PortfolioItem) {
        this.order = {Isin: item.Isin, Type: 'Invest', Amount: 0, Currency: 'GBP'};
    }
    private Raise(item: PortfolioItem) {
        this.order = {Isin: item.Isin, Type: 'Raise', Amount: 0, Currency: 'GBP'};
    }
}
</script>