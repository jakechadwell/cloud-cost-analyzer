<template>
  <b-navbar toggleable type="dark" variant="dark" class="p-4 mb-5">
    <b-navbar-brand href="/">Eclipse Training App </b-navbar-brand>

    

    <!-- <b-collapse id="navbar-toggle-collapse" is-nav> -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item href="/employees" class="p-2">Employees</b-nav-item>
        <b-nav-item href="/trainings" class="p-2">Trainings</b-nav-item>
        <b-nav-item href="/attributes" class="p-2">Attributes</b-nav-item>
      </b-navbar-nav>
    <!-- </b-collapse> -->
  </b-navbar>
</template>

<style>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

  nav-link{
    font-size: small;
  }

  .navbar-nav{
    flex-direction: row!important;
  }

  nav {
    font-family: "Futura XBlk BT", sans-serif;
  }
</style>