import React from 'react';
import { Box, Card, Grid, Typography } from '@mui/material';
import NotificationSettingDetails from './notification-setting-details';
  
   export default function NotificationSetting() {
    return (
      <Box sx={{ flexGrow: 1 }} component="div">
        <Grid container rowSpacing={1} columns={{ xs: 4, sm: 8, md: 12 }}>
  
          <Grid item xs={4} sm={7} md={12} component="div">
            <Typography variant="h4" gutterBottom>
               通知設定
            </Typography>
          </Grid>
  
          <Grid item xs={4} sm={7} md={12} component="div">
            <Card>
              <Grid container spacing={{ xs: 1 }} columns={{ xs: 4, sm: 8, md: 12 }}>
                <Grid item xs={3} sm={7} md={8} component="div">
                  <NotificationSettingDetails />
                </Grid>
              </Grid>
            </Card>
          </Grid>
  
        </Grid >
      </Box >
    );
  }