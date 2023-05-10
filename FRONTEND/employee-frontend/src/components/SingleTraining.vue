<template>
    <div >

    
    <div class="training-body" v-for="training in trainings" v-bind:key="training.trainingid">
        <div class="thumbnail">

        </div>
        <div class="training-details">
            <h3 class="cloud">{{ training.cloudid }}</h3>
            <p class="details">{{ training.trainingDetail }}</p>
        </div>
        <div class="path">
            <h3>{{ training.trainingPath }}</h3>
        </div>
        <div class="dates">
            <p>{{ training.trainingDates }}</p>
        </div>
        <div class="view-btn">
            <button>View Course</button>
        </div>
    </div>
    </div>
</template>

<script>
import TrainingDataService from '../service/TrainingDataService';

    export default {
    name: "SingleTraining",
    data() {
      return {
        trainings: [],
        message: "",
      };
    },
    methods: {
      refreshTrainings() {
        TrainingDataService.retrieveAllTrainings.then((res) => {
          this.trainings = res.data;
        });
      }
    },
    created() {
      this.refreshTrainings();
    },
  };
</script>

<style scoped>

</style>