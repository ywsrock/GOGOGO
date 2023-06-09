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
        <Box component="div" sx={{ width: '100%', height: '100%', overflow: 'scroll' }}>
            <HomeItem></HomeItem>
        </Box >

    )
}
