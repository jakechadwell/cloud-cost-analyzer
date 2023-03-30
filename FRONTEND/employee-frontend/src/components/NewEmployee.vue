<template>
    <div>
      <h3>New Employee</h3>
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
            <label>Middle Name</label>
            <input type="text" class="form-control"
             v-model="middleName" />
          </fieldset>
          <fieldset class="form-group">
            <label>Email ID</label>
            <input type="text" class="form-control"
             v-model="emailid" />
          </fieldset>
          <fieldset class="form-group">
            <label>SAP ID</label>
            <input type="text" class="form-control"
             v-model="sapid" />
          </fieldset>
          <fieldset class="form-group">
            <label>Organization</label>
            <input type="text" class="form-control"
             v-model="organization" />
          </fieldset>
          <fieldset class="form-group">
            <label>Role</label>
            <input type="text" class="form-control"
             v-model="role" />
          </fieldset>
          <fieldset class="form-group">
            <label>Reporting To</label>
            <input type="text" class="form-control"
             v-model="reportingto" />
          </fieldset>
          <fieldset class="form-group">
            <label>Phone Number</label>
            <input type="text" class="form-control" 
            v-model="phonenumber" />
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
    name: "NewEmployee",
    data() {
      return {
        firstName: "",
        lastName: "",
        middleName: "",
        emailid: "",
        sapid: "",
        organization: "",
        role: "",
        reportingto: "",
        phonenumber: "",
        errors: [],
      };
    },
    methods: {
      refreshEmployeeDetails() {
        EmployeeDataService.retrieveEmployee(this.emailid).then((res) => {
          this.firstName = res.data.firstName;
          this.lastName = res.data.lastName;
          this.middleName = res.data.middleName;
          this.emailid = res.data.emailid;
          this.sapid = res.data.sapid;
          this.organization = res.data.organization;
          this.role = res.data.role;
          this.reportingto = res.data.reportingto;
          this.phonenumber = res.data.phonenumber;
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
            EmployeeDataService.createEmployee({
              firstName: this.firstName,
              lastName: this.lastName,
              middleName: this.middleName,
              emailid: this.emailid,
              sapid: this.sapid,
              organization: this.organization,
              role: this.role,
              reportingto: this.reportingto,
              phonenumber: this.phonenumber,
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