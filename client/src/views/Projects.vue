<template>
  <b-container>
    <div v-if="loading" class="d-flex justify-content-center mt-5">
      <b-icon
        icon="circle-fill"
        animation="throb"
        font-scale="2"
        variant="primary"
      ></b-icon>
    </div>
    <div v-else>
      <h2>Pinned</h2>
      <repository
        class="mb-3"
        v-for="repo in pinned"
        :key="repo.url"
        :repo="repo"
      />

      <h2>Recent</h2>
      <repository
        class="mb-3"
        v-for="repo in recent"
        :key="repo.url"
        :repo="repo"
      />
    </div>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Repository from '../components/Repository.vue'
import { Repository as Repo } from '../types'
import { getProjects } from '../services/api'

@Component({
  components: {
    Repository
  }
})
export default class Projects extends Vue {
  private pinned!: Repo[]
  private recent!: Repo[]
  private loading = true

  async created(): Promise<void> {
    const user = await getProjects()
    this.pinned = user.pinnedItems.nodes
    this.recent = user.repositoriesContributedTo.nodes
    this.loading = false
  }
}
</script>
