// import { Helmet } from 'react-helmet-async';
// import { faker } from '@faker-js/faker';
import { faker } from '@faker-js/faker/locale/ja';
import { Box, Grid, Typography } from '@mui/material';
// @mui
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import {ShowHistory} from '../api/api'
// components
// import Iconify from '../components/iconify';
// sections
// import {
//   AppNewsUpdate,
//   AppOrderTimeline
// } from '../sections/@dashboard/app';

import {useEffect,useState} from 'react';
import ShowStatus from "./showStatus";
import ShowProcess from "./showProcess";


// ----------------------------------------------------------------------

export default function History() {
  const theme = useTheme();
  const matches = useMediaQuery(theme.breakpoints.down('sm'));
  faker.setLocale('ja')

  const [HisData, setData] = useState({result:{records:[]}})

  //全件データ取得
  useEffect( () => {
    (async () =>{
      const data  =  await ShowHistory()
       setData(data)
    })()
  }, [])
  

  return (
      <Box  sx={{ flexGrow: 1,alignItems:"center" }}  component="div" maxWidth={matches ? "85vw" : "96vw"}>

        {/* <Grid container spacing={3}> */}
        <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 1, md: 3 }} columns={{ xs: 4, sm: 8, md: 12 }}>
          <Grid item xs={4} sm={7} md={12} component="div">
            <Typography variant="h4" sx={{ mb: 5 }}>
              依頼一覧
            </Typography>
          </Grid>

          {/* <Grid item xs={12} md={6} lg={8}> */}
          <Grid item xs={4} sm={4} md={10} component="div">
          <ShowStatus
              title="依頼一覧"
              list={HisData.result.records.map((item, index) => ({
                key: faker.datatype.uuid(),
                id: item.$id.value,
                title: item.request_title.value,
                description: item.request_contents.value,
                image: `af052.ico`,
                postedAt: item.create_date.value,
              }))
              }
            />
            {/* <ShowStatus
              title="依頼一覧"
              list={[...Array(15)].map((_, index) => ({
                id: faker.datatype.uuid(),
                title: faker.name.fullName(),
                description: faker.lorem.paragraph(),
                image: `af052.ico`,
                postedAt: faker.date.recent(),
              }))
              }
            /> */}
          </Grid>

          {/* <Grid item xs={12} md={6} lg={4}> */}
          <Grid item xs={4} sm={4} md={2} component="div">
            <ShowProcess
              title="対応進捗"
              list={[...Array(6)].map((_, index) => ({
                id: faker.datatype.uuid(),
                title: [
                  '問い合わせ受付',
                  '対応依頼済',
                  '対応中',
                  '対応完了',
                  '確認中',
                  '確認完了',
                  '終了',
                ][index],
                type: `order${index + 1}`,
                time: faker.date.past(10),
              }))}
            />
          </Grid>
        </Grid>
      </Box>
  
  );
}