import axios from "axios";
import config from "../config";

const userAPI = {
    getByUsername,
}

function getByUsername(username) {
    return axios.get(`${config.apiUrl}/users/${username}`, config.cors)
        .then(res => res.data);
}

export default userAPI;