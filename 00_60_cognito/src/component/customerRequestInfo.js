import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';
import FormControl from '@mui/material/FormControl';
import Input from '@mui/material/Input';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import * as React from 'react';
import { useState } from 'react';
import { NewRequest } from '../api/api';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';

const initData = {
    userName: "Demo",
    customerName: "笹塚太郎",
    requestContent: "発信DEMO作成",
    office: "笹塚センター",
    phoneNo: "08066322950",
    status: "1",
    receptionContent: "",
}

const StatusData = [
    {
        key: 0,
        value: "依頼中"
    },
    {
        key: 1,
        value: "対応中"
    },
    {
        key: 2,
        value: "対応完了"
    },
    {
        key: 3,
        value: "確認完了"
    },
]

export default function CustomerRequestInfo() {
    const [Info, setInfo] = useState(initData)

    const [open, setOpen] = React.useState(false);
    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down('md'));
  
    const handleClickOpen = () => {
      setOpen(true);
    };
  
    const handleClose = () => {
      setOpen(false);
    };

    const createQuest = () => {
        // console.log(Info.userName)
        NewRequest(Info)
        handleClickOpen()
    }

    const MyDiolog = ()=> {
        return (
            <Dialog
            fullScreen={fullScreen}
            open={open}
            onClose={handleClose}
            aria-labelledby="responsive-dialog-title"
          >
            <DialogTitle id="responsive-dialog-title">
              {"処理結果"}
            </DialogTitle>
            <DialogContent>
              <DialogContentText>
                依頼内容を作成しました。
                指定電話番号に通知完了。
              </DialogContentText>
            </DialogContent>
            <DialogActions>
              <Button autoFocus onClick={handleClose}>
                閉じる
              </Button>
              {/* <Button onClick={handleClose} autoFocus>
                Agree
              </Button> */}
            </DialogActions>
          </Dialog>
        )
    }

    return (
        <Box sx={{ flexGrow: 1 }} component="form">
            <MyDiolog></MyDiolog>    

            {/* 0,600,900 */}
            <Grid container spacing={{ xs: 2 }} columns={{ xs: 4, sm: 8, md: 12 }}>
                <Grid xs={4} sm={8} md={12} component="div">
                    {/* <Box> */}
                    <Typography variant='h4' component='h1'>■お客様相談</Typography>
                    <Typography variant='h5' component='h1'>閲覧者一覧</Typography>

                    <Divider></Divider>
                    {/* </Box> */}
                    <br></br>
                </Grid>
                <Grid xs={4} sm={4} md={2} component="div">
                    <FormControl variant="standard" fullWidth>
                        <InputLabel htmlFor="userName">受付者</InputLabel>
                        <Input id="userName" name="userName" variant="outlined" defaultValue={Info.userName} onBlur={(e) => {
                            setInfo({ ...Info, userName: e.target.value })
                        }} />
                    </FormControl>
                </Grid>
                <Grid xs={4} sm={4} md={2} component="div">
                    <FormControl variant="standard" fullWidth>
                        <InputLabel htmlFor="customerName">お客様名</InputLabel>
                        <Input id="customerName" name="customerName" variant='outlined' defaultValue={Info.customerName} onBlur={(e) => {
                            setInfo({ ...Info, customerName: e.target.value })
                        }

                        } />
                    </FormControl>
                </Grid>
                <Grid xs={4} sm={8} md={12} component="div">
                    <FormControl variant="standard" fullWidth>
                        <TextField id="requestContent" name="requestContent" defaultValue={Info.userName} label="問い合わせ内容"
                            multiline rows={4} variant='outlined' onBlur={(e) => {
                                setInfo({ ...Info, requestContent: e.target.value })
                            }} />
                    </FormControl>
                </Grid>

                <Grid xs={4} sm={4} md={2} component="div">
                    <FormControl variant="standard" fullWidth>
                        <InputLabel htmlFor="office">営業所：</InputLabel>
                        <Input id="office" name="office" defaultValue={Info.office} variant='outlined' onBlur={(e) => {
                            setInfo({ ...Info, office: e.target.value })
                        }} />
                    </FormControl>
                </Grid>
                <Grid xs={4} sm={4} md={2} component="div">
                    <FormControl variant="standard" fullWidth>
                        <InputLabel htmlFor="phoneNo">対応者電話番号：</InputLabel>
                        <Input id="phoneNo" name="phoneNo" defaultValue={Info.phoneNo} variant='outlined' onBlur={(e) => {
                            setInfo({ ...Info, phoneNo: e.target.value })
                        }} />
                    </FormControl>
                </Grid>
                <Grid xs={4} sm={4} md={2} component="div">
                    <FormControl variant="standard" fullWidth>
                        <InputLabel htmlFor="phoneNo">状態：</InputLabel>
                        {/* <Input id="phoneNo" defaultValue={Info.phoneNo} variant='outlined' /> */}
                        <Select
                            labelId="status-label"
                            id="status"
                            name="status"
                            value={Info.status}
                            onChange={(event) => setInfo({ ...initData, status: event.target.value })}
                            label="status"
                        >
                            <MenuItem value="">
                                <em>None</em>
                            </MenuItem>
                            {
                                StatusData.map((v, index) => (

                                    <MenuItem value={v.key} key={index}>{v.value}</MenuItem>
                                )
                                )
                            }

                        </Select>
                    </FormControl>
                </Grid>

                <Grid xs={4} sm={8} md={12} component="div">
                    <FormControl variant="standard" fullWidth>
                        <TextField id="receptionContent" name="receptionContent" defaultValue={Info.receptionContent} label="回答内容" multiline rows={4} variant='outlined' onBlur={(e) => {
                            setInfo({ ...Info, phoneNo: e.target.value })
                        }} />
                    </FormControl>
                </Grid>

                <Grid xs={4} sm={8} md={12} component="div">
                 
                    <Button variant="contained" color="success" size='large' onClick={createQuest} >
                        確認
                    </Button>
                    &nbsp;
                    <Button variant="outlined" color="error" size='large' onClick={() => {
                        //    window.location.href ="/";
                        window.location.reload()
                    }}>
                        取消
                    </Button>
                    &nbsp;
                    <Button variant="contained" color="success" size='large' disabled={true}>
                        編集
                    </Button>
                </Grid>

            </Grid>
        </Box>
    );
}