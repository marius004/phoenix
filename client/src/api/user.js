import axios from "axios";
import config from "../config";

const userAPI = {
    getByUsername,
    getGravatar,
}

function getByUsername(username) {
    return axios.get(`${config.apiUrl}/users/${username}`, config.cors)
        .then(res => res.data);
}

function getGravatar(id) {
    return axios.get(`${config.apiUrl}/users/${id}/gravatar`, config.cors)
        .then(res => res.data)
}

export default userAPI;