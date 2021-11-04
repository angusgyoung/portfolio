<template>
  <b-navbar type="dark" variant="dark">
    <b-navbar-nav>
      <b-nav-item to="/" active>Home</b-nav-item>
      <b-nav-item to="/projects">Projects</b-nav-item>
    </b-navbar-nav>

    <b-navbar-nav class="ml-auto">
      <b-nav-item
        ><a href="github.com/angusgyoung">
          <b-icon icon="github" variant="info" /> </a
      ></b-nav-item>
      <b-nav-item>
        <a href="https://www.linkedin.com/in/angus-young-278b886b/"
          ><b-icon icon="linkedin" variant="info"
        /></a>
      </b-nav-item>
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
