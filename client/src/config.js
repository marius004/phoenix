const config = {
    apiUrl: "http://localhost:8080/api",
    backendDomain: "http://localhost:8080",
    cors: {
        "withCredentials": true,
        "Access-Control-Allow-Origin": "http://6cd7-178-138-98-245.ngrok.io",
        "Access-Control-Allow-Credentials": true,
    },
    submissionsLimit: 25, // used on the submissions page for the maximum number of submissions to fetch
};

export default config;