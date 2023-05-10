<template>
    <div class="emp-table">
      <h3 class="mt-5">All Employees</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="">
        <table class="table">
          <thead>
            <tr> 
              <th>Employee ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Dept</th>
              <th>Cloud</th>
              <th>Training Attended</th>
              <th>Training Path</th>
              <th>Email</th>
              <th>Infographics</th>
              <th>Avatar</th>
              <th>Update</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="employee in employees" v-bind:key="employee.emailid">
            
              <td>{{ employee.employeeid }}</td>
              <td>{{ employee.firstName }}</td>
              <td>{{ employee.lastName }}</td>
              <td>{{ employee.dept }}</td>
              <td>{{ employee.cloud }}</td>
              <td>{{ employee.trainingAttended }}</td>
              <td>{{ employee.trainingPath }}</td>
              <td>{{ employee.email }}</td>
              <td>{{ employee.infographics }}</td>
              <td>{{ employee.avatar }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateEmployee(employee.employeeid)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteEmployee(employee.employeeid)">
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
      updateEmployee(employeeid) {
        this.$router.push(`/employee/${employeeid}`);
      },
      deleteEmployee(employeeid) {
        EmployeeDataService.deleteEmployee(employeeid).then(() => {
          this.refreshEmployees();
        });
      },
    },
    created() {
      this.refreshEmployees();
    },
  };
  </script>

<style scoped>
  .emp-table{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>