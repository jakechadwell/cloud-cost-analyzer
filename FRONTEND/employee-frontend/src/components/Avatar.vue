<template>
    <div class="container form-container">
        <div class="card">
            <p v-show="error" class="text-sm danger">{{ errorMsg }}</p>
            <h2>Select an Avatar</h2>
            <form @submit="validateAndSubimt">
            <div class="card-body">
                
                <label for="001">
                    <input type="radio" id="001" value="1" checked v-model="avatar" />
                    <img src="../img/001.png" alt="avatar" class="avatar rounded-circle img-fluid" >                    
                </label>
            </div>
            <div class="card-body">
                
                <label for="002">
                    <input type="radio" id="002" value="2" v-model="avatar" />   
                    <img src="../img/002.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="003">
                    <input type="radio" id="003" value="3" v-model="avatar" />
                    <img src="../img/003.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="004">
                    <input type="radio" id="004" value="4" v-model="avatar" />
                    <img src="../img/004.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="005">
                    <input type="radio" id="005" value="5" v-model="avatar" />
                    <img src="../img/005.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="006">
                   <input type="radio" id="006" value="6" v-model="avatar" />
                   <img src="../img/006.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="007">
                    <input type="radio" id="007" value="7" v-model="avatar" />
                    <img src="../img/007.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="008">
                    <input type="radio" id="008" value="8" v-model="avatar" />
                    <img src="../img/008.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="009">
                    <input type="radio" id="009" value="9" v-model="avatar" />
                    <img src="../img/009.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="010">
                    <input type="radio" id="010" value="10" v-model="avatar" />
                    <img src="../img/010.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            
            </form>
            <button class="save-btn" @click="validateAndSubimt">
               Save 
            </button>
        </div>
    </div>

</template>

<script>
import EmployeeDataService from '../service/EmployeeDataService';

export default {
  name: 'Avatar',
  data(){
    return{
      avatar: "",
      error: false,
      errorMsg: "Sorry an error has occured",
    }
  },
  computed: {
    currentUser() {
        return this.$store.state.auth.user;
    },
    employeeid() {
        return this.$route.params.employeeid;
    },
  },
  methods: {
    validateAndSubimt(e){
        e.preventDefault();
        this.error = false;
        if(this.avatar==null){
           this.errorMsg = "There was a problem, please try again.";
        }
        if(this.error==false){
            EmployeeDataService.updateAvatar(this.employeeid, {
                avatar: this.avatar
            }).then(
                EmployeeDataService.retrieveEmployee(this.employeeid).then((res) =>{
                    this.$router.push(`/account/${res.data.email}`)
                })   
            )
            console.log(`Avatar ${this.avatar} was chosen.`)
        }
    }
  },
  created(){

  }
};
</script>

<style scoped>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

    button {
        background: #01b9ff;

        border: 0;

        padding: 10px 20px;

        margin-top: 20px;

        color: white;

        border-radius: 20px;


        margin-left: auto;
        margin-right: auto;

        margin-bottom: 3em;
    }

    h2 {
        text-align: center;
        margin-top: 1em;
        font-family: "Futura XBlk BT", sans-serif;
    }

    .card-body{
        width: 128px;
        padding: 0 !important;
    }

    div .container{
        margin-top: 5% !important;
    }

    form{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        align-items: center;
        justify-content: center;
    }

    form .card-body{
        margin: 3.75em;
        flex: initial;
    }

    .card{
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        margin: auto;
        -webkit-box-shadow: 0 0.5rem 1rem rgba(0, 0, 16, 0.19), 0 0.3rem 0.3rem rgba(0, 0, 16, 0.23);
        box-shadow: 0 0.5rem 1rem rgba(0, 0, 16, 0.19), 0 0.3rem 0.3rem rgba(0, 0, 16, 0.23);
        background-color: rgb(255, 255, 255);
        padding: 0.8rem;
        border: none;
        border-radius: 2%;
    }


    input[type="radio"]{
        position: absolute;
        opacity: 0;
        width: 0;
        height: 0;
    }

    [type="radio"]{
        cursor: pointer;
    }

    [type="radio"]:checked + img{
        transition: 100ms;
        transform: scale(1.2);
        transition-timing-function: ease-in-out;
        border: #01b9ff solid medium;
        outline: #01b9ff solid 2px;
    }

</style>