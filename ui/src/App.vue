<template>
  <div>
      <nav class="navbar has-shadow">
            <div class="section">
                <div class="navbar-brand">
                  <h4 class="title is-2">AdAlpha Solutions</h4>
                </div>
            </div>
      </nav>
    <HelloWorld/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from '@/components/HelloWorld.vue';

@Component({
  components: {
    HelloWorld,
  },
})
export default class App extends Vue {}
</script>

<style>
</style>
