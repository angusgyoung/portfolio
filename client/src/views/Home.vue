<template>
  <b-container>
    <b-row align-v="center">
      <b-col class="mt-5 p-5">
        <h1>Angus Young</h1>
        <h3>Software Engineer</h3>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
          sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
          Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris 
          nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla 
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in 
          culpa qui officia deserunt mollit anim id est laborum.
        </p>
        <p>
          Take a look at some of my projects. 
        </p>
        <div>
          <b-button class="mr-2" variant="outline-info">Projects</b-button>
          <b-button class="mr-2" variant="outline-info">Gists</b-button>
        </div>
      </b-col>
    </b-row>
    <b-row align-h="center" align-v="end" class="fixed-bottom">
      <b-col>
        <div class="d-flex justify-content-center">
          <b-icon class="h2 m-3" icon="github" variant="primary"/>
          <b-icon class="h2 m-3" icon="linkedin" variant="primary"/>
        </div>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class Home extends Vue {}
</script>

<style lang="scss">
@import '../main.scss';



</style>
