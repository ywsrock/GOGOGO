import {
    Avatar,
    Box,
    Button,
    Card,
    CardActions,
    CardContent,
    Divider,
    Typography
  } from '@mui/material';
  import { React,useState,useEffect } from "react";
  import { Auth } from 'aws-amplify';
  import Stack from '@mui/material/Stack';
  import { deepOrange } from '@mui/material/colors';
 
  export const AccountProfile = () => {
    const [user, setUser] = useState();
    
    async function getUserInfo() {
        const AccountProfile = await Auth.currentSession()
        setUser(AccountProfile.idToken.payload);
        console.dir(AccountProfile.idToken.payload)
    }
    useEffect(() => {
        getUserInfo()
       
    }, [])

    const userData = {
        avatar: '/assets/avatars/avatar-anika-visser.png',
        name: user ? user['cognito:username'] : "No Data",
        email: user ? user['email'] : "No Data"
      };

  return(
    <Card>
      <CardContent>
        <Box
          sx={{
            alignItems: 'center',
            display: 'flex',
            flexDirection: 'column'
          }}
        >
          <Stack direction="row" spacing={2}>
            <Avatar
              sx={{ bgcolor: deepOrange[500] }}
              alt="Remy Sharp"
              src="/broken-image.jpg"
            >
              SA
            </Avatar>
          </Stack>

          <Typography
            gutterBottom
            variant="h5"
          >
            {userData.name}
          </Typography>
          <Typography
            color="text.secondary"
            variant="body2"
          >
            {userData.email}
          </Typography>
        </Box>
      </CardContent>
      <Divider />
      <CardActions>
        <Button
          fullWidth
          variant="text"
        >
          画像取込み
        </Button>
      </CardActions>
    </Card>
  )};