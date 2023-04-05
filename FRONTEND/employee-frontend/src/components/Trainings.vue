<template>
    <div class="container">
      <h3>All Trainings</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="container">
        <table class="table">
          <thead>
            <tr> 
              <th>Training ID</th>
              <th>Cloud ID</th>
              <th>Training Detail</th>
              <th>Training Path</th>
              <th>Training Dates</th>
              <th>Update</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="training in trainings" v-bind:key="training.trainingid">
            
              <td>{{ training.trainingid }}</td>
              <td>{{ training.cloudid }}</td>
              <td>{{ training.trainingDetail }}</td>
              <td>{{ training.trainingPath }}</td>
              <td>{{ training.trainingDates }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateTraining(training.trainingid)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteTraining(training.trainingid)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="row">
          <button class="btn btn-success" 
          v-on:click="addTraining()">Add</button>
        </div>
      </div>
    </div>
  </template>
  <script>
  import axios from "axios";
  import TrainingDataService from "../service/TrainingDataService";
  
  export default {
    name: "Trainings",
    data() {
      return {
        trainings: [],
        message: "",
      };
    },
    methods: {
      refreshTrainings() {
        axios.get("http://localhost:9080/trainings").then((res) => {
          this.trainings = res.data;
        });
      },
      addTraining() {
        this.$router.push(`/training/new`);
      },
      updateTraining(trainingid) {
        this.$router.push(`/training/${trainingid}`);
      },
      deleteTraining(trainingid) {
        TrainingDataService.deleteTraining(trainingid).then(() => {
          this.refreshTrainings();
        });
      },
    },
    created() {
      this.refreshTrainings();
    },
  };
  </script>