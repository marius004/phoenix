import axios from "axios";
import config from "../config";
import authenticationService from "./authentication.service";

const userService = {
    getUserId,
}

function getUserId() {
    const authToken = authenticationService.getAuthToken();
    return authToken?.user?.id
}

export default userService;