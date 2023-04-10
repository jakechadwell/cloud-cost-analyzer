<template>
    <div class="container">
      <h3>All Attributes</h3>
      <div v-if="message" class="alert alert-success">
        {{ this.message }}</div>
      <div class="container">
        <table class="table">
          <thead>
            <tr> 
              <th>Attribute ID</th>
              <th>Attribute Name</th>
              <th>Stage</th>
              <th>Attribute Type</th>
              <th>Count</th>
              <th>Course</th>
              <th>Update</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="attribute in attributes" v-bind:key="attribute.attributeid">
            
              <td>{{ attribute.attributeid }}</td>
              <td>{{ attribute.attributeName }}</td>
              <td>{{ attribute.stage }}</td>
              <td>{{ attribute.attributeType }}</td>
              <td>{{ attribute.count }}</td>
              <td>{{ attribute.course }}</td>
              <td>
                <button class="btn btn-warning" 
                v-on:click="updateAttribute(attribute.attributeid)">
                  Update
                </button>
              </td>
              <td>
                <button class="btn btn-danger" 
                v-on:click="deleteAttribute(attribute.attributeid)">
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="row">
          <button class="btn btn-success" 
          v-on:click="addAttribute()">Add</button>
        </div>
      </div>
    </div>
  </template>
  <script>
  import axios from "axios";
  import AttributeDataService from "../service/AttributeDataService";
  
  export default {
    name: "Attributes",
    data() {
      return {
        attributes: [],
        message: "",
      };
    },
    methods: {
      refreshAttributes() {
        axios.get("http://localhost:9080/attributes").then((res) => {
          this.attributes = res.data;
        });
      },
      addAttribute() {
        this.$router.push(`/attribute/new`);
      },
      updateAttribute(attributeid) {
        this.$router.push(`/attribute/${attributeid}`);
      },
      deleteAttribute(attributeid) {
        AttributeDataService.deleteAttribute(attributeid).then(() => {
          this.refreshAttributes();
        });
      },
    },
    created() {
      this.refreshAttributes();
    },
  };
  </script>