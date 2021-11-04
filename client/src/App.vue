<template>
  <div id="app">
    <navigation />
    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
    <footer id="footer" class="mt-auto py-4 text-center">
      Angus Young 2021
    </footer>
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

<style lang="scss">
@import 'main.scss';

#app {
  height: 100vh;
}

#footer {
  background-color: $dark-grey;
  color: white;
}
</style>
