<template>
    <div class="container form-container">
        <div class="card">
            <p v-show="error" class="text-sm danger">{{ errorMsg }}</p>
            <h2>Select an Avatar</h2>
            <div class="buttons">
                <button class="save-btn" @click="validateAndSubimt">
                    Save 
                </button>
                <button class="cancel-btn" @click="cancel">
                    Cancel
                </button>
            </div>
            <form @submit="validateAndSubimt">
            <div class="card-body">
                
                <label for="1">
                    <input type="radio" id="1" value="1" v-model="avatar" />
                    <img src="../img/001.png" alt="avatar" class="avatar rounded-circle img-fluid" >                    
                </label>
            </div>
            <div class="card-body">
                
                <label for="2">
                    <input type="radio" id="2" value="2" v-model="avatar" />   
                    <img src="../img/002.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="3">
                    <input type="radio" id="3" value="3" v-model="avatar" />
                    <img src="../img/003.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="4">
                    <input type="radio" id="4" value="4" v-model="avatar" />
                    <img src="../img/004.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="5">
                    <input type="radio" id="5" value="5" v-model="avatar" />
                    <img src="../img/005.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="6">
                   <input type="radio" id="6" value="6" v-model="avatar" />
                   <img src="../img/006.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="7">
                    <input type="radio" id="7" value="7" v-model="avatar" />
                    <img src="../img/007.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="8">
                    <input type="radio" id="8" value="8" v-model="avatar" />
                    <img src="../img/008.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="9">
                    <input type="radio" id="9" value="9" v-model="avatar" />
                    <img src="../img/009.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            <div class="card-body">
                <label for="10">
                    <input type="radio" id="10" value="10" v-model="avatar" />
                    <img src="../img/010.png" alt="avatar" class="avatar rounded-circle img-fluid" >
                </label>
            </div>
            
            </form>
            
            
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
                EmployeeDataService.retrieveEmployee(+this.employeeid).then((res) =>{
                    this.$router.push(`/account/${res.data.email}`)
                })   
            )
            console.log(`Avatar ${this.avatar} was chosen.`)
        }
    },
    cancel(){
        EmployeeDataService.retrieveEmployee(+this.employeeid).then((res) =>{
            this.$router.push(`/account/${res.data.email}`)
        })
    }
  },
};
</script>

<style scoped>
@import url('https://fonts.cdnfonts.com/css/futura-xblk-bt');

    button {
        background: #01b9ff;

        border: 0;

        padding: 15px 25px;

        margin-top: 20px;

        color: white;

        border-radius: 20px;

        font-family: "Futura XBlk BT", sans-serif;

        margin-left: 30%;
        margin-right: -5em;

    }

    .cancel-btn{
        background: #dd3e31;
    }

    .buttons{
        display: flex;
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
        /* box-shadow: 0 0.5rem 1rem rgba(0, 0, 16, 0.19), 0 0.3rem 0.3rem rgba(0, 0, 16, 0.23); */
        box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
        background-color: rgb(255, 255, 255);
        padding: 0.8rem;
        border: none;
        border-radius: 2%;
        margin-bottom: 3em;
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