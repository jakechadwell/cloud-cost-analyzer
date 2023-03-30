<template>
    <div>
      <h3>New Training</h3>
      <div class="container">
        <form @submit="validateAndSubmit">
          <div v-if="errors.length">
            <div
              class="alert alert-danger"
              v-bind:key="index"
              v-for="(error, index) in errors"
            >
              {{ error }}
            </div>
          </div>
          <fieldset class="form-group">
            <label>Category</label>
            <input type="text" class="form-control"
             v-model ="category" />
          </fieldset>
          <fieldset class="form-group">
            <label>Course Name</label>
            <input type="text" class="form-control"
             v-model="courseName" />
          </fieldset>
          <fieldset class="form-group">
            <label>Course Code</label>
            <input type="text" class="form-control"
             v-model="courseCode" />
          </fieldset>
          <button class="btn btn-success" 
          type="submit">Save</button>
        </form>
      </div>
    </div>
  </template>
  <script>
  import TrainingDataService from "../service/TrainingDataService";
  
  export default {
    name: "Training",
    data() {
      return {
        category: "",
        courseName: "",
        courseCode: "",
        errors: [],
      };
    },
    methods: {
      refreshTrainingDetails() {
        TrainingDataService.retrieveTraining(this.coursecode).then((res) => {
          this.category = res.data.category;
          this.courseName = res.data.courseName;
          this.courseCode = res.data.courseCode;
        });
      },
      validateAndSubmit(e) {
        e.preventDefault();
        this.errors = [];
        if (!this.category) {
          this.errors.push("Enter valid values");
        } else if (this.category.length < 2) {
          this.errors.push
          ("Enter atleast 2 characters in Category.");
        }
        if (!this.courseName) {
          this.errors.push("Enter valid values");
        } else if (this.courseName.length < 4) {
          this.errors.push
          ("Enter atleast 4 characters in Course Name");
        }
        if (this.errors.length === 0) {
            TrainingDataService.createTraining({
              category: this.category,
              courseName: this.courseName,
              courseCode: this.courseCode,
            }).then(() => {
              this.$router.push("/trainings");
            });
        }
      },
    },
    created() {
      this.refreshTrainingDetails();
    }
  };
  </script>