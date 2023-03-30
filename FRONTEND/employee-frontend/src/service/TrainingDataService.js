import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class TrainingDataService {

    retrieveAllTrainings() {
        return axios.get(`${USER_API_URL}/trainings`);
    }

    retrieveTraining(courseCode) {

        return axios.get(`${USER_API_URL}/trainings/${courseCode}`);
    }

    deleteTraining(courseCode) {

        return axios.delete(`${USER_API_URL}/trainings/${courseCode}`);
    }

    updateTraining(courseCode, training) {

        return axios.put(`${USER_API_URL}/trainings/${courseCode}`, training);
    }

    createTraining(training) {

        return axios.post(`${USER_API_URL}/trainings`, training);
    }   
}

export default new TrainingDataService()