<template>
  <b-navbar type="dark" variant="dark">
    <b-navbar-nav>
      <b-nav-item to="/" active>Home</b-nav-item>
      <b-nav-item to="/projects">Projects</b-nav-item>
    </b-navbar-nav>
  </b-navbar>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class Navigation extends Vue {}
</script>
