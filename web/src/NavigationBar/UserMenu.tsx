import React from "react";
import {Dropdown, Image} from "react-bootstrap";
import "./UserMenu.css"
import {UserData} from "../Shared/Models/Auth";

type Props = {
    data: UserData
    renderLogoutButton: () => JSX.Element
}

/**
 * User menu
 * @param data is an user info
 * @param renderLogoutButton is an render factory for logout button
 * @constructor
 */
export const UserMenu: React.FC<Props> = ({data, renderLogoutButton}) => {
    return <Dropdown className={"user-info"} align="end">
        <Dropdown.Toggle variant="outline" id="dropdown-autoclose-outside">
            <Image className={"user-avatar"} fluid={true} roundedCircle={true} src={data.avatar}/>
        </Dropdown.Toggle>

        <Dropdown.Menu>
            <Dropdown.Item>
                {renderLogoutButton()}
            </Dropdown.Item>
        </Dropdown.Menu>
    </Dropdown>
}