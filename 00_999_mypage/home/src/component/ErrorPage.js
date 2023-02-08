import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import { Link } from "react-router-dom";

const ErrorPage = ({ code }) => {
    let msg = ""
    switch (code) {
        case undefined:
        case "404":
            msg = "We got lost in the fog."
            break;
        case "NetorkError":
            msg = "O..H No Something went wrong on server side."
            break;
        default:
            msg = "We are under maintenance."
            break;
    }
    return (
        <div className='Component'>
            <Alert Alert severity="error" >
                <AlertTitle>{code ? code : "404"}</AlertTitle>
                <Link to="/home"> <strong>{msg}</strong></Link>
            </Alert >
        </div>
    )
}

export default ErrorPage;