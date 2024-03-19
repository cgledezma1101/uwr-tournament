import React from "react";
import { Button, Checkbox, FormControlLabel, Box, TextField } from "@mui/material";
import GoogleIcon from "@mui/icons-material/Google";
import FacebookIcon from "@mui/icons-material/Facebook";

const Login = (): React.ReactElement => {
    // const [email, setEmail] = useState("");
    // const [password, setPassword] = useState("");
    // const [rememberMe, setRememberMe] = useState(false);

    // const handleEmailChange = (e: React.FormEvent<HTMLInputElement>): void => {
    //     const target = e.target as HTMLInputElement;
    //     setEmail(target.value);
    // };

    // const handlePasswordChange = (e: ChangeEventHandler<HTMLInputElement>): any => {
    //     const target = e.target as HTMLInputElement;
    //     setPassword(target.value);
    // };

    // const handleRememberMeChange = (e: any) => {
    //     setRememberMe(e.target.checked);
    // };

    const handleSubmit = (e: any): void => {
        e.preventDefault();
        // Add your login logic here
    };

    return (
        <Box
            className="login-container"
            alignItems={"center"}
            display={"flex"}
            flexDirection={"column"}
            justifyContent={"center"}
            sx={{ borderRadius: 1, padding: 5 }}
        >
            <h2>Login</h2>

            <Button
                style={{ minWidth: "300px", marginTop: "15px", marginBottom: "15px" }}
                variant="outlined"
                size="large"
                startIcon={<FacebookIcon />}
            >
                Log in with Facebook
            </Button>

            <Button
                style={{ minWidth: "300px", marginTop: "15px", marginBottom: "15px" }}
                variant="outlined"
                size="large"
                startIcon={<GoogleIcon />}
            >
                Log in with Google
            </Button>

            <form
                style={{
                    width: "100%",
                    alignItems: "center",
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "center",
                    margin: "10px",
                }}
                onSubmit={handleSubmit}
            >
                <TextField
                    style={{ minWidth: "300px", marginTop: "10px", marginBottom: "10px" }}
                    type="email"
                    required
                    label="Email address"
                    variant="filled"
                ></TextField>

                <TextField
                    style={{ minWidth: "300px", marginTop: "10px", marginBottom: "10px" }}
                    type="password"
                    required
                    label="Password"
                    variant="filled"
                ></TextField>

                <FormControlLabel control={<Checkbox defaultChecked />} label="Remember Me" />

                <Button
                    style={{ minWidth: "300px", marginTop: "15px", marginBottom: "15px" }}
                    variant="contained"
                    size="large"
                >
                    Log In
                </Button>
            </form>

            <div className="forgot-password">
                <a href="#">Forgot your password?</a>
            </div>
            <div className="sign-up">
                <a href="#">Sign Up</a>
            </div>
        </Box>
    );
};

export default Login;
