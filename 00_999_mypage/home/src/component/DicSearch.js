import { Button, CardActions, CardContent } from '@mui/material';
import Card from '@mui/material/Card';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Unstable_Grid2';
import React, { useState } from 'react';
import { API } from '../api/api';


export const SearchBar = (props) => {
    const [word, setWord] = useState("")
    // const [url, setUrl] = useState("/tool/dic/v1/c")

    // var keyword = ""
    // const { WordData, WordIsError } = useDictionalWord(keyword)


    const Query = async (w) => {
        let url = `/tool/dic/v1/c?key=${w}`
        let data1 = await API.get(url).then((res) => res.data)
        props.setWordData(data1)
    }

    return (
        <Grid container columns={{ xs: 12, sm: 12, md: 12 }} direction="row" justifyContent="center" alignItems="flex-start" spacing={1} rowSpacing={0}>
            <Grid xs={12} sm={12} md={6}>
                <Card variant='outlined' sx={{ borderStyle: 'none' }}>
                    <CardContent>
                        <TextField fullWidth={true} sx={{ height: '4em' }} variant="outlined"
                            value={word} onChange={(e) => setWord(e.target.value)}></TextField>
                    </CardContent>
                </Card>
            </Grid>
            <Grid xs={12} sm={12} md={2} spacing={2}>
                <Card variant='outlined' sx={{ borderStyle: 'none' }}>
                    <CardActions>
                        <Button variant="contained" fullWidth={true} color="success" sx={{ height: '4rem', }}
                            onClick={() => Query(word)}>
                            {/* <Link to={url}>検索</Link> */}
                            検索
                        </Button>
                    </CardActions>
                </Card>
            </Grid>
        </Grid>
    )

}

