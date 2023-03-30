import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class EmployeeDataService {

    retrieveAllEmployees() {
        return axios.get(`${USER_API_URL}/employees`);
    }

    retrieveEmployee(emailid) {

        return axios.get(`${USER_API_URL}/employees/${emailid}`);
    }

    deleteEmployee(emailid) {

        return axios.delete(`${USER_API_URL}/employees/${emailid}`);
    }

    updateEmployee(emailid, employee) {

        return axios.put(`${USER_API_URL}/employees/${emailid}`, employee);
    }

    createEmployee(employee) {

        return axios.post(`${USER_API_URL}/employees`, employee);
    }   
}

export default new EmployeeDataService()