import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {RootState} from "./Root";

export interface UserState {
    isAuthorized: boolean
    credentials: {
        token: string,
        username: string,
        avatarUrl: string,
    }
}

const initialUserState: UserState = {
    isAuthorized: false,
    credentials: {
        token: "",
        username: "",
        avatarUrl: "",
    }

}

export const userSlice = createSlice({
    name: "user",
    initialState: initialUserState,
    reducers: {
        setCredentials: (state, action: PayloadAction<{token: string, avatarUrl: string, username: string}>) => {
            state.credentials = action.payload
            state.isAuthorized = true
        },
    }
})

export const selectSignedIn = (state: RootState) => state.user.isAuthorized
export const selectUserInfo = (state: RootState) => state.user.credentials

export const { setCredentials, } = userSlice.actions