import React from 'react'
import Card from '@mui/material/Card';
import { Button, CardActions, CardContent, CardHeader, Divider, Typography } from '@mui/material';
import Avatar from '@mui/material/Avatar';
import { SearchBar } from './DicSearch';
import { Box } from '@mui/system';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import CampaignIcon from '@mui/icons-material/Campaign';
import IconButton from '@mui/material/IconButton';
import ShowList from './ShowList';
import ShowHistory from './ShowHistory';
import Statistic from './Statistic';


export const Tools = () => {

    return (
        <Card sx={{ width: '100vw;height:200vh' }} >
            <audio controls="" preload="none" style={{ height: "50px" }}>
                <source src="" type="audio/mp3" media="all"></source>
            </audio>
            <CardContent>
                <SearchBar></SearchBar>
                <Box >
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
                                    sx={{ mt: -1 }}
                                />
                            </Card>
                        </Grid>

                        <Grid xs={4} sm={8} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', mt: -7 }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>音節:</span> cel・e・bra  音節: cel・e・brat &nbsp;&nbsp;
                                        {/* <Divider variant="middle" /> */}
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>読方:</span> /séləbrèɪt(米国英語), ˈsɛl.ɪ.bɹeɪt(英国英語)  読み方: /séləbrèɪt(米国英語)
                                        {/* <Divider component="div" variant="inset" />
                                    <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>意味・対訳::</span> /séləbrèɪt(米国英語), ˈsɛl.ɪ.bɹeɪt(英国英語)  読み方: /séləbrèɪt(米国英語) */}
                                    </Typography>
                                    <Divider></Divider>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid xs={4} sm={8} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', mt: -5 }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>対訳:</span> celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                        celebrateの意味・対訳は、挙行する、(式を挙げて)祝う、ほめたたえる、ほめて世に知らせる、などです。
                                    </Typography>
                                    <Divider></Divider>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid xs={4} sm={8} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', mt: -5 }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>例文:</span>
                                        <Divider></Divider>
                                    </Typography>
                                    <ShowList></ShowList>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid xs={4} sm={8} md={9} mdOffset={2} lgOffset={2} >
                            <Card sx={{ maxWidth: '100%', borderStyle: 'none', mt: -5 }} variant='outlined'>
                                <CardContent>
                                    <Typography variant='body1' >
                                        <span style={{ backgroundColor: "#ffc107", color: "#212529", fontSize: "1.2rem" }}>履歴:
                                        </span>
                                        <Divider></Divider>
                                    </Typography>
                                    <ShowHistory></ShowHistory>
                                </CardContent>
                            </Card>
                        </Grid>
                    </Grid>
                </Box>
            </CardContent>
        </Card >
    )
}


