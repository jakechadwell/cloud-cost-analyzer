<template>
    <div class="body">
        <div class="welcome-banner">
            <div class="welcome-banner-txt">
                <h2>Hello {{ firstName }}!</h2>
                <p>It's good to see you again.</p>
            </div>
            <div class="welcome-banner-avatar">
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
            </div>

            
        </div>
        <div class="search-progress-container">
            <div class="search">
                <font-awesome-icon :icon="['fas', 'magnifying-glass']" /><input type="text">
            </div>
            <div class="progress-section">
                <div class="attended">
                    <h1>11</h1>
                    <h4> Trainings Attended</h4>
                </div>
                <div class="toDo">
                    <h1>4</h1>
                    <h4> Upcoming Trainings</h4>
                </div>
            </div>
        </div>
        
        <div class="training-container">
            <div class="training-container-txt">
                <div class="training-container-txt-header">
                    <h3>Trainings</h3>
                </div>
                <div class="training-container-txt-content">
                    <h4>All Courses</h4>
                    <h4>AWS</h4>
                    <h4>Google Cloud</h4>
                    <h4>Microsoft Azure</h4>
                </div>
            </div>

            <div class="training-container-content">
                <div class="training-body" v-for="training in trainings" v-bind:key="training.trainingid">
                    <div class="thumbnail">

                    </div>
                    <div class="training-details">
                        <h5 class="cloud">{{ training.cloudid }}</h5>
                        <p class="details">{{ detailShorten(training.trainingDetail) }}</p>
                    </div>
                    <div class="path text-center">
                        <h5>{{ training.trainingPath }}</h5>
                    </div>
                    <div class="dates">
                        <p>{{ training.trainingDates }}</p>
                    </div>
                    <div class="view-btn">
                        <button @click="trainingRedirect">View Course</button>
                    </div>
                </div>
            </div>
        </div>
        <div class="statistics">
            <div class="infographic">
                {{ this.infographics }}
            </div>
        </div>
    </div>
</template>

<script>
import AuthDataService from '../service/AuthDataService';
import EmployeeDataService from '../service/EmployeeDataService';
import TrainingDataService from '../service/TrainingDataService';
export default {
  name: 'TrainingDashboard',
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
      trainings: [],
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
      EmployeeDataService.retrieveEmployeeByEmail(this.currentUser.email).then((res)=>{
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
    refreshTrainings(){
        TrainingDataService.retrieveAllTrainings().then((res)=>{
            this.trainings = res.data
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
    },
    detailShorten(details){
        if(details.length > 20){
            return details.substring(0, 20)+'...';
        }else{
            return details;
        }
    },
    getDate(dates){
        return 
    },
    trainingRedirect(training){
        return this.$router.push(`/training/${training}`)
    }
  },
  created(){
    this.refreshUserDetails();
    this.refreshRole();
    this.refreshTrainings();
  }
};  
</script>

<style scoped>
@import url('https://fonts.cdnfonts.com/css/fox-corporation-logo');

    .body{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
        margin-top: 4em;
    }

    .body > div{
        /* background-color: red; */
        width: 40%;
        margin-left: 1em;
        margin-right: 1em;
        margin-bottom: 3em;
    }

    .welcome-banner{
        display: flex;
        align-items: center;
        justify-content: space-around;
        border-radius: 1%;
        padding: 1.5em;
        box-shadow: rgba(100, 100, 111, 0.2) 1px 7px 29px 0px;
    }

    .avatar{
        box-shadow: rgba(100, 100, 111, 0.2) 1px 7px 29px 0px;

    }

    .search{
        border-radius: 2%;
        background-color: #f5f5f7;
        width: 50%;
    }

    .search>input {
        background-color: #f5f5f7; 
        border: none;
        outline: none;
    }

    .search>svg{
        margin-left: .5em;
    }

    .progress-section{
        margin-top: 1em;
        display: flex;
        justify-content: space-between;
    }

    .progress-section>div{
        background-color: #f5f5f7;
        border-radius: 2%;
        display: flex;
        align-items: center;
        padding: 1.5em;
        padding-top: 2em;
        padding-bottom: 2em;
    }

    .progress-section>div>h1{
        margin-right: .5em;
        font-family: 'FONTS', sans-serif;
    }

    .training-container-txt-content{
        display: flex;
        justify-content: space-around;
        flex-direction: row;
        margin-bottom: 2em;
    }

    .training-container-txt-content>h4{
        font-size: 20px;
    }
    
    .training-container-txt-header{
        display: flex;
        margin-bottom: 1em;
    }

    .training-body{
        background-color: #f5f5f7;
        border-radius: 2%;
        padding: 1em;
        margin-bottom: 1em;
        display: flex;
        justify-content: space-around;
        align-items: center;
    }

    .training-body>div{
        width: 15em;
    }

    .training-details{
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .view-btn>button{

        background-color: #f5f5f7;

        border: 2px solid black;

        padding: 10px;

        /* margin-top: 20px; */

        color: black;

        border-radius: 8px;
    }

</style>