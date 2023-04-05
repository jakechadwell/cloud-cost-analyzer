import axios from 'axios'

const USER_API_URL = 'http://localhost:9080'

class TrainingDataService {

    retrieveAllTrainings() {
        return axios.get(`${USER_API_URL}/trainings`);
    }

    retrieveTraining(trainingid) {

        return axios.get(`${USER_API_URL}/trainings/${trainingid}`);
    }

    deleteTraining(trainingid) {

        return axios.delete(`${USER_API_URL}/trainings/${trainingid}`);
    }

    updateTraining(trainingid, training) {

        return axios.put(`${USER_API_URL}/trainings/${trainingid}`, training);
    }

    createTraining(training) {

        return axios.post(`${USER_API_URL}/trainings`, training);
    }   
}

export default new TrainingDataService()