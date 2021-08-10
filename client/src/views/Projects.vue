<template>
<b-container>
    <div v-if="loading" class="d-flex justify-content-center mt-5">
        <b-icon icon="circle-fill" animation="throb" font-scale="2" variant="primary"></b-icon>
    </div>
    <div v-else>
        <h2>Pinned</h2>
        <repository class="mb-3" v-for="repo in userData.pinnedItems.nodes" :key="repo.url" :repository="repo"/>

        <h2>Recent</h2>
        <repository class="mb-3" v-for="repo in userData.repositoriesContributedTo.nodes" :key="repo.url" :repository="repo"/>
    </div>
</b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Repository from '../components/Repository.vue'
import { getUserDetails } from '../services/github'

@Component({
    components: {
        Repository
    }
})
export default class Projects extends Vue {

    private userData = {}
    private loading = true

    async created() {
        this.userData = await getUserDetails()
        console.log(this.userData)
        this.loading = false
    }
}
</script>