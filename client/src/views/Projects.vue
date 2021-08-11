<template>
<b-container>
    <div v-if="loading" class="d-flex justify-content-center mt-5">
        <b-icon icon="circle-fill" animation="throb" font-scale="2" variant="primary"></b-icon>
    </div>
    <div v-else>
        <h2>Pinned</h2>
        <repository class="mb-3" v-for="repo in pinned" :key="repo.url" :repository="repo"/>

        <h2>Recent</h2>
        <repository class="mb-3" v-for="repo in recent" :key="repo.url" :repository="repo"/>
    </div>
</b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Repository from '../components/Repository.vue'
import { getProjects } from '../services/github'

@Component({
    components: {
        Repository
    }
})
export default class Projects extends Vue {

    private projects
    private pinned
    private recent
    private loading = true

    async created() {
        this.projects = await getProjects()
        this.pinned = this.projects.pinnedItems.nodes
        this.recent = this.projects.repositoriesContributedTo.nodes
        this.loading = false
    }
}
</script>