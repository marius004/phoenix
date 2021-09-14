import axios from "axios";
import config from "config";

const problemAPI = {
    getByQuery,
    getById,
    getByName,
    getAll,
    create,
    update,
    delete: deleteProblem,
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

function create(problem) {
    return axios.post(`${config.apiUrl}/problems`, problem, 
        config.cors).then(res => res.data);
}

function update(problemName, data) {
    console.log(problemName);
    return axios.put(`${config.apiUrl}/problems/${problemName}`, data, 
        config.cors).then(res => res.data);
}

function deleteProblem(problemName) {
    return axios.delete(`${config.apiUrl}/problems/${problemName}`, 
        config.cors).then(res => res.data);
}

export default problemAPI;