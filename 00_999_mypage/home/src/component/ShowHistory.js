import * as React from 'react';
import Box from '@mui/material/Box';
import {
    DataGrid,
    gridPageCountSelector,
    gridPageSelector,
    useGridApiContext,
    useGridSelector,
} from '@mui/x-data-grid';
import { useDemoData } from '@mui/x-data-grid-generator';
import Pagination from '@mui/material/Pagination';
import Statistic from './Statistic';

function CustomPagination() {
    const apiRef = useGridApiContext();
    const page = useGridSelector(apiRef, gridPageSelector);
    const pageCount = useGridSelector(apiRef, gridPageCountSelector);
    return (
        <Pagination
            color="primary"
            count={pageCount}
            page={page + 1}
            siblingCount={0}
            onChange={(event, value) => apiRef.current.setPage(value - 1)}
        />
    );
}

export default function ShowHistory() {
    const { data } = useDemoData({
        dataSet: 'Commodity',
        rowLength: 100,
        maxColumns: 6,
    });

    return (
        <Box sx={{ height: 350, width: '100%' }}>
            <DataGrid
                pagination

                pageSize={10}
                rowsPerPageOptions={[5]}
                components={{
                    Pagination: CustomPagination,
                }}
                {...data}
            />
            < Statistic></Statistic>
        </Box>
    );
}
