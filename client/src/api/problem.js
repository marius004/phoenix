import axios from "axios";
import config from "config";

const problemAPI = {
    getByQuery,
    getById,
    getByName,
    getAll,
};

function getByQuery(query) {
    return axios.get(`${config.apiUrl}/problems?${query}`, config.cors)
        .then(res => res.data)
}

function getById(id) {
    return axios.get(`${config.apiUrl}/problems?id=${id}`, config.cors)
        .then(res => res.data)
}

function getByName(name) {
    return axios.get(`${config.apiUrl}/problems/${name}`, config.cors)
        .then(res => res.data)
}

function getAll() {
    return axios.get(`${config.apiUrl}/problems`, config.cors)
        .then(res => res.data)
}

export default problemAPI;