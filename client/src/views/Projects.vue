<template>
  <div>
    <repository
      :class="'mb-3 ' + (index % 2 == 0 ? 'filled' : '')"
      v-for="(repo, index) in repositories"
      :key="repo.url"
      :repo="repo"
    />
  </div>
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
  private repositories = [] as Repo[]

  async created(): Promise<void> {
    const user = await getProjects()

    if (user?.pinnedItems?.nodes) {
      this.repositories = this.repositories.concat(
        user.pinnedItems.nodes.map(
          (repositoryEdge) => repositoryEdge.repository
        )
      )
    }

    // add repositories contributed to, excluding pinned
    if (user?.repositoriesContributedTo?.nodes) {
      this.repositories = this.repositories.concat(
        user.repositoriesContributedTo.nodes
          .map((repositoryEdge) => repositoryEdge.repository)
          .filter((repo) => !this.repositories.some((r) => r.url === repo.url))
      )
    }
  }
}
</script>

<style lang="scss">
@import '../main.scss';

.filled {
  background-color: $off-white;
}
</style>
