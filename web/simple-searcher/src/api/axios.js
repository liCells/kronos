import axios from 'axios'

const req = axios.create({
    baseURL: 'http://localhost:18000',
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: false,
    timeout: 3000,
})

req.interceptors.response.use(
    response => {
        return response.data
    }
)

export default req
