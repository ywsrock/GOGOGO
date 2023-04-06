// @mui
import { Box, Button, Card, CardHeader, Divider, Grid, ListItemButton, Stack, Typography } from '@mui/material';
import PropTypes from 'prop-types';
import React from 'react';
// components
import Iconify from '../components/iconify';
import Scrollbar from '../components/scrollbar';
import { Link } from "react-router-dom";

// ----------------------------------------------------------------------

ShowStatus.propTypes = {
  title: PropTypes.string,
  subheader: PropTypes.string,
  list: PropTypes.array.isRequired,
};

export default function ShowStatus({ title, subheader, list, ...other }) {
  return (
    <Card {...other}>
      <Grid container rowSpacing={1} columns={{ xs: 4, sm: 8, md: 12 }}>
        <Grid item xs={4} sm={8} md={12} component="div">
          <CardHeader title={title} subheader={subheader} sx={{ backgroundColor: "#ffc107" }} />
        </Grid>
        <Divider />
        <Grid item xs={4} sm={7} md={12} component="div">
        <Scrollbar sx={{ height: "80vh" }}>
            <Stack spacing={2} sx={{ p: 2, pr: 0, height: "80vh" }}>
              {list.map((news) => {
                return (
                  <React.Fragment key={news.key} >
                  {/* <Link to="/CustomerRequestInfo/1">     */}
                    <NewsItem key={news.id} news={news} /> 
                  {/* </Link> */}
                    <Divider />
                  </React.Fragment>
                )
              })}
            </Stack>
          </Scrollbar>
        </Grid>

        <Divider />
        <Grid item xs={4} sm={8} md={12} component="div">
          <Box sx={{ p: 2, textAlign: 'right' }}>
            <Button size="small" color="inherit" endIcon={<Iconify icon={'eva:arrow-ios-forward-fill'} />}>
              View all
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Card >
  );
}

// ----------------------------------------------------------------------

NewsItem.propTypes = {
  news: PropTypes.shape({
    description: PropTypes.string,
    image: PropTypes.string,
    postedAt: PropTypes.string,
    title: PropTypes.string,
  }),
};

function NewsItem({ news }) {
  const { image, title, description, postedAt,id} = news;

  return (
    <Stack direction="row" alignItems="center" spacing={2}>
      <ListItemButton>
        <Box component="img" alt={title} src={image} sx={{ width: 20, height: 20, borderRadius: 1.5, flexShrink: 0 }} />
        <Box sx={{ minWidth: 240, flexGrow: 1 }}>
          <Link  to={`/ShowCustomerRequestInfo/${id}`} color="inherit" variant="subtitle2" underline="hover" style={{ textDecoration: 'none', color: 'inherit' }}>
            {title}
          </Link>

          <Typography variant="body2" sx={{ color: 'text.secondary' }} >
            {description}
          </Typography>
        </Box>

        <Typography variant="caption" sx={{ pr: 3, flexShrink: 0, color: 'text.secondary' }}>
          {postedAt}
        </Typography>

      </ListItemButton>
    </Stack>
  );
}
