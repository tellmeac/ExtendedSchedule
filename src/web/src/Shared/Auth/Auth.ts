import axios from "axios";

export function updateAuthorization(token: string) {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`
}