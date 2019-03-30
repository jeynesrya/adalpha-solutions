<template>
    <div class="section">
        <div class="columns">
            <div class="column">
                First column
            </div>
            <div class="column">
                <div class="field">
                    <div class="control">
                        <input class="input is-primary" type="text" placeholder="Primary input">
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
    protected msg = 'Hello, Zoe!';
}
</script>

<style>
</style>
