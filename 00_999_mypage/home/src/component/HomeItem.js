import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Paper from '@mui/material/Paper';
import List from '@mui/material/List';
import { Divider, ListItemButton } from '@mui/material';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import Avatar from '@mui/material/Avatar';
import { Link } from 'react-router-dom';
import { useContentsList } from '../api/api';
import ErrorPage from './ErrorPage';
import { deepOrange } from '@mui/material/colors';


export default function HomeItem() {
    const { data, isError } = useContentsList("/contents")
    if (isError) return <ErrorPage code="NetorkError"></ErrorPage>

    return (
        <React.Fragment>
            <CssBaseline />
            <Paper square sx={{ pb: '50px', }}>
                {/* <Typography variant="h5" gutterBottom component="div" sx={{ p: 2, pb: 0 }}>
                    Inbox
                </Typography> */}
                <List sx={{ mb: 2 }}>
                    {data.map(({ date, id, Contents }) => (
                        <React.Fragment key={id}>
                            <ListSubheader sx={{ bgcolor: 'background.paper' }}>
                                {date}
                            </ListSubheader>
                            {Contents.map(({ id, title, subTitle, url }) => (
                                <React.Fragment key={id}>
                                    <ListItemButton>
                                        <ListItemAvatar>
                                            {/* <Avatar alt="Profile Picture" src={url} /> */}
                                            <Avatar sx={{ bgcolor: deepOrange[500] }} alt="Profile Picture"> {id}</Avatar>
                                        </ListItemAvatar>
                                        <Link to={`Contents/${url}`} target='_blank' style={{ textDecoration: 'none', color: 'black' }}>
                                            <ListItemText primary={title} secondary={subTitle} />
                                        </Link>
                                    </ListItemButton>
                                </React.Fragment>
                            ))}
                        </React.Fragment>
                    ))}
                </List>
            </Paper>
        </React.Fragment >
    );
}
