<template>
<div class="form">
    <h1 class="title">Order Sheet</h1>

    <div class="field">
    <label class="label">Isin</label>
    <div class="control">
        <input class="input" type="text" v-model="order.Isin" disabled>
    </div>
    </div>

    <div class="field">
    <label class="label">Type</label>
    <div class="control">
        <input class="input" type="text" v-model="order.Type" disabled/>
    </div>
    </div>

    <div class="field">
    <label class="label">Amount</label>
    <div class="control">
        <input class="input" type="text" v-model="order.Amount">
    </div>
    </div>

    <div class="field">
    <label class="label">Currency</label>
    <div class="control">
        <input class="input" type="text" v-model="order.Currency">
    </div>
    </div>

    <div class="field is-grouped" v-if="showProgress !== true">
    <div class="control">
        <button class="button is-link" @click="submit()">Submit</button>
    </div>
    </div>
    <div class="column is-12" v-else>
        <progress class="progress is-small is-primary" max="100"></progress>
    </div>

    <div class="field" v-if="showResponse === true">
        <p>{{response.message}}</p>
    </div>
</div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import { UnitOrder, CurrencyOrder, InternalOrder, Order } from '../order';
import { portfolioService } from '../portfolio.service';

@Component
export default class OrderSheet extends Vue {
    @Prop() order!: InternalOrder;

    protected showProgress = false;
    protected showResponse = false;
    protected response = {"code": 0, "message": ""};

    private submit() {
        this.showProgress = !this.showProgress;
        let orderToSubmit: Order;

        if (this.order.Currency === "N/A") {
            orderToSubmit = {"Isin": this.order.Isin, "Amount": Number(this.order.Amount)}
        } else {
            orderToSubmit = {"Isin": this.order.Isin, "Amount": Number(this.order.Amount), "Currency": this.order.Currency}
        }

        return portfolioService.updatePortfolio(orderToSubmit, this.order.Type.toLowerCase())
                                .then((response) => {
                                    this.response.code = response.status;
                                    this.response.message = response.data;
                                    this.showResponse = true;
                                    this.showProgress = false;
                                });
    }
}
</script>