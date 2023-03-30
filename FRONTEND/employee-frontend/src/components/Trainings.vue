<template>
    <div class="container">
      <h3>All Trainings</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="container">
        <table class="table">
          <thead>
            <tr> 
              <th>Category</th>
              <th>Course Name</th>
              <th>Course ID</th>
              <th>Update</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="training in trainings" v-bind:key="training.courseCode">
            
              <td>{{ training.category }}</td>
              <td>{{ training.courseName }}</td>
              <td>{{ training.courseCode }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateTraining(training.courseCode)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteTraining(training.courseCode)">
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
          console.log(this.trainings)
        });
      },
      addTraining() {
        this.$router.push(`/training/new`);
      },
      updateTraining(courseCode) {
        this.$router.push(`/training/${courseCode}`);
      },
      deleteTraining(courseCode) {
        TrainingDataService.deleteTraining(courseCode).then(() => {
          this.refreshTrainings();
        });
      },
    },
    created() {
      this.refreshTrainings();
    },
  };
  </script>