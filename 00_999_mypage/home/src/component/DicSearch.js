import Grid from '@mui/material/Unstable_Grid2';
import Card from '@mui/material/Card';
import { Button, CardActions, CardContent } from '@mui/material';
import TextField from '@mui/material/TextField';

export const SearchBar = () => {
    return (
        <Grid container columns={{ xs: 12, sm: 12, md: 12 }} direction="row" justifyContent="center" alignItems="flex-start" spacing={1} rowSpacing={0}>
            <Grid xs={12} sm={12} md={6}>
                <Card variant='outlined' sx={{ borderStyle: 'none' }}>
                    <CardContent>
                        <TextField fullWidth='true' sx={{ height: '4em' }} label="keyword" variant="outlined"></TextField>
                    </CardContent>
                </Card>
            </Grid>
            <Grid xs={12} sm={12} md={2} spacing={2}>
                <Card variant='outlined' sx={{ borderStyle: 'none' }}>
                    <CardActions>
                        <Button variant="contained" fullWidth='true' color="success" sx={{ height: '4rem', }}>
                            検索
                        </Button>
                    </CardActions>
                </Card>
            </Grid>
        </Grid>
    )
}