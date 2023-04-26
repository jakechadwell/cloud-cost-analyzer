<template>

  <section>
    <div class="con">
        <div class="main-card card top">
          <div class="card-body av text-center ">
            <div class="avatar-container">
              <img v-if="avatar==1" src="../img/001.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==2" src="../img/002.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==3" src="../img/003.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==4" src="../img/004.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==5" src="../img/005.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==6" src="../img/006.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==7" src="../img/007.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==8" src="../img/008.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==9" src="../img/009.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
                <img v-if="avatar==10" src="../img/010.png" alt="avatar"
                class="avatar rounded-circle img-fluid" style="width: 150px;">
              

              <div class="content" @click.prevent="avatarRedirect(employeeid)">
                <div class="text">Edit Avatar</div>
              </div>
            </div>
            <div class="edit-btn-div rounded-circle img-fluid">
              <img @click.prevent="editRedirect(employeeid)" title="Edit Employee Details" src="../img/edit.png" alt="edit"
              class="rounded-circle img-fluid edit-btn">
            </div>
            

            <h5 class="my-3">{{ firstName }} {{ lastName }}</h5>
            <p class="text-muted mb-1">Dept: {{ dept }}</p>
            <p class="text-muted mb-4">Cloud: {{ cloud }}</p>
          </div>
        </div>
    </div>
      
    <div class="con">
        <div class="main-card card top">
          <div class="card-body text-center">
            <h2>Trainings Completed :</h2>
              <div class="container card-container my-4">
                <div class="training-card card">
                  <div class="card-body text-center">
                      <h5>Cloud: </h5>
                      <p class="text-muted"></p>
                      <p>Description: </p>
                      <p>Dates: </p>
                  </div>
              </div>

              <div class="training-card card">
                  <div class="card-body text-center">
                      <h5>Cloud: </h5>
                      <p class="text-muted"></p>
                      <p>Description: </p>
                      <p>Dates: </p>
                  </div>
              </div>

              <div class="training-card card">
                  <div class="card-body text-center">
                      <h5>Cloud: </h5>
                      <p class="text-muted"></p>
                      <p>Description: </p>
                      <p>Dates: </p>
                  </div>
              </div>
            </div>
          </div>
        </div>
    </div>
  </section>
</template>

<script>
import AuthDataService from '../service/AuthDataService';
import EmployeeDataService from '../service/EmployeeDataService';
export default {
  name: 'Account',
  data(){
    return{
      employeeid: "",
      firstName: "",
      lastName: "",
      dept: "",
      cloud: "",
      trainingAttended: "",
      trainingPath: "",
      email: "",
      infographics: "",
      avatar: "",
      role: "",
      errors: [],
    }
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
    accountEmail() {
      return this.$route.params.email;
    },
  },
  methods: {
    refreshUserDetails(){
      EmployeeDataService.retrieveEmployeeByEmail(this.accountEmail).then((res)=>{
        this.employeeid = res.data.employeeid;
        this.firstName = res.data.firstName;
        this.lastName = res.data.lastName;
        this.dept = res.data.dept;
        this.cloud = res.data.cloud;
        this.trainingAttended = res.data.trainingAttended;
        this.trainingPath = res.data.trainingPath;
        this.email = res.data.email;
        this.infographics = res.data.infographics;
        this.avatar = res.data.avatar;
      })
    },
    refreshRole(){
      AuthDataService.retrieveCredentials(this.accountEmail).then((res)=>{
        this.role = res.data.role;
      })
    },
    editRedirect(employeeid){
      return this.$router.push(`/employee/${employeeid}`)
    },
    avatarRedirect(employeeid){
      return this.$router.push(`/edit/avatar/${employeeid}`)
    }
  },
  created(){
    this.refreshUserDetails();
    this.refreshRole();
    
  }
  // mounted() {
  //   if (!this.currentUser) {
  //     this.$router.push('/signin');
  //   }
  // }
};  
</script>

<style scoped>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

  .con{
    width: 100%;
    padding-right: 5em;
    padding-left: 5em;
    margin-left: auto;
    margin-right: auto;

  }

  .card{
    border: none;
    margin-top: 2em;
    align-content: center !important;
  }

  .main-card{
    box-shadow: rgba(100, 100, 111, 0.2) 1px 7px 29px 0px;
  }

  .training-card{
    border: rgba(100, 100, 111, 0.2) solid thin;
    width: 30%;
  }

  .card-container{
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    align-content: center;
    justify-content: space-around;
    align-items: center;
  }

  .edit-btn {
    width: 35px;
    position: inherit;
    /* transform: translate(50px, 115px); */
    cursor: pointer;
  }

  .edit-btn-div{
    transform: translate(30px, 119px);
    width: 35px;
    height: 35px;
    position: absolute;
  }

  .edit-btn-div :hover{
    transform: scale(1.25);
    transition: all 100ms ease-in-out;
    box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
  }

  .avatar{
    position: inherit;
  }

  .av{
    display: flex;
    flex-direction: column;
    align-content: center;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    
  }
  
  .avatar-container{
    max-width: 150px;
    max-height: 150px;
    border-radius: 50%;
  }

  .avatar-container:hover{
    opacity: 80%;
  }

  .content{
    opacity: 0;
    position: absolute;
    -webkit-transition: all 100ms ease-out;
    -webkit-transition: all 100ms ease-in-out;
    font-size: smaller;
    transform: translate(0, -9em);
    height: 150px;
    width: 150px;
    border-radius: 50%;
    display: flex;
    flex-direction: row;
    align-content: center;
    align-items: center;
    justify-content: center;
  }

  .text{
    font-family: "Futura XBlk BT", sans-serif;
  }

  .content:hover{
    opacity: 1;
    cursor: pointer;
  }

</style>