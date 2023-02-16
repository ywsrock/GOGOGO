import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import * as React from 'react';
import { StsDataInit } from '../state/State';

export default function ModalTable({ type = "d", StsData = StsDataInit }) {
    const rows = type === "d" ? StsData.dc : StsData.dm

    return (
        <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="right">{type === "d" ? "Day" : "Month"}</TableCell>
                        <TableCell align="right">{type === "d" ? "DayCount" : "MonthCount"}</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map((row, index) => (
                        <TableRow
                            key={index}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell align="right">{type === "d" ? row.day : row.month}</TableCell>
                            <TableCell align="right">{type === "d" ? row.dayCount : row.monthCount}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}
