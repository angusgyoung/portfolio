<template>
  <div>
    <repository
      :class="'mb-3 ' + (index % 2 == 0 ? 'filled' : '')"
      v-for="(repo, index) in pinned"
      :key="repo.url"
      :repo="repo"
    />
    <repository
      :class="'mb-3 ' + (index % 2 == 0 ? 'filled' : '')"
      v-for="(repo, index) in recent"
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
  private pinned: Repo[] = []
  private recent: Repo[] = []

  async created(): Promise<void> {
    const user = await getProjects()
    this.pinned = user.pinnedItems.nodes
    this.recent = user.repositoriesContributedTo.nodes
  }
}
</script>

<style lang="scss">
@import '../main.scss';

.filled {
  background-color: $off-white;
}
</style>
