// export const UserInfo = () => {
//     const [user, setUser] = useState();
//     async function getUserInfo() {
//         const UserInfo = await Auth.currentSession()
//         setUser(UserInfo.idToken.payload);
//     }
//     return (
//         <>
//             <header className="App-header">
//                 <div>UserName : {user ? user.name : "No Data"}</div>
//                 <div>Cognito UserName : {user ? user['cognito:username'] : "No Data"}</div>
//                 <div>Group : {user ? user['cognito:groups'] : "No Data"}</div>
//                 <div><button id='button' onClick={() => getUserInfo()}>Show User Info</button></div>
//             </header>
//         </>
//     )
// }
// import Head from 'next/head';
import { Box, Container, Stack, Typography, Unstable_Grid2 as Grid } from '@mui/material';
import { I18n } from 'aws-amplify';
import { React } from "react";
import { vocabularies } from '../menu/const';
import { AccountProfile } from './account-profile';
import { AccountProfileDetails } from './account-profile-details';
I18n.putVocabularies(vocabularies);
I18n.setLanguage('ja');

export const UserInfo = () => {
    return (
        <Box
            component="main"
            sx={{
                flexGrow: 1,
                py: 8
            }}
        >
            <Container maxWidth="lg">
                <Stack spacing={3}>
                    <div>
                        <Typography variant="h4">
                            アカウント
                        </Typography>
                    </div>
                    <div>
                        <Grid
                            container
                            spacing={3}
                        >
                            <Grid
                                xs={12}
                                md={6}
                                lg={4}
                            >
                                <AccountProfile />
                            </Grid>
                            <Grid
                                xs={12}
                                md={6}
                                lg={8}
                            >
                                <AccountProfileDetails />
                            </Grid>
                        </Grid>
                    </div>
                </Stack>
            </Container>
        </Box>
)};
