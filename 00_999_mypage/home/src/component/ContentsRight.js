import { Box } from "@mui/system"
// import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';

const img = require("../assets/gopher.png")

//  backgroundColor: '#c5e1a5',
export const RightContents = ({ ds }) => {
    return <Box sx={{
        width: '100%', height: `${ds.isdmd ? '70vh' : '100vh'}`, backgroundColor: 'white',
        justifyContent: 'center', alignItems: 'center', display: 'flex',
    }} >
        <Card variant="inline">
            <CardMedia
                component="img"
                // height="100%"
                // width="100%"
                image={img}
                alt="so cool" />
        </Card>
    </Box>;

}