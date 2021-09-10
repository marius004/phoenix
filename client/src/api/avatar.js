import axios from "axios";
import config from "config";

const avatarAPI = {
    get,
}

function get(userId, size) {
    console.log(`${config.apiUrl}/avatar/${userId}?size=${size}`)
    return axios.get(`${config.apiUrl}/avatar/${userId}?size=${size}`, 
        config.cors).then(res => res.data)
}

export default avatarAPI;