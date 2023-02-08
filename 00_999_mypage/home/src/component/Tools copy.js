import React from 'react'
import Card from '@mui/material/Card';
import { Button, CardActionArea, CardActions, CardContent, CardHeader, CardMedia, Divider, Stack, Typography } from '@mui/material';
import { blue, red } from '@mui/material/colors';
import Avatar from '@mui/material/Avatar';
import { SearchBar } from './DicSearch';
import { Box } from '@mui/system';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import { experimentalStyled as styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import CampaignIcon from '@mui/icons-material/Campaign';
import IconButton from '@mui/material/IconButton';
import { Diversity1 } from '@mui/icons-material';
import ShowList from './ShowList';


export const Tools = () => {
    const Item = styled(Paper)(({ theme }) => ({
        backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
        ...theme.typography.body2,
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    }));
    return (
        <Card sx={{ width: '100vw;height:200vh' }} >
            <CardHeader></CardHeader>
            <CardContent>
                <SearchBar></SearchBar>

                <Box sx={{ mt: -4 }}>

                    <Grid container spacing={{ xs: 2, md: 3 }} columns={{ xs: 4, sm: 8, md: 12, lg: 12 }}  >
                        <Grid xs={4} sm={8} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', }} variant='outlined'>
                                <CardHeader
                                    component='h1'
                                    disableTypography='true'
                                    title="what"
                                    avatar={
                                        <Avatar sx={{ bgcolor: 'rgba(0, 0, 0, 0)' }} aria-label="recipe" >
                                            <IconButton aria-label="fingerprint" color="success">
                                                <CampaignIcon />
                                            </IconButton>
                                        </Avatar>
                                    }
                                    sx={{ mb: -2 }}
                                />
                            </Card>

                        </Grid>

                        <Grid xs={4} md={1}>
                            <audio controls="" preload="none" style={{ height: "50px" }}>
                                <source src="" type="audio/mp3" media="all"></source>
                            </audio>

                            {/* <Button onclick="document.getElementById('demo2').play()">&#9205;</Button> */}
                            {/*<Button onclick="document.getElementById('demo2').pause()">&#9208;</Button>
                            <Button onclick="document.getElementById('demo2').muted= true">&#128263;</Button>
                            <Button onclick="document.getElementById('demo2').muted= false">&#128266;</Button>
                            <Button onclick="document.getElementById('demo2').volume+=0.1">+</Button>
                            <Button onclick="document.getElementById('demo2').volume-=0.1">-</Button> */}
                            {/* <MediaPlayer ></MediaPlayer> */}
                        </Grid>


                        <Grid xs={4} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>音節:</span> cel・e・bra  音節: cel・e・brat &nbsp;&nbsp;
                                        {/* <Divider variant="middle" /> */}
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>読み方:</span> /séləbrèɪt(米国英語), ˈsɛl.ɪ.bɹeɪt(英国英語)  読み方: /séləbrèɪt(米国英語)
                                        {/* <Divider component="div" variant="inset" />
                                    <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>意味・対訳::</span> /séləbrèɪt(米国英語), ˈsɛl.ɪ.bɹeɪt(英国英語)  読み方: /séləbrèɪt(米国英語) */}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                        <Grid xs={4} md={1}>
                            <audio controls="" preload="none" style={{ height: "50px" }}>
                                <source src="" type="audio/mp3" media="all"></source>
                            </audio>
                        </Grid>

                        <Grid xs={4} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', }} variant='outlined'>
                                <Divider></Divider>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>意味・対訳:</span> celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                        <Grid xs={4} md={1}>
                            <audio controls="" preload="none" style={{ height: "50px" }}>
                                <source src="" type="audio/mp3" media="all"></source>
                            </audio>
                        </Grid>


                        <Grid xs={4} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>SAMPLE:</span>
                                        {/* <Divider></Divider> */}
                                    </Typography>
                                    <ShowList></ShowList>
                                </CardContent>
                            </Card>
                        </Grid>
                        <Grid xs={4} md={1}>
                            <audio controls="" preload="none" style={{ height: "50px" }}>
                                <source src="" type="audio/mp3" media="all"></source>
                            </audio>
                        </Grid>

                    </Grid>
                </Box>
            </CardContent>

        </Card >
    )
}


