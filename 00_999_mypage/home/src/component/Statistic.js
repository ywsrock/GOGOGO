import { Typography } from '@mui/material';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Drawer from '@mui/material/Drawer';
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import * as React from 'react';
import ModalTable from './ModalTable';
import { useDictionalSts } from '../api/api';

export default function Statistic() {
    const { StsData, StsIsError } = useDictionalSts()

    const [state, setState] = React.useState({
        top: false,
    });

    const toggleDrawer = (anchor, open) => (event) => {
        if (event.type === 'keydown' && (event.key === 'Tab' || event.key === 'Shift')) {
            return;
        }

        setState({ ...state, [anchor]: open });
    };

    const list = (anchor) => (
        <Box
            sx={{ width: anchor === 'top' || anchor === 'bottom' ? 'auto' : 250 }}
            role="presentation"
            onClick={toggleDrawer(anchor, false)}
            onKeyDown={toggleDrawer(anchor, false)}
        >
            <Grid container spacing={{ xs: 1, md: 1 }} columns={{ xs: 4, sm: 8, md: 12, lg: 12 }}  >
                <Grid xs={4} sm={8} md={6} lg={6}>

                    <Box sx={{ height: '98%', overflow: 'scroll' }} maxHeight='50vh'>
                        <Typography variant='h4'>日統計：</Typography>
                        <ModalTable type="d" StsData={StsData}></ModalTable>
                    </Box>
                </Grid>
                <Grid xs={4} sm={8} md={6} lg={6}>
                    <Box sx={{ height: '98%', overflow: 'scroll' }} maxHeight='50vh'>
                        <Typography variant='h4'>月統計：</Typography>
                        <ModalTable type="m" StsData={StsData}></ModalTable>
                    </Box>
                </Grid>
            </Grid>

        </Box >
    );

    return (
        <div>
            {['top'].map((anchor) => (
                <React.Fragment key={anchor}>
                    <Button onClick={toggleDrawer(anchor, true)}>統計</Button>
                    <Drawer
                        anchor={anchor}
                        open={state[anchor]}
                        onClose={toggleDrawer(anchor, false)}
                    >
                        {list(anchor)}
                    </Drawer>
                </React.Fragment>
            ))}
        </div>
    );
}
