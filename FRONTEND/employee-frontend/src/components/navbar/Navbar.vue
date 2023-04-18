<template>
  <b-navbar toggleable type="dark" variant="dark" class="p-4 mb-5">
    <b-navbar-brand href="/">Eclipse Training App </b-navbar-brand>

    

    <!-- <b-collapse id="navbar-toggle-collapse" is-nav> -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item href="/signin" class="p-2">Login</b-nav-item>
        <b-nav-item href="/signup" class="p-2">Register</b-nav-item>
      </b-navbar-nav>
    <!-- </b-collapse> -->
  </b-navbar>
</template>

<style>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

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