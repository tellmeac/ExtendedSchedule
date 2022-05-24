import {AxiosRequestConfig} from "axios";

export function applyAuthorization(config: AxiosRequestConfig) {
    const token = localStorage.getItem("token") || "undefined"
    const value = `Bearer: ${token}`

    if (config.headers) {
        config.headers["Authorization"] = value
    } else {
        config.headers = {"Authorization": value}
    }
    return config
}

export function storeUserJwtToken(token: string) {
    localStorage.setItem("token", token)
}