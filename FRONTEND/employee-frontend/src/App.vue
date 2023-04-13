<template>
  <div>
    <Navbar/>
    <div class="container">
        <router-view/>
    </div>
  </div>
</template>

<script>
import Navbar from './components/navbar/Navbar.vue';

export default {
  name: "app",
  components: {
    Navbar
  }
};
</script>

<style>
@import 
url(https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css);
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

  h3{
    text-align: center;
    font-family: "Futura XBlk BT", sans-serif;
    margin-bottom: 1.5em;
  }

  

</style>