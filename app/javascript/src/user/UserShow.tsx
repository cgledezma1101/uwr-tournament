import React from "react";
import { Link } from "react-router-dom";

const UserShow = (): JSX.Element => (
    <div>
        This is the profile of the user want to go to <Link to="/v2/tournament">tournament</Link>?
    </div>
);

export default UserShow;
