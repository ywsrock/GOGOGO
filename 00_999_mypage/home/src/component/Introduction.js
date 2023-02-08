import { useTheme } from '@mui/material/styles';
import { useMediaQuery } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { ActionCard, LeftContents, SelfIntroduction } from './ContentsLeft';
import { RightContents } from "./ContentsRight";
import { useEffect, useState } from 'react';
import { useAbout } from '../api/api';
import ErrorPage from './ErrorPage';

export const Introduction = () => {
    // media Query
    const theme = useTheme();
    const down_md = useMediaQuery(theme.breakpoints.down('md'))
    // window height
    const [dsHeight, setDisHeight] = useState(window.innerHeight);
    // resize monitor
    useEffect(() => {
        window.addEventListener('resize', () => {
            setDisHeight(window.innerHeight)
        });
        return () => window.removeEventListener('resize', () => { })
    })

    // md　Introduction left
    const base_cmps = [];
    const ds_style = (() => {
        if (down_md) {
            return {
                Name: "h4",
                SubTitle: "h6",
                PhoneNo: "h6",
                Email: "h6",
                Ads: "h6",
                isdmd: down_md
            }
        }
        return false;
    })();

    // About Data Fetch
    const { data, isError } = useAbout("/about")
    if (isError) return <ErrorPage code="NetorkError"></ErrorPage>

    // base display left | right
    base_cmps.push(
        <LeftContents ds={ds_style} >
            {
                creatCardActionNavMeanu(data)
            }
        </LeftContents>, <RightContents ds={ds_style} />)

    return (
        <Grid container spacing={0} columns={{ xs: 2, md: 6, lg: 12 }} >
            <Grid xs={2} md={3} lg={6} sx={{ height: `${down_md ? '60vh' : dsHeight}px`, overflow: 'auto' }} component='div'>
                {
                    !down_md && base_cmps[0]
                }
                {
                    !down_md && Informations(data)
                }
            </Grid>
            <Grid xs={2} md={3} lg={6} sx={{ height: `${dsHeight}px`, overflow: 'auto' }} component='div'>
                {
                    down_md && base_cmps[1]
                }
                {
                    down_md ? base_cmps[0] : base_cmps[1]
                }
                {
                    down_md && Informations(data)
                }
            </Grid>
        </Grid >
    )
}

function Informations(data) {
    return (data.results.map((info, index) => {
        return (<SelfIntroduction ds={{ Name: 'body1' }} title={info.title} id={info.id} content={info.content} key={index}>
            {creatCardActionNavMeanu(data)}
        </SelfIntroduction>
        );
    }));
}

function creatCardActionNavMeanu(data) {
    return data.results.filter((info) => info.id).map((item, index) => {
        return (
            <ActionCard id={item.id} title={item.title} key={index} />
        );
    });
}