<template>
  <b-container id="home" fluid>
    <b-row>
      <b-col>
        <b-img
          :src="user.avatarUrl"
          :alt="user.avatarUrl"
          rounded="circle"
          center
          fluid
        />
      </b-col>
    </b-row>
    <b-row class="text-center">
      <b-col>
        <h1>{{ user.name }}</h1>
        <h3 class="text-muted">{{ user.login }}</h3>
        <p :style="'color:' + (user.isHireable ? '#00ff0d' : '#c637ff') + ';'">
          {{ hireMessage() }}
        </p>
        <p>
          {{ user.bio }}
        </p>
      </b-col>
    </b-row>

    <b-row align-h="center">
      <b-col>
        <div class="d-flex justify-content-center">
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
    </b-row>

    <footer class="fixed-bottom">
      <div class="d-flex justify-content-center">
        <a :href="user.url">
          <b-icon class="h2 m-3" icon="github" variant="info" />
        </a>
        <b-icon class="h2 m-3" icon="linkedin" variant="info" />
      </div>
    </footer>
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
  private user!: User = {}
  private organisations!: Organisation[] = []

  async created(): Promise<void> {
    this.user = await getUser()

    this.organisations = this.user.organizations.nodes
  }

  hireMessage(): string {
    if (this.user.isHireable) {
      return 'Available'
    } else return 'Currently employed ' + this.user.company
  }
}
</script>

<style lang="scss" scoped>
@import '../main.scss';

#home {
  background-color: $dark;
  color: $white;
  height: 100vh;
  padding-top: 3rem;
}

.org-img {
  width: 3rem;
}
</style>
