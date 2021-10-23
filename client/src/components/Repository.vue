<template>
  <b-container class="p-5" fluid>
    <b-container>
      <b-row align-h="between">
        <b-col>
          <h2>{{ repo.name }}</h2>
          <h5 class="text-muted">
            Last updated
            {{
              new Date(repo.updatedAt).toLocaleDateString('en-gb', {
                year: 'numeric',
                month: 'long',
                day: 'numeric'
              })
            }}
          </h5>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <p>{{ repo.description }}</p>
        </b-col>
      </b-row>
      <b-row>
        <b-col v-if="languages.length > 0">
          <b-badge
            class="mr-1"
            v-for="language in languages"
            :key="language.name"
            pill
            :style="'background-color: ' + language.color"
            >{{ language.name }}</b-badge
          >
        </b-col>
      </b-row>
      <b-row>
        <b-col v-if="topics.length > 0">
          <b-badge
            class="mr-1"
            v-for="topic in topics"
            :key="topic.name"
            pill
            variant="primary"
            >{{ topic.topic.name }}</b-badge
          >
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-button
            class="mt-1"
            size="sm"
            :href="repo.url"
            variant="outline-info"
            >View on GitHub <b-icon icon="github" />
          </b-button>
        </b-col>
      </b-row>
    </b-container>
  </b-container>
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
