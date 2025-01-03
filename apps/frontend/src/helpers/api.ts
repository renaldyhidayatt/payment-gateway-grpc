import axios from 'axios'

const myApi = axios.create({
    baseURL:  "http://localhost:5000/api",
    headers: {
        "Content-Type": "application/json",
        Accept: "application/json"
    },
    withCredentials: true,
})

export default myApi;