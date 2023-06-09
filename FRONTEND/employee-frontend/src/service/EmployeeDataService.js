import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class EmployeeDataService {

    retrieveAllEmployees() {
        return axios.get(`${USER_API_URL}/employees`);
    }

    retrieveEmployee(employeeid) {

        return axios.get(`${USER_API_URL}/employees/${employeeid}`);
    }

    retrieveEmployeeByEmail(email) {
        return axios.get(`${USER_API_URL}/employee/${email}`);
    }

    deleteEmployee(employeeid) {

        return axios.delete(`${USER_API_URL}/employees/${employeeid}`);
    }

    updateEmployee(employeeid, employee) {

        return axios.put(`${USER_API_URL}/employees/${employeeid}`, employee);
    }

    createEmployee(employee) {

        return axios.post(`${USER_API_URL}/employees`, employee);
    }   

    updateAvatar(employeeid, avatar){
        return axios.put(`${USER_API_URL}/edit/avatar/${employeeid}`, avatar);
    }
}

export default new EmployeeDataService()