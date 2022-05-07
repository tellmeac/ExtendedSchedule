import React from "react";
import {Dropdown, Image} from "react-bootstrap";
import "./UserMenu.css"
import {UserAuthContent} from "../Shared/Models/Auth";

type Props = {
    data: UserAuthContent
    renderLogoutButton: () => JSX.Element
}

export const UserMenu: React.FC<Props> = ({data, renderLogoutButton}) => {
    return <div className={"user-info"}>
        <Dropdown align="end">
            <Dropdown.Toggle id="dropdown-autoclose-outside">
                <Image className={"user-avatar"} fluid={true} roundedCircle={true} src={data.avatar}/>
            </Dropdown.Toggle>

            <Dropdown.Menu>
                <Dropdown.Item>
                    {renderLogoutButton()}
                </Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    </div>
}