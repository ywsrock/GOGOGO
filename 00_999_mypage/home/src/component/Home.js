import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import { experimentalStyled as styled } from '@mui/material/styles';
import React from 'react';
import { useHome } from '../api/api';
import './Component.css';
import ErrorPage from './ErrorPage';
import HomeItem from './HomeItem';

const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));


export const Home = () => {
    const { data, isError } = useHome("/home")
    if (isError) return <ErrorPage code="NetorkError"></ErrorPage>
    return (
        // <div>
        //     <Box>
        //         <Grid container spacing={{ xs: 0.5, md: 1 }} columns={{ xs: 1, sm: 4, md: 3, lg: 12 }}>
        //             {data.results.map((item, index) => (
        //                 <Grid xs={1} sm={2} md={1} lg={4} key={index} sx={{ ml: '', py: '1em' }}>
        //                     {/* <div style={{ flexGrow: 1, justifyContent: 'center', alignItems: 'center', display: 'flex' }}> */}
        //                     <div style={{ flexGrow: 1, justifyContent: 'center', alignItems: 'start', display: 'flex' }} >
        //                         <HomeCard item={item} className='Component-Home' />
        //                     </div>
        //                 </Grid>
        //             ))}
        //         </Grid>
        //     </Box >
        // </div >
        <Box component="div" sx={{ width: '100vw', height: '100vh', overflow: 'scroll' }}>
            <HomeItem></HomeItem>
        </Box >
        // <Box>
        //     <HomeItem></HomeItem>
        // </Box >
    )
}
