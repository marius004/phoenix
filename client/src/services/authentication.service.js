import axios from "axios";
import config from "../config";
import Cookies from "js-cookie";
import base64 from "base-64";

const authenticationService = {
    login,
    signup,
    logout,
    isLoggedIn,
    getAuthToken,
}

function login(username, password) {
    return axios.post(`${config.apiUrl}/auth/login`, {
        username,
        password,
    }, config.cors);
}

function getAuthToken() {
    const data = Cookies.get("auth-token");
    try {
        const decoded = base64.decode(data);
        return JSON.parse(decoded); 
    }catch(e) {
        return null;
    }
}

function signup(username, password, email) {
    return axios.post(`${config.apiUrl}/auth/signup`, {
        username,
        password,
        email,
    }, config.cors);
}

function logout() {
    return axios.post(`${config.apiUrl}/auth/logout`, {}, config.cors);
}

function isLoggedIn() {
    const authToken = getAuthToken();
    return authToken != null && authToken.token != null;
}

export default authenticationService;