<template>
    <div>
      <h3>Attribute</h3>
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
            <label>Attribute Name</label>
            <input type="text" class="form-control"
             v-model ="attributeName" />
          </fieldset>
          <fieldset class="form-group">
            <label>Stage</label>
            <input type="text" class="form-control"
             v-model="stage" />
          </fieldset>
          <fieldset class="form-group">
            <label>Attribute Type</label>
            <input type="text" class="form-control"
             v-model="attributeType" />
          </fieldset>
          <fieldset class="form-group">
            <label>Count</label>
            <input type="text" class="form-control"
             v-model="count" />
          </fieldset>
          <fieldset class="form-group">
            <label>Course</label>
            <input type="text" class="form-control"
             v-model="course" />
          </fieldset>
          <button class="btn btn-success" 
          type="submit">Save</button>
        </form>
      </div>
    </div>
  </template>
  <script>
  import AttributeDataService from "../service/AttributeDataService";
  
  export default {
    name: "Attribute",
    data() {
      return {
        attributeName: "",
        stage: "",
        attributeType: "",
        count: "",
        course: "",
        errors: [],
      };
    },
    computed: {
      attributeid() {
        return this.$route.params.attributeid;
      },
    },
    methods: {
      refreshAttributeDetails() {
        AttributeDataService.retrieveAttribute(this.attributeid).then((res) => {
          this.attributeName = res.data.attributeName;
          this.stage = res.data.stage;
          this.attributeType = res.data.attributeType;
          this.count = res.data.count;
          this.course = res.data.course;
        });
      },
      validateAndSubmit(e) {
        e.preventDefault();
        this.errors = [];
        if (!this.attributeName) {
          this.errors.push("Enter valid values");
        }
        if (this.errors.length === 0) {
            AttributeDataService.updateAttribute(this.attributeid, {
                attributeName: this.attributeName,
                stage: this.stage,
                attributeType: this.attributeType,
                count: this.count,
                course: this.course
            }).then(() => {
              this.$router.push("/attributes");
            });
          
        }
      },
    },
    created() {
      this.refreshAttributeDetails();
    }
  };
  </script>