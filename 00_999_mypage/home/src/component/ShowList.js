import { ListItemButton } from '@mui/material';
import CssBaseline from '@mui/material/CssBaseline';
import List from '@mui/material/List';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemText from '@mui/material/ListItemText';
import Paper from '@mui/material/Paper';
import { split, trim } from 'lodash';
import * as React from 'react';

export default function ShowList(props) {
    const messages = split(props.WordData.w.InfosJp, "\n").filter(v => trim(v) !== "")
    return (
        <React.Fragment>
            <CssBaseline />
            <Paper square elevation={0} sx={{
                pb: '50px', height: '200px', overflowY: 'scroll', borderColor: '#4caf50'
            }} variant="outlined" component='div'>
                <List sx={{ mb: 2, borderStyle: 'none' }} component='div'>
                    {/* {messages.map(({ id, primary, secondary, person }) => ( */}
                    {messages.map((v, id) => (
                        <React.Fragment key={id}>
                            <ListItemButton dense={true} divider>
                                <ListItemAvatar>
                                    {/* <Avatar alt="Profile Picture" src={person} /> */}
                                    {id + 1}
                                </ListItemAvatar>
                                <ListItemText primary={v} />
                            </ListItemButton>
                        </React.Fragment>
                    ))}
                </List>
            </Paper>
        </React.Fragment >
    );
}
