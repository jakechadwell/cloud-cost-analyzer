<template>
    <div class="container">
      <form>
        <h2>Sign In</h2>
        <p v-show="error" class="text-sm danger">{{ errorMsg }}</p>
        <div>
          <label> Email </label>
          <input type="email" required v-model="email" />
  
          <label> Password </label>
          <input type="password" required v-model="password" />
        </div>
      </form>
  
      <div class="sub">
        <button class="px-3" @click="handleLogin" :disabled="loading || password.length < 5 || email.length < 6">Login<font-awesome-icon class="pl-2" :icon="['fas', 'right-long']" bounce size="lg" style="color: #ffffff;" />
        <span v-show="loading" class="spinner-border spinner-border-sm"></span>
        </button>
      </div>
    </div>
</template>
<script>
import AuthDataService from '../service/AuthDataService'

    export default {
        name: 'LoginPage',
        
        data() {
            return {
                email: '',
                password: '',
                role: '',
                loading: false,
                error: false,
                errorMsg: `An error occurred, please try again`
            }
        },
        computed: {
            loggedIn(){
                return this.$store.state.auth.status.loggedIn;
            }
        },
        created(){
            if(this.loggedIn){
                this.$router.push('/account');
            }
        },
        methods: {
            handleLogin(){
                this.loading = true;
                if(this.email && this.password){
                    AuthDataService.retrieveCredentials(this.email).then((res) =>{
                        this.$store.dispatch('auth/login', {email: this.email, password: this.password, role: res.data.role}).then(
                        (response) => {
                            if(response.data.tokenString){
                                localStorage.setItem('user', JSON.stringify(response.data))
                                this.$router.push('/account/' + JSON.parse(localStorage.getItem('user')).email);
                            }
                            
                        },
                        error => {
                            this.loading = false;
                            this.error = true;
                            this.errorMsg = error.toString();
                        }
                        )
                    })  
                }
            }
        }
    }
</script>
<style scoped>
form {
        max-width: 420px;

        margin: 30px auto;

        background: white;

        text-align: left;

        padding: 40px;

        border-radius: 10px;
    }

    label {
        color: #aaa;

        display: inline-block;

        margin: 25px 0 15px;

        font-size: 0.6em;

        text-transform: uppercase;

        letter-spacing: 1px;

        font-weight: bold;
    }

    input {
        display: block;

        padding: 10px 6px;

        width: 100%;

        box-sizing: border-box;

        border: none;

        border-bottom: 1px solid #ddd;

        color: #555;
    }

    button {
        background: #01b9ff;

        border: 0;

        padding: 10px 20px;

        margin-top: 20px;

        color: white;

        border-radius: 20px;
    }

    .sub {
        text-align: center;
    }

    h2{
        text-align: center;
        padding-bottom: 1em;
    }
</style>