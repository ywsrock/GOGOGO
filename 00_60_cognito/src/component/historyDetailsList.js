// import { Helmet } from 'react-helmet-async';
// @mui
import {
  Avatar, Box, Button, Card, Checkbox, Grid, IconButton, MenuItem, Paper, Popover, Stack, Table, TableBody,
  TableCell, TableContainer,
  TablePagination, TableRow, Typography
} from '@mui/material';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import { filter } from 'lodash';
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
// mock
// import USERLIST from '../_mock/user';
import { ShowHistory } from '../api/api';
import Iconify from '../components/iconify';
import Label from '../components/label';
import Scrollbar from '../components/scrollbar';
// setting
import { RowsPerPage } from '../menu/setting';
// sections
import { UserListHead, UserListToolbar } from '../sections/@dashboard/user';

const TABLE_HEAD = [
  { id: 'name', label: '対応者', alignRight: false },
  { id: 'company', label: 'お客様名', alignRight: false },
  { id: 'role', label: '問い合わせ', alignRight: false },
  { id: 'isVerified', label: '対応日', alignRight: false },
  { id: 'status', label: 'Status', alignRight: false },
  { id: '' },
];
// ----------------------------------------------------------------------

function descendingComparator(a, b, orderBy) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function getComparator(order, orderBy) {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function applySortFilter(array, comparator, query) {
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  if (query) {
    return filter(array, (_user) => _user.name.toLowerCase().indexOf(query.toLowerCase()) !== -1);
  }
  return stabilizedThis.map((el) => el[0]);
}

export default function HistoryDetailsList() {


  const [USERLIST, setUSERLIST] = useState([])

  //全件データ取得
  useEffect(() => {
    (async () => {
      const data = await ShowHistory()
      // const users = []

      const users = data.result.records.map((v, i) => ({
        // let user = {
        "id": v.$id.value,
        "name": v.create_user.value.name,
        "company": v.create_user.value.name,
        "isVerified": v.create_date.value,
        "status": v.request_status.value,
        "role": v.request_title.value,
        // }
      }))
      
      setUSERLIST(users)
    })()
  }, [])

  //const [message, setMessage] = React.useState('');
  //const data = useMovieData();
  // const handleRowClick = (params) => {
  //   setMessage(`Movie "${params.row.title}" clicked`);
  //   //<Link to="/ShowCustomerRequestInfo" style={{ textDecoration: 'none', color: 'inherit' }}></Link>
  // };

  const [open, setOpen] = useState(null);

  const [page, setPage] = useState(0);

  const [order, setOrder] = useState('asc');

  const [selected, setSelected] = useState([]);

  const [orderBy, setOrderBy] = useState('name');

  const [filterName, setFilterName] = useState('');

  const [rowsPerPage, setRowsPerPage] = useState(RowsPerPage);

  const handleOpenMenu = (event) => {
    setOpen(event.currentTarget);
  };

  const handleCloseMenu = () => {
    setOpen(null);
  };

  const handleRequestSort = (event, property) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleSelectAllClick = (event) => {
    if (event.target.checked) {
      const newSelecteds = USERLIST.map((n) => n.name);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleClick = (event, name) => {
    const selectedIndex = selected.indexOf(name);
    let newSelected = [];

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, name);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex), 
        selected.slice(selectedIndex + 1),
      );
    }
    setSelected(newSelected);
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setPage(0);
    setRowsPerPage(parseInt(event.target.value, 10));
  };

  const handleFilterByName = (event) => {
    setPage(0);
    setFilterName(event.target.value);
  };

  const isSelected = (name) => selected.indexOf(name) !== -1;

  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - USERLIST.length) : 0;

  const filteredUsers = applySortFilter(USERLIST, getComparator(order, orderBy), filterName);

  const isNotFound = !filteredUsers.length && !!filterName;

  const theme = useTheme();
  const matches = useMediaQuery(theme.breakpoints.down('sm'));

  return (
    <Box sx={{ flexGrow: 1 }} component="div">

      <Grid container rowSpacing={1} columns={{ xs: 4, sm: 8, md: 12 }}>

        <Grid item xs={4} sm={7} md={12} component="div">
          <Typography variant="h4" gutterBottom>
            履歴一覧
          </Typography>
        </Grid>

        <Grid item xs={4} sm={7} md={12} component="div">
          <Card>
            <Grid container spacing={{ xs: 1 }} columns={{ xs: 4, sm: 8, md: 12 }}>
              <Grid item xs={3} sm={7} md={8} component="div">
                <UserListToolbar numSelected={selected.length} filterName={filterName} onFilterName={handleFilterByName} />
              </Grid>
              <Grid item xs={3} sm={7} md={4} component="div">
                <div>
                  <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} sx={{ backgroundColor: '#ffc107', mt: '20px', ml: '20px' }}>
                    <Link to="/CustomerRequestInfo" style={{ textDecoration: 'none', color: 'inherit' }}>新規作成</Link>
                  </Button>
                </div>
              </Grid>
            </Grid>

            <Scrollbar sx={{ height: '75vh', width: matches ? '85vw' : '93vw' }}>
              <TableContainer component='div' >
                <Table sx={{ minWidth: 20 }} aria-label="simple table">
                  {/* <caption>A basic table example with a caption</caption> */}
                  <UserListHead
                    order={order}
                    orderBy={orderBy}
                    headLabel={TABLE_HEAD}
                    rowCount={USERLIST.length}
                    numSelected={selected.length}
                    onRequestSort={handleRequestSort}
                    onSelectAllClick={handleSelectAllClick}
                  />
                  <TableBody>
                    {filteredUsers.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row, index) => {
                      const { id, name, role, status, company, avatarUrl, isVerified } = row;
                      // const selectedUser = selected.indexOf(name) !== -1;
                      const selectedUser = isSelected(row.name);
                      const labelId = `enhanced-table-checkbox-${index}`;

                      return (
                        <TableRow 
                        hover 
                        onClick={(event) => handleClick(event, row.name)} 
                        role="checkbox" 
                        aria-checked={selectedUser} 
                        tabIndex={-1} 
                        key={row.$id} 
                        selected={selectedUser}
                        >
                          <TableCell padding="checkbox">
                            <Checkbox 
                            checked={selectedUser} 
                            inputProps={{
                            'aria-labelledby': labelId,
                          }} 
                          onChange={(event) => handleClick(event, name)} />
                          </TableCell>

                          <TableCell component="th" scope="row" padding="none">
                            <Stack direction="row" alignItems="center" spacing={2}>
                              <Avatar alt={name} src={avatarUrl} />
                              <Typography variant="subtitle2" noWrap>
                              <Link  to={`/ShowCustomerRequestInfo/${id}`} color="inherit" variant="subtitle2" underline="hover" style={{ textDecoration: 'none', color: 'inherit' }}>
                                {name}
                              </Link>
                              </Typography>
                            </Stack>
                          </TableCell>

                          <TableCell align="left" sx={{ whiteSpace: 'nowrap' }}>{company}</TableCell>

                          <TableCell align="left" sx={{ whiteSpace: 'nowrap' }}>{role}</TableCell>

                          <TableCell align="left" sx={{ whiteSpace: 'nowrap' }}>{isVerified}</TableCell>

                          <TableCell align="left" sx={{ whiteSpace: 'nowrap' }}>
                            {/* <Label color={(status === 'banned' && 'error') || 'success'}>{sentenceCase(status)}</Label> */}
                            <Label color={(status === 'banned' && 'error') || 'success'}>{status}</Label>
                            
                          </TableCell>

                          <TableCell align="right" sx={{ whiteSpace: 'nowrap' }}>
                            <IconButton size="large" color="inherit" onClick={handleOpenMenu}>
                              <Iconify icon={'eva:more-vertical-fill'} />
                            </IconButton>
                          </TableCell>
                        </TableRow>
                      );
                    })}
                    {emptyRows > 0 && (
                      <TableRow style={{ height: 53 * emptyRows }}>
                        <TableCell colSpan={6} />
                      </TableRow>
                    )}
                  </TableBody>

                  {isNotFound && (
                    <TableBody>
                      <TableRow>
                        <TableCell align="center" colSpan={6} sx={{ py: 3 }}>
                          <Paper
                            sx={{
                              textAlign: 'center',
                            }}
                          >
                            <Typography variant="h6" paragraph>
                              Not found
                            </Typography>

                            <Typography variant="body2">
                              No results found for &nbsp;
                              <strong>&quot;{filterName}&quot;</strong>.
                              <br /> Try checking for typos or using complete words.
                            </Typography>
                          </Paper>
                        </TableCell>
                      </TableRow>
                    </TableBody>
                  )}
                </Table>
              </TableContainer>
            </Scrollbar>

            <TablePagination
              rowsPerPageOptions={[10, 10, 25]}
              component="div"
              count={USERLIST.length}
              rowsPerPage={rowsPerPage}
              page={page}
              onPageChange={handleChangePage}
              onRowsPerPageChange={handleChangeRowsPerPage}
              sx={{ float: 'left' }}
            />
          </Card>
        </Grid>

        <Grid item xs={4} sm={7} md={12} component="div">
          <Popover
            open={Boolean(open)}
            anchorEl={open}
            onClose={handleCloseMenu}
            anchorOrigin={{ vertical: 'top', horizontal: 'left' }}
            transformOrigin={{ vertical: 'top', horizontal: 'right' }}
            PaperProps={{
              sx: {
                p: 1,
                width: 140,
                '& .MuiMenuItem-root': {
                  px: 1,
                  typography: 'body2',
                  borderRadius: 0.75,
                },
              },
            }}
          >
            <MenuItem>
              <Iconify icon={'eva:edit-fill'} sx={{ mr: 2 }} />
              Edit
            </MenuItem>

            <MenuItem sx={{ color: 'error.main' }}>
              <Iconify icon={'eva:trash-2-outline'} sx={{ mr: 2 }} />
              Delete
            </MenuItem>
          </Popover>
        </Grid>
      </Grid >
    </Box >
  );
}