import axios from "axios";
import config from "config";

const authenticationAPI = {
    login,
    signup,
    logout,
}

function login(username, password) {
    return axios.post(`${config.apiUrl}/auth/login`, {
        username,
        password,
    }, config.cors)
        .then(res => res.data);
}

function signup(username, password, email) {
    return axios.post(`${config.apiUrl}/auth/signup`, {
        username,
        password,
        email,
    }, config.cors)
        .then(res => res.data);
}

function logout() {
    return axios.post(`${config.apiUrl}/auth/logout`, {}, config.cors)
        .then(res => res.data);
}

export default authenticationAPI;