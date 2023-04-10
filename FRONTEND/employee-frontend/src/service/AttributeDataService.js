import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class AttributeDataService {

    retrieveAllAttributes() {
        return axios.get(`${USER_API_URL}/attributes`);
    }

    retrieveAttribute(attributeid) {

        return axios.get(`${USER_API_URL}/attributes/${attributeid}`);
    }

    deleteAttribute(attributeid) {

        return axios.delete(`${USER_API_URL}/attributes/${attributeid}`);
    }

    updateAttribute(attributeid, attribute) {

        return axios.put(`${USER_API_URL}/attributes/${attributeid}`, attribute);
    }

    createAttribute(attribute) {

        return axios.post(`${USER_API_URL}/attributes`, attribute);
    }   
}

export default new AttributeDataService()