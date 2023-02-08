import Card from '@mui/material/Card';
import { CardHeader, Typography } from "@mui/material"
import { Box } from "@mui/system"
import Button from '@mui/material/Button';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';


export const LeftContents = ({ ds, id, children }) => {

    return <Box id={id} sx={{
        width: '100%', height: `${ds.isdmd ? '70vh' : '100vh'}`, backgroundColor: 'white',
        justifyContent: 'center', alignItems: 'center', display: 'flex',
    }}>
        <Card variant="inlined">
            <CardContent>
                <Typography variant={ds ? ds.Name : 'h2'} >
                    YAN WENSHUAI
                </Typography>
                <Typography variant={ds ? ds.SubTitle : 'h5'} gutterBottom align="left">
                    &nbsp;&nbsp;SYSTEM ENGINEER IN JAPAN
                </Typography>
                <Typography variant={ds ? ds.PhoneNo : 'h5'} gutterBottom sx={{ mt: 6 }}>
                    &nbsp;&nbsp;<b>Tel   </b> 080-9342-xxxx
                </Typography>
                <Typography variant={ds ? ds.Email : 'h5'} gutterBottom>
                    &nbsp;&nbsp;<b>Email </b>ywsrock@gmail.com
                </Typography>
                <Typography variant={ds ? ds.Ads : 'h5'} gutterBottom>
                    &nbsp;&nbsp;<b>Addr  </b> 東京都江東区
                </Typography>
                <Typography variant={ds ? ds.Ads : 'h5'} >
                    &nbsp;&nbsp;<b>From  </b> 中国内モンゴル
                </Typography>
            </CardContent>
            <CardActions>
                {children}
            </CardActions>
        </Card>
    </Box >;
}

export const ActionCard = ({ id, title }) => {
    return (<>
        <Button size="smail" href={`#${id}`} >{title}</Button>
    </>
    )
}

export const SelfIntroduction = ({ ds, id, title, content, children }) => {
    return <Box id={id} sx={{
        width: '100%', height: '100vh', backgroundColor: 'white',
        justifyContent: 'center', alignItems: 'center', display: 'flex'

    }}>
        <Card variant="inlined" sx={{ height: '100vh', width: '90%' }} >
            <CardHeader title={title}></CardHeader>
            <CardContent sx={{ height: '60%', overflow: 'auto' }} >
                {/* <Typography variant='h6' gutterBottom >
                    {title}
                </Typography> */}
                <Typography variant={ds ? ds.Name : 'body1'} >
                    {content}
                </Typography>
            </CardContent>
            <CardActions>
                {children}
            </CardActions>
        </Card>
    </Box>
}
