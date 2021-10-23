<template>
  <div id="app">
    <navigation />
    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Navigation from './components/Navigation.vue'

@Component({
  components: {
    Navigation
  }
})
export default class App extends Vue {}
</script>

<style lang="scss"></style>
