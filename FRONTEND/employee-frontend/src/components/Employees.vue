<template>
    <div class="container">
      <h3>All Employees</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="container">
        <table class="table">
          <thead>
            <tr> 
              <th>First Name</th>
              <th>Last Name</th>
              <th>Middle Name</th>
              <th>Email ID</th>
              <th>SAP ID</th>
              <th>Organization</th>
              <th>Role</th>
              <th>Reporting To</th>
              <th>Phone Number</th>
              <th>Update</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="employee in employees" v-bind:key="employee.emailid">
            
              <td>{{ employee.firstName }}</td>
              <td>{{ employee.lastName }}</td>
              <td>{{ employee.middleName }}</td>
              <td>{{ employee.emailid }}</td>
              <td>{{ employee.sapid }}</td>
              <td>{{ employee.organization }}</td>
              <td>{{ employee.role }}</td>
              <td>{{ employee.reportingto }}</td>
              <td>{{ employee.phonenumber }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateEmployee(employee.emailid)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteEmployee(employee.emailid)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="row">
          <button class="btn btn-success" 
          v-on:click="addEmployee()">Add</button>
        </div>
      </div>
    </div>
  </template>
  <script>
  import axios from "axios";
  import EmployeeDataService from "../service/EmployeeDataService";
  
  export default {
    name: "Employees",
    data() {
      return {
        employees: [],
        message: "",
      };
    },
    methods: {
      refreshEmployees() {
        axios.get("http://localhost:9080/employees").then((res) => {
          this.employees = res.data;
        });
      },
      addEmployee() {
        this.$router.push(`/employee/new`);
      },
      updateEmployee(emailid) {
        this.$router.push(`/employee/${emailid}`);
      },
      deleteEmployee(emailid) {
        EmployeeDataService.deleteEmployee(emailid).then(() => {
          this.refreshEmployees();
        });
      },
    },
    created() {
      this.refreshEmployees();
    },
  };
  </script>