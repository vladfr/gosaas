<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
     <div v-if="!$auth.loading">
      <!-- show login when not authenticated -->
      <button v-if="!$auth.isAuthenticated" @click="login">Log in</button>
      <!-- show logout when authenticated -->
      <button v-if="$auth.isAuthenticated" @click="logout">Log out</button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
  components: {
  },
  methods: {
    // Log the user in
    login() {
      this.$auth.loginWithRedirect({})
    },
    // Log the user out
    logout() {
      this.$auth.logout({
        returnTo: window.location.origin
      })
    },
  }
})
export default class Home extends Vue {}
</script>
