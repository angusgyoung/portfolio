<template>
    <b-nav>
        <b-nav-item active>Home</b-nav-item>
        <b-nav-item>Projects</b-nav-item>
        <b-nav-item>Resume</b-nav-item>
        <b-nav-item>Gists</b-nav-item>
    </b-nav>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component({
    components: {}
})
export default class Navigation extends Vue {}
</script>