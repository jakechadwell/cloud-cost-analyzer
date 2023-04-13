<template>
  <b-navbar toggleable type="dark" variant="dark" class="p-4 mb-5">
    <b-navbar-brand href="/">Eclipse Training App </b-navbar-brand>

    <b-navbar-toggle target="navbar-toggle-collapse">
      <template #default="{ expanded }">
        <!-- <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon> -->
        <i v-if="expanded" class="bi bi-chevron-double-up"></i>
        <!-- <b-icon v-else icon="chevron-bar-down"></b-icon> -->
        <i v-else class="bi bi-chevron-double-down"></i>
      </template>
    </b-navbar-toggle>

    <b-collapse id="navbar-toggle-collapse" is-nav>
      <b-navbar-nav class="ml-auto">
        <b-nav-item href="/employees" class="p-2">Employees</b-nav-item>
        <b-nav-item href="/trainings" class="p-2">Trainings</b-nav-item>
        <b-nav-item href="/attributes" class="p-2">Attributes</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>

<style>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');
@import url('https://fonts.cdnfonts.com/css/made-soulmaze');

  nav-link{
    font-size: small;
  }

  nav {
    font-family: "Futura XBlk BT", sans-serif;
    /* font-family: 'MADE Soulmaze', sans-serif; */
    /* font-family: 'MADE Soulmaze Outline', sans-serif; */
  }
</style>