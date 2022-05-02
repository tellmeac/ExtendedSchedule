import {GoogleLoginResponse} from "react-google-login";

export interface UserAuthContent {
    id_token: string
    avatar: string
    name: string
}

export function getUserAuthContentFromResponse(r: GoogleLoginResponse): UserAuthContent {
    const basic = r.getBasicProfile()
    return {
        id_token: r.tokenId,
        avatar: basic.getImageUrl(),
        name: basic.getName()
    }
}