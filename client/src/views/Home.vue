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
      <b-row class="mt-5" align-v="center">
        <b-col>
          <h1>{{ user.name }}</h1>
          <h3 class="text-muted">{{ user.login }}</h3>
          <p
            :style="'color:' + (user.isHireable ? '#00ff0d' : '#c637ff') + ';'"
          >
            {{ hireMessage() }}
          </p>
          <p>
            {{ user.bio }}
          </p>
          <p>Take a look at some of my projects.</p>
          <div>
            <b-button to="/projects" variant="outline-info">Projects</b-button>
          </div>
          <div class="mt-3">
            <a
              v-for="organisation in organisations"
              :key="organisation.login"
              :href="organisation.url"
            >
              <b-img
                class="org-img m-2"
                :src="organisation.avatarUrl"
                :alt="organisation.name"
                rounded="circle"
              />
            </a>
          </div>
        </b-col>
        <b-col>
          <b-img
            :src="user.avatarUrl"
            :alt="user.avatarUrl"
            rounded="circle"
            fluid
          />
        </b-col>
      </b-row>
      <b-row align-h="center" align-v="end" class="fixed-bottom">
        <b-col>
          <div class="d-flex justify-content-center">
            <a :href="user.url">
              <b-icon class="h2 m-3" icon="github" variant="primary" />
            </a>
            <b-icon class="h2 m-3" icon="linkedin" variant="primary" />
          </div>
        </b-col>
      </b-row>
    </div>
  </b-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { User, Organisation } from '../types'
import { getUser } from '../services/api'

@Component({
  components: {}
})
export default class Home extends Vue {
  private user!: User
  private organisations!: Organisation[]
  private loading = true

  async created(): Promise<void> {
    this.user = await getUser()

    this.organisations = this.user.organizations.nodes
    this.loading = false
  }

  hireMessage(): string {
    if (this.user.isHireable) {
      return 'Available'
    } else return 'Currently employed ' + this.user.company
  }
}
</script>

<style lang="scss">
@import '../main.scss';

.org-img {
  width: 3rem;
}
</style>
