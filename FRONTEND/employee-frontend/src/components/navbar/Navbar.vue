<template>
  <b-navbar toggleable type="dark" variant="dark" class="p-4">
    <b-navbar-brand href="/">Eclipse Training App </b-navbar-brand>

      <b-navbar-nav v-if="currentUser" class="ml-auto">
        <b-nav-item href="/dashboard" class="p-2">Dashboard</b-nav-item>
        <b-nav-item @click.prevent="accountRedirect(currentUser.data.email)" class="p-2">Account</b-nav-item>
        <b-nav-item @click.prevent="logOut" class="p-2">Logout</b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav v-else class="ml-auto">
        <b-nav-item href="/signin" class="p-2">Login</b-nav-item>
        <b-nav-item href="/signup" class="p-2">Register</b-nav-item>
      </b-navbar-nav>

      
  </b-navbar>
</template>

<script>
  export default {
    name: "navbar",
    computed: {
      currentUser(){
        return this.$store.state.auth.user;
      }
    },
    methods: {
      logOut(){
        this.$store.dispatch('auth/logout');
        this.$router.push('/signin')
      },
      accountRedirect(email){
        this.$router.push(`/account/${email}`)
      },
      log(email){
        console.log(email)
      }
    }
  }

</script>

<style>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');
@import url('https://fonts.cdnfonts.com/css/fox-corporation-logo');
  body{
    font-size: 1.25em !important;
  }

  .navbar-nav{
    flex-direction: row!important;
  }

  .navbar-nav .nav-link{
    padding: 0 .5em 0 .5em !important;
  }

  nav {
    font-family: "Futura XBlk BT", sans-serif;
  }
</style>