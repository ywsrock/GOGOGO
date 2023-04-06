import { Box, Card, Grid, Typography } from '@mui/material';
import SystemSettingDetails from './system-setting-details';

   export default function SystemSetting() {
    return (
      <Box sx={{ flexGrow: 1 }} component="div">
        <Grid container rowSpacing={1} columns={{ xs: 4, sm: 8, md: 12 }}>
  
          <Grid item xs={4} sm={7} md={12} component="div">
            <Typography variant="h4" gutterBottom>
               システム設定
            </Typography>
          </Grid>
  
          <Grid item xs={4} sm={7} md={12} component="div">
            <Card>
              <SystemSettingDetails />
            </Card>
          </Grid>
  
        </Grid >
      </Box >
    );
  }