import axios from "axios";
import config from "config";

const problemService = {
    getByName,
    getAll,
    get,
};

function get(queryArgs) {
    return axios.get(`${config.apiUrl}/problems?${queryArgs}`);
}

function getByName(name) {
    return axios.get(`${config.apiUrl}/problems/${name}`);
}

function getAll() {
    return axios.get(`${config.apiUrl}/problems`);
}

export default problemService;