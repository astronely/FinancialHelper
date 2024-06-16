import axios from "axios";
import {redirect} from "react-router-dom";

export const ProfileLoader = async () => {
    let username
    await axios.get("http://localhost:8080/user/test", {withCredentials: true})
        .then(response => {
            if (response.status === 200) {
                username = response.data.username;
            }
        })
        .catch(error => {
            if (error.response.status !== 401) {
                console.log(error.response)
            }
        });
    if (username === undefined) {
        // console.log("redirect test")
        return redirect("/")
    }

    return username
}