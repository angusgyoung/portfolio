<template>
  <b-card
    :title="repository.name"
    :sub-title="repository.primaryLanguage.name"
    class="repository-card"
  >
    <b-card-text>
      {{ repository.description }}
    </b-card-text>
    <b-row>
      <div v-if="topics && topics.length">
        <b-col>
          <b-badge
            class="mr-2"
            v-for="topic in topics"
            :key="topic.name"
            pill
            variant="primary"
            >{{ topic.topic.name }}</b-badge
          >
        </b-col>
      </div>
      <div v-if="languages && languages.length">
        <b-col>
          <b-badge
            class="mr-2"
            v-for="language in languages"
            :key="language.name"
            pill
            :style="'background-color: ' + language.color"
            >{{ language.name }}</b-badge
          >
        </b-col>
      </div>
      <b-col>
        <b-button
          class="btn-sm float-right"
          :href="repository.url"
          variant="outline-info"
          >View on Github <b-icon icon="github"
        /></b-button>
      </b-col>
    </b-row>
  </b-card>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { Topic, Language, Repository as Repo } from '../types'

@Component({
  components: {}
})
export default class Repository extends Vue {
  @Prop({ default: {} })
  repo!: Repo

  private topics!: Topic[]
  private languages!: Language[]

  created(): void {
    this.topics = this.repo.repositoryTopics.nodes
    this.languages = this.repo.languages.nodes
  }
}
</script>

<style lang="scss">
@import '../main.scss';

.repository-card {
  background-color: $navy;
}
</style>
