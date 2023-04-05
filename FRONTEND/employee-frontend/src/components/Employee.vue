<template>
    <div>
      <h3>Employee</h3>
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
            <label>First Name</label>
            <input type="text" class="form-control"
             v-model ="firstName" />
          </fieldset>
          <fieldset class="form-group">
            <label>Last Name</label>
            <input type="text" class="form-control"
             v-model="lastName" />
          </fieldset>
          <fieldset class="form-group">
            <label>Dept</label>
            <input type="text" class="form-control"
             v-model="dept" />
          </fieldset>
          <fieldset class="form-group">
            <label>Cloud</label>
            <input type="text" class="form-control"
             v-model="cloud" />
          </fieldset>
          <fieldset class="form-group">
            <label>Training Attended</label>
            <input type="text" class="form-control"
             v-model="trainingAttended" />
          </fieldset>
          <fieldset class="form-group">
            <label>Training Path</label>
            <input type="text" class="form-control"
             v-model="trainingPath" />
          </fieldset>
          <fieldset class="form-group">
            <label>Email</label>
            <input type="text" class="form-control"
             v-model="email" />
          </fieldset>
          <fieldset class="form-group">
            <label>Infographics</label>
            <input type="text" class="form-control"
             v-model="infographics" />
          </fieldset>
          <button class="btn btn-success" 
          type="submit">Save</button>
        </form>
      </div>
    </div>
  </template>
  <script>
  import EmployeeDataService from "../service/EmployeeDataService";
  
  export default {
    name: "Employee",
    data() {
      return {
        firstName: "",
        lastName: "",
        dept: "",
        cloud: "",
        trainingAttended: "",
        trainingPath: "",
        email: "",
        infographics: "",
        errors: [],
      };
    },
    computed: {
      employeeid() {
        return this.$route.params.employeeid;
      },
    },
    methods: {
      refreshEmployeeDetails() {
        EmployeeDataService.retrieveEmployee(this.employeeid).then((res) => {
          this.firstName = res.data.firstName;
          this.lastName = res.data.lastName;
          this.dept = res.data.dept;
          this.cloud = res.data.cloud;
          this.trainingAttended = res.data.trainingAttended;
          this.trainingPath = res.data.trainingPath;
          this.email = res.data.email;
          this.infographics = res.data.infographics;
        });
      },
      validateAndSubmit(e) {
        e.preventDefault();
        this.errors = [];
        if (!this.firstName) {
          this.errors.push("Enter valid values");
        } else if (this.firstName.length < 4) {
          this.errors.push
          ("Enter atleast 4 characters in First Name");
        }
        if (!this.lastName) {
          this.errors.push("Enter valid values");
        } else if (this.lastName.length < 4) {
          this.errors.push
          ("Enter atleast 4 characters in Last Name");
        }
        if (this.errors.length === 0) {
            EmployeeDataService.updateEmployee(this.employeeid, {
                employeeid: this.employeeid,
                firstName: this.firstName,
                lastName: this.lastName,
                dept: this.dept,
                cloud: this.cloud,
                trainingAttended: this.trainingAttended,
                trainingPath: this.trainingPath,
                email: this.email,
                infographics: this.infographics
            }).then(() => {
              this.$router.push("/employees");
            });
          }
        }
      },
    created() {
      this.refreshEmployeeDetails();
    }
  };
  </script>