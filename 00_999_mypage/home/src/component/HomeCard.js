import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ShareIcon from '@mui/icons-material/Share';
import Avatar from '@mui/material/Avatar';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardHeader from '@mui/material/CardHeader';
import CardMedia from '@mui/material/CardMedia';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import { red } from '@mui/material/colors';
import { styled } from '@mui/material/styles';
import * as React from 'react';
import './HomeCard.css';

const ExpandMore = styled((props) => {
    const { expand, ...other } = props;
    return <IconButton {...other} />;
})(({ theme, expand }) => ({
    transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
        duration: theme.transitions.duration.shortest,
    }),
}));

export default function HomeCard({ item }) {
    const [expanded, setExpanded] = React.useState(false);

    const handleExpandClick = () => {
        setExpanded(!expanded);
    };

    const cardImg = require("../assets/img.png");

    return (
        <Card className='p1_Card'>
            <CardHeader
                avatar={
                    <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
                        {item.id}
                    </Avatar>
                }
                // action={
                //     <IconButton aria-label="settings">
                //         <MoreVertIcon />
                //     </IconButton>
                // }
                title={item.subTitle}
                subheader={`${item.startDate} 〜 ${item.endDate}`}
            />
            <CardMedia
                component="img"
                height="50"
                image={cardImg}
                alt="Paella dish"
            />
            <CardContent>
                <Typography variant="body2" color="text.secondary" sx={{ color: 'black' }}>
                    {item.description}
                </Typography>
            </CardContent>
            <CardActions disableSpacing>
                <IconButton aria-label="add to favorites">
                    <FavoriteIcon />
                </IconButton>
                <IconButton aria-label="share">
                    <ShareIcon />
                </IconButton>
                <ExpandMore
                    expand={expanded}
                    onClick={handleExpandClick}
                    aria-expanded={expanded}
                    aria-label="show more"
                >
                    <ExpandMoreIcon />
                </ExpandMore>
            </CardActions>
            <Collapse in={expanded} timeout="auto" unmountOnExit>
                <CardContent>
                    <Typography paragraph>Method:</Typography>
                    <Typography paragraph>
                        Heat 1/2 cup of the broth in a pot until simmering, add saffron and set
                        aside for 10 minutes.
                    </Typography>
                    <Typography paragraph sx={{ color: 'black' }}>
                        {item.description}
                    </Typography>
                    <Typography paragraph>
                        {item.description}
                    </Typography>
                    <Typography>
                        Set aside off of the heat to let rest for 10 minutes, and then serve.
                    </Typography>
                </CardContent>
            </Collapse>
        </Card >
    );
}
