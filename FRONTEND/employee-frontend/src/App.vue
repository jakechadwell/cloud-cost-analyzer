<template>
  <div class="container">
    <div class="navbar-header text-center">
      <a class="navbar-brand" href="/employees">
        View Employees
        </a>
      <a class="navbar-brand" href="/trainings">
        View Trainings
      </a><br/><br/>
    </div>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: "app"
};
</script>

<style>
@import 
url(https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css);

</style>