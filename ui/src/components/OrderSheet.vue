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

    <div class="field is-grouped">
    <div class="control">
        <button class="button is-link">Submit</button>
    </div>
    <div class="control">
        <button class="button is-text">Cancel</button>
    </div>
    </div>
</div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import { Order } from '../order';

@Component
export default class OrderSheet extends Vue {
    @Prop() order!: Order;    
}
</script>