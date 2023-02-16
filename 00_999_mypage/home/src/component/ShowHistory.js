import Box from '@mui/material/Box';
import Pagination from '@mui/material/Pagination';
import {
    DataGrid,
    gridPageCountSelector,
    gridPageSelector,
    useGridApiContext,
    useGridSelector,
} from '@mui/x-data-grid';
import * as React from 'react';
import { useDictionalHistory } from '../api/api';
import ErrorPage from './ErrorPage';
import Statistic from './Statistic';

function CustomPagination() {
    const apiRef = useGridApiContext();
    const page = useGridSelector(apiRef, gridPageSelector);
    const pageCount = useGridSelector(apiRef, gridPageCountSelector);
    return (
        <Pagination
            color="primary"
            count={10}
            page={page + 1}
            siblingCount={0}
            onChange={(event, value) => apiRef.current.setPage(value - 1)}
        />
    );
}


export default function ShowHistory() {
    // const { data, isError } = useDictional("/tool/dic/v1/c")
    const { HistoryData, HistoryIsError } = useDictionalHistory()
    if (HistoryIsError) return <ErrorPage></ErrorPage>

    const rows = HistoryData.h.map((obj) => {
        return {
            id: obj.ID,
            col1: obj.Word,
            col2: obj.Info,
            col3: obj.CreatedAt,
        }
    })

    const columns = [
        { field: 'col1', headerName: 'Word', width: 150 },
        { field: 'col2', headerName: 'Info', width: 800 },
        { field: 'col3', headerName: 'CreatedDate', width: 150 },
    ]

    const hisData = {
        rows,
        columns
    }

    return (
        <Box sx={{ height: 350, width: '100%', border: '1px solid #4caf50' }} component='div'>
            <DataGrid
                pagination
                pageSize={10}
                rowsPerPageOptions={[5]}
                components={{
                    Pagination: CustomPagination,
                }}
                {...hisData}
            />
            < Statistic></Statistic>
        </Box>
    );
}
