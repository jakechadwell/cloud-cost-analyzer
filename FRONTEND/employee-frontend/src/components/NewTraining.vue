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
            <label>Training ID</label>
            <input type="text" class="form-control"
             v-model ="trainingid" />
          </fieldset>
          <fieldset class="form-group">
            <label>Cloud ID</label>
            <input type="text" class="form-control"
             v-model ="cloudid" />
          </fieldset>
          <fieldset class="form-group">
            <label>Training Detail</label>
            <input type="text" class="form-control"
             v-model ="trainingDetail" />
          </fieldset>
          <fieldset class="form-group">
            <label>Training Path</label>
            <input type="text" class="form-control"
             v-model="trainingPath" />
          </fieldset>
          <fieldset class="form-group">
            <label>Training Dates</label>
            <input type="text" class="form-control"
             v-model="trainingDates" />
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
        trainingid: "",
        cloudid: "",
        trainingDetail: "",
        trainingPath: "",
        trainingDates: "",
        errors: [],
      };
    },
    methods: {
      refreshTrainingDetails() {
        TrainingDataService.retrieveTraining(this.trainingid).then((res) => {
          this.trainingid = res.data.trainingid;
          this.cloudid = res.data.cloudid;
          this.trainingDetail = res.data.trainingDetail;
          this.trainingPath = res.data.trainingPath;
          this.trainingDates = res.data.trainingDates;
        });
      },
      validateAndSubmit(e) {
        e.preventDefault();
        this.errors = [];
        if (!this.trainingDetail) {
          this.errors.push("Enter valid values");
        } else if (this.trainingDetail.length < 2) {
          this.errors.push
          ("Enter atleast 2 characters in Training Detail.");
        }
        if (this.errors.length === 0) {
            TrainingDataService.createTraining({
              trainingid: this.trainingid,
              cloudid: this.cloudid,
              trainingDetail: this.trainingDetail,
              trainingPath: this.trainingPath,
              trainingDates: this.trainingDates
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