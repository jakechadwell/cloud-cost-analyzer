<template>
        <div id="notfound">
            <div class="notfound">
                <div class="notfound-404">
                    <h1>4<span>0</span>1</h1>
                </div>
                <h2>Unauthorized Access</h2>
                <button @click="redirect">Back to Homepage</button>
            </div>
	    </div>
    
</template>
  
<script>
    export default {
        name: 'Unauthorized',

        methods: {
            redirect(){
                this.$router.push("/")
            }
        }
    }
    
</script>

<style>
    * {
        -webkit-box-sizing: border-box;
                box-sizing: border-box;
    }   

    body {
        padding: 0;
        margin: 0;
    }

    #notfound {
        position: relative;
        height: 100vh;
        background: #f6f6f6;
    }

    #notfound .notfound {
        position: absolute;
        left: 50%;
        top: 45%;
        -webkit-transform: translate(-50%, -50%);
            -ms-transform: translate(-50%, -50%);
                transform: translate(-50%, -50%);
    }

    .notfound {
        max-width: 767px;
        width: 100%;
        line-height: 1.4;
        padding: 110px 40px;
        text-align: center;
        background: #fff;
        -webkit-box-shadow: 0 15px 15px -10px rgba(0, 0, 0, 0.1);
                box-shadow: 0 15px 15px -10px rgba(0, 0, 0, 0.1);
    }

    .notfound .notfound-404 {
        position: relative;
        height: 180px;
    }

    .notfound .notfound-404 h1 {
        font-family: 'Roboto', sans-serif;
        position: absolute;
        left: 50%;
        top: 50%;
        -webkit-transform: translate(-50%, -50%);
            -ms-transform: translate(-50%, -50%);
                transform: translate(-50%, -50%);
        font-size: 165px;
        font-weight: 700;
        margin: 0px;
        color: #262626;
        text-transform: uppercase;
    }

    .notfound .notfound-404 h1>span {
        color: #00b7ff;
    }

    .notfound h2 {
        font-family: 'Roboto', sans-serif;
        font-size: 22px;
        font-weight: 400;
        text-transform: uppercase;
        color: #151515;
        margin-top: 0px;
        margin-bottom: 25px;
    }

    button {
        background: #01b9ff;

        border: 0;

        padding: 10px 20px;

        margin-top: 20px;

        color: white;

        border-radius: 20px;
    }

</style>