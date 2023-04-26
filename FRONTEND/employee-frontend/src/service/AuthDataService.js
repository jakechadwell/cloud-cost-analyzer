import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class AuthDataService {

    retrieveCredentials(email){
        return axios.get(`${USER_API_URL}/credentials/${email}`);
    }

    signUp(credentials) {
        return axios.post(`${USER_API_URL}/signup`, credentials);
    }

    signIn(credentials) {
        return axios.post(`${USER_API_URL}/signin`, credentials)
    }

    logout(){
        localStorage.removeItem('user')
        sessionStorage.clear()
    }
 
}

export default new AuthDataService()