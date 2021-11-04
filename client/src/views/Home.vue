<template>
  <div id="home">
    <b-container class="py-4">
      <b-row>
        <b-col>
          <b-img
            :src="user.avatarUrl"
            :alt="user.avatarUrl"
            rounded="circle"
            center
            fluid
            id="profile-image"
            class="my-5"
          />
        </b-col>
      </b-row>

      <b-row class="text-center">
        <b-col>
          <h2>{{ user.name }}</h2>
          <h3 class="text-muted">{{ user.login }}</h3>
          <p
            :style="'color:' + (user.isHireable ? '#00ff0d' : '#c637ff') + ';'"
          >
            {{ hireMessage() }}
          </p>
          <h4>
            {{ user.bio }}
          </h4>
        </b-col>
      </b-row>

      <b-row>
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
    </b-container>

    <b-jumbotron text-variant="dark" fluid>
      <p class="lead">
        Skilled developer with experience implementing highly scalable and
        diverse software solutions, reacting dynamically to changing
        requirements and working with multidiciplinary teams. I continually
        challenge myself by improving my skills and exploring new technologies.
      </p>
    </b-jumbotron>

    <b-jumbotron
      bg-variant="dark"
      lead=" I have experience working with the following tools and technologies"
      class="mb-0"
      fluid
    >
      <b-row>
        <b-col>
          <ul>
            <li>Java</li>
            <li>Go</li>
            <li>Vue</li>
            <li>Javascript/Typescript</li>
            <li>HTML/CSS</li>
            <li>MySQL</li>
            <li>MongoDB</li>
            <li>Elasticsearch</li>
            <li>
              Automated Deployment and Testing (Jenkins, Checkmarx, Cucumber)
            </li>
          </ul>
        </b-col>
        <b-col>
          <ul>
            <li>ActiveMQ/RabbitMQ</li>
            <li>Linux</li>
            <li>Git</li>
            <li>Kubernetes</li>
            <li>
              AWS
              <ul>
                <li>IAM</li>
                <li>EKS</li>
                <li>ECR</li>
                <li>EC2</li>
                <li>RDS</li>
              </ul>
            </li>
          </ul>
        </b-col>
      </b-row>
    </b-jumbotron>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { User, Organisation } from '../types'
import { getUser } from '../services/api'

@Component({
  components: {}
})
export default class Home extends Vue {
  private user = {} as User
  private organisations = [] as Organisation[]

  async created(): Promise<void> {
    this.user = await getUser()

    if (this.user.organizations?.nodes)
      this.organisations = this.user?.organizations?.nodes
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
}

#profile-image {
  width: 18rem;
}

.org-img {
  width: 3rem;
}
</style>
