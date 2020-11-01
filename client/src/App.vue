<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
      <span v-if="$auth.isAuthenticated"> |
        <router-link to="/profile">Profile</router-link>
      </span>
      <div v-if="!$auth.loading">
        <!-- show login when not authenticated -->
        <a href="#" v-if="!$auth.isAuthenticated" v-on:click.prevent="login">Log in</a>
        <!-- show logout when authenticated -->
        <a href="#" v-if="$auth.isAuthenticated" v-on:click.prevent="logout">Log out</a>
      </div>
    </div>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
